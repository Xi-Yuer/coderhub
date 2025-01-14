package repository

import (
	"coderhub/model"
	"coderhub/shared/storage"
	"context"

	"gorm.io/gorm"
)

type CommentRelationLikeRepository interface {
	Create(ctx context.Context, commentRelationLike *model.CommentRelationLike) error
	Delete(ctx context.Context, commentRelationLike *model.CommentRelationLike) error
	Get(ctx context.Context, commentRelationLike *model.CommentRelationLike) bool
	List(ctx context.Context, commentID int64) (int64, error)
	// 批量获取评论点赞数
	BatchList(ctx context.Context, commentIDs []int64) (map[int64]int64, error)
}

type commentRelationLikeRepository struct {
	DB    *gorm.DB
	Redis storage.RedisDB
}

func NewCommentRelationLikeRepository(db *gorm.DB, redis storage.RedisDB) CommentRelationLikeRepository {
	return &commentRelationLikeRepository{
		DB:    db,
		Redis: redis,
	}
}

// 创建评论点赞
func (r *commentRelationLikeRepository) Create(ctx context.Context, commentRelationLike *model.CommentRelationLike) error {
	return r.DB.Create(commentRelationLike).Error
}

// 删除评论点赞
func (r *commentRelationLikeRepository) Delete(ctx context.Context, commentRelationLike *model.CommentRelationLike) error {
	return r.DB.Delete(commentRelationLike, "comment_id = ? AND user_id = ?", commentRelationLike.CommentID, commentRelationLike.UserID).Error
}

// 获取评论是否被用户点赞
func (r *commentRelationLikeRepository) Get(ctx context.Context, commentRelationLike *model.CommentRelationLike) bool {
	var count int64
	r.DB.Model(commentRelationLike).Where("comment_id = ? AND user_id = ?", commentRelationLike.CommentID, commentRelationLike.UserID).Count(&count)
	return count > 0
}

// 获取评论点赞数
func (r *commentRelationLikeRepository) List(ctx context.Context, commentID int64) (int64, error) {
	var count int64
	r.DB.Model(&model.CommentRelationLike{}).Where("comment_id = ?", commentID).Count(&count)
	return count, nil
}

// 批量获取评论点赞数
func (r *commentRelationLikeRepository) BatchList(ctx context.Context, commentIDs []int64) (map[int64]int64, error) {
	commentRelationLikes := make([]model.CommentRelationLike, 0)
	err := r.DB.Where("comment_id IN (?)", commentIDs).Find(&commentRelationLikes).Error
	if err != nil {
		return nil, err
	}
	commentLikeCountMap := make(map[int64]int64)
	for _, commentRelationLike := range commentRelationLikes {
		if _, ok := commentLikeCountMap[commentRelationLike.CommentID]; !ok {
			commentLikeCountMap[commentRelationLike.CommentID] = 0
		}
		commentLikeCountMap[commentRelationLike.CommentID]++
	}
	return commentLikeCountMap, nil
}
