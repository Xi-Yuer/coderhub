package repository

import (
	"coderhub/model"
	"coderhub/shared/CacheDB"
	"context"

	"gorm.io/gorm"
)

type CommentRelationLikeRepository interface {
	Create(ctx context.Context, commentRelationLike *model.CommentRelationLike) error
	Delete(ctx context.Context, commentRelationLike *model.CommentRelationLike) error
	Get(ctx context.Context, commentRelationLike *model.CommentRelationLike) error
}

type commentRelationLikeRepository struct {
	DB    *gorm.DB
	Redis CacheDB.RedisDB
}

func NewCommentRelationLikeRepository(db *gorm.DB, redis CacheDB.RedisDB) CommentRelationLikeRepository {
	return &commentRelationLikeRepository{
		DB:    db,
		Redis: redis,
	}
}

func (r *commentRelationLikeRepository) Create(ctx context.Context, commentRelationLike *model.CommentRelationLike) error {
	return r.DB.Create(commentRelationLike).Error
}

func (r *commentRelationLikeRepository) Delete(ctx context.Context, commentRelationLike *model.CommentRelationLike) error {
	return r.DB.Delete(commentRelationLike).Error
}

func (r *commentRelationLikeRepository) Get(ctx context.Context, commentRelationLike *model.CommentRelationLike) error {
	return r.DB.Where(commentRelationLike).First(commentRelationLike).Error
}
