package repository

import (
	"coderhub/model"
	"coderhub/shared/CacheDB"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// CommentRepository 评论仓储接口
type CommentRepository interface {
	Create(ctx context.Context, comment *model.Comment) error
	GetByID(ctx context.Context, id int64) (*model.Comment, error)
	Delete(ctx context.Context, id int64) error
	ListByArticleID(ctx context.Context, articleID int64, page int64, pageSize int64) ([]model.Comment, int64, error)
	ListReplies(ctx context.Context, parentID int64) ([]model.Comment, error)
	UpdateLikeCount(ctx context.Context, id int64, increment int32) error
}

var ErrConcurrentUpdate = errors.New("并发更新冲突，请重试")

// commentRepository 评论仓储实现
type commentRepository struct {
	DB       *gorm.DB
	Redis    CacheDB.RedisDB
	minLikes int32
}

// NewCommentRepository 创建评论仓储实例
func NewCommentRepository(db *gorm.DB, rdb CacheDB.RedisDB) CommentRepository {
	return &commentRepository{
		DB:       db,
		Redis:    rdb,
		minLikes: 10,
	}
}

// Create 创建评论
func (r *commentRepository) Create(ctx context.Context, comment *model.Comment) error {
	return r.DB.WithContext(ctx).Create(comment).Error
}

// GetByID 根据ID获取评论，优先从缓存获取
func (r *commentRepository) GetByID(ctx context.Context, id int64) (*model.Comment, error) {
	comment := &model.Comment{}
	cacheKey := comment.CacheKeyByID(id)

	// 尝试从缓存获取
	data, err := r.Redis.Get(cacheKey)
	if err == nil {
		err = json.Unmarshal([]byte(data), comment)
		if err == nil {
			return comment, nil
		}
	}

	// 从数据库获取
	err = r.DB.WithContext(ctx).First(comment, id).Error
	if err != nil {
		return nil, err
	}

	// 热门评论写入缓存
	if comment.LikeCount >= r.minLikes {
		go func() {
			err := r.cacheComment(comment)
			if err != nil {

			}
		}()
	}

	return comment, nil
}

// Delete 删除评论并清除缓存
func (r *commentRepository) Delete(ctx context.Context, id int64) error {
	err := r.DB.WithContext(ctx).Delete(&model.Comment{}, id).Error
	if err != nil {
		return err
	}
	go func() {
		err := r.deleteCache(id)
		if err != nil {

		}
	}()
	return nil
}

// UpdateLikeCount 更新点赞数并处理缓存，使用Redis分布式锁
func (r *commentRepository) UpdateLikeCount(ctx context.Context, id int64, increment int32) error {
	// 重试相关配置
	maxRetries := 3
	retryDelay := 100 * time.Millisecond

	// 锁相关的配置
	lockKey := fmt.Sprintf("comment_lock:%d", id)
	lockTTL := 3 * time.Second

	// 重试循环
	for i := 0; i < maxRetries; i++ {
		lockValue := fmt.Sprintf("%d:%d", id, time.Now().UnixNano())

		// 获取分布式锁
		acquired, err := r.Redis.SetNX(lockKey, lockValue, lockTTL)
		if err != nil {
			return fmt.Errorf("获取分布式锁失败: %w", err)
		}

		if !acquired {
			// 如果是最后一次重试，则返回错误
			if i == maxRetries-1 {
				return fmt.Errorf("系统繁忙，请稍后重试")
			}
			// 等待一段时间后重试
			time.Sleep(retryDelay)
			continue
		}

		// 获取到锁后的处理逻辑...
		defer func() {
			// 使用Lua脚本确保安全释放锁
			script := r.Redis.NewScript(`
				if redis.call("GET", KEYS[1]) == ARGV[1] then
					return redis.call("DEL", KEYS[1])
				end
				return 0
			`)
			_, releaseErr := script.Run(ctx, r.Redis.Pipeline(), []string{lockKey}, lockValue).Result()
			if releaseErr != nil {
				fmt.Printf("释放锁失败: %v\n", releaseErr)
			}
		}()

		// 更新逻辑
		err = r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
			var comment model.Comment
			err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&comment, id).Error
			if err != nil {
				return err
			}

			// 更新点赞数
			result := tx.Model(&comment).
				Where("id = ? AND version = ?", id, comment.Version).
				Updates(map[string]interface{}{
					"like_count": gorm.Expr("like_count + ?", increment),
					"version":    gorm.Expr("version + 1"),
				})
			if result.RowsAffected == 0 {
				return ErrConcurrentUpdate
			}

			// 重新查询最新数据
			var updatedComment model.Comment
			if err := tx.First(&updatedComment, id).Error; err != nil {
				return err
			}

			// 使用最新数据更新缓存
			if updatedComment.LikeCount >= r.minLikes {
				if err := r.cacheComment(&updatedComment); err != nil {
					return err // 缓存更新失败时回滚事务
				}
			} else {
				if err := r.deleteCache(id); err != nil {
					return err // 缓存删除失败时回滚事务
				}
			}

			return nil
		})

		return err
	}

	return fmt.Errorf("更新失败，请稍后重试")
}

// ListByArticleID 获取文章的评论列表
func (r *commentRepository) ListByArticleID(ctx context.Context, articleID int64, page int64, pageSize int64) ([]model.Comment, int64, error) {
	var comments []model.Comment
	var total int64
	err := r.DB.WithContext(ctx).
		Where("article_id = ? AND parent_id = 0", articleID).
		Order("created_at DESC").
		Offset(int((page - 1) * pageSize)).
		Limit(int(pageSize)).
		Find(&comments).Error
	if err != nil {
		return nil, 0, err
	}
	err = r.DB.WithContext(ctx).Model(&model.Comment{}).Where("article_id = ? AND parent_id = 0", articleID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	return comments, total, nil
}

// ListReplies 获取评论的回复列表
func (r *commentRepository) ListReplies(ctx context.Context, parentID int64) ([]model.Comment, error) {
	var replies []model.Comment
	err := r.DB.WithContext(ctx).
		Where("parent_id = ?", parentID).
		Order("created_at ASC").
		Find(&replies).Error
	return replies, err
}

// 内部辅助方法
func (r *commentRepository) cacheComment(comment *model.Comment) error {
	data, err := json.Marshal(comment)
	if err != nil {
		return err
	}
	cacheKey := comment.CacheKeyByID(comment.ID)
	return r.Redis.Set(cacheKey, string(data))
}

func (r *commentRepository) deleteCache(id int64) error {
	cacheKey := (&model.Comment{}).CacheKeyByID(id)
	return r.Redis.Del(cacheKey)
}
