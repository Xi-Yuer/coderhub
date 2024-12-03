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

type ArticleRepository interface {
	CreateArticle(article *model.Articles) error
	GetArticleByID(id int64) (*model.Articles, error)
	UpdateArticle(article *model.Articles) error
	LikeArticle(id int64, increment int64) error
	DeleteArticle(id int64) error
}
type ArticleRepositoryImpl struct {
	DB       *gorm.DB
	Redis    CacheDB.RedisDB
	minLikes int32
}

func NewArticleRepositoryImpl(db *gorm.DB, rdb CacheDB.RedisDB) *ArticleRepositoryImpl {
	return &ArticleRepositoryImpl{
		DB:       db,
		Redis:    rdb,
		minLikes: 10,
	}
}

func (r *ArticleRepositoryImpl) CreateArticle(article *model.Articles) error {
	if err := r.DB.Create(article).Error; err != nil {
		return err
	}
	// 创建后设置缓存
	return r.setCache(article.CacheKeyByID(article.ID), article)
}

func (r *ArticleRepositoryImpl) GetArticleByID(id int64) (*model.Articles, error) {
	var article model.Articles
	key := article.CacheKeyByID(id)

	// 尝试从缓存获取
	if cached, err := r.getCache(key); err == nil {
		return cached, nil
	}

	// 简化数据库查询
	if err := r.DB.First(&article, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("文章不存在: %v", id)
		}
		return nil, err
	}

	// 异步设置缓存，避免影响主流程
	go func() {
		_ = r.setCache(key, &article)
	}()

	return &article, nil
}

func (r *ArticleRepositoryImpl) UpdateArticle(article *model.Articles) error {
	if article.ID <= 0 {
		return errors.New("无效的文章ID")
	}

	// 使用事务确保数据一致性
	return r.DB.Transaction(func(tx *gorm.DB) error {
		// 检查文章是否存在
		var count int64
		if err := tx.Model(&model.Articles{}).Select("id").Where("id = ?", article.ID).Count(&count).Error; err != nil {
			return err
		}
		if count == 0 {
			return errors.New("文章不存在")
		}

		// 更新数据库
		if err := tx.Model(article).Updates(article).Error; err != nil {
			return err
		}

		// 删除旧缓存并设置新缓存
		key := article.CacheKeyByID(article.ID)
		if err := r.delCache(key); err != nil {
			return err
		}
		return r.setCache(key, article)
	})
}

func (r *ArticleRepositoryImpl) DeleteArticle(id int64) error {
	article, err := r.GetArticleByID(id)
	if err != nil {
		return err
	}
	err = r.delCache(article.CacheKeyByID(id))
	if err != nil {
		return err
	}
	return r.DB.Delete(&model.Articles{}, id).Error
}

// 给文章点赞（需要考虑高并发，微服务架构）
func (r *ArticleRepositoryImpl) LikeArticle(id int64, increment int64) error {
	// 重试相关配置
	maxRetries := 3
	retryDelay := 100 * time.Millisecond

	// 锁相关的配置
	lockKey := fmt.Sprintf("article_lock:%d", id)
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
			_, releaseErr := script.Run(context.Background(), r.Redis.Pipeline(), []string{lockKey}, lockValue).Result()
			if releaseErr != nil {
				fmt.Printf("释放锁失败: %v\n", releaseErr)
			}
		}()

		// 更新逻辑
		err = r.DB.WithContext(context.Background()).Transaction(func(tx *gorm.DB) error {
			var article model.Articles
			err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&article, id).Error
			if err != nil {
				return err
			}

			// 更新点赞数
			result := tx.Model(&article).
				Where("id = ?", id).
				Updates(map[string]interface{}{
					"like_count": gorm.Expr("like_count + ?", increment),
					"version":    gorm.Expr("version + 1"),
				})
			if result.RowsAffected == 0 {
				return errors.New("并发更新失败")
			}

			// 重新查询最新数据
			var updatedArticle model.Articles
			if err := tx.First(&updatedArticle, id).Error; err != nil {
				return err
			}

			// 使用最新数据更新缓存
			if updatedArticle.LikeCount >= int64(r.minLikes) {
				if err := r.setCache(article.CacheKeyByID(id), &updatedArticle); err != nil {
					return err // 缓存更新失败时回滚事务
				}
			} else {
				if err := r.delCache(article.CacheKeyByID(id)); err != nil {
					return err // 缓存删除失败时回滚事务
				}
			}

			return nil
		})

		return err
	}

	return fmt.Errorf("更新失败，请稍后重试")
}

func (r *ArticleRepositoryImpl) getCache(key string) (*model.Articles, error) {
	var article model.Articles
	data, err := r.Redis.Get(key)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(data), &article); err != nil {
		return nil, err
	}
	return &article, nil
}

func (r *ArticleRepositoryImpl) setCache(key string, article *model.Articles) error {
	data, err := json.Marshal(article)
	if err != nil {
		return err
	}
	return r.Redis.Set(key, string(data))
}

func (r *ArticleRepositoryImpl) delCache(key string) error {
	return r.Redis.Del(key)
}
