package repository

import (
	"coderhub/model"
	"coderhub/shared/CacheDB"
	"context"

	"gorm.io/gorm"
)

type ArticlesRelationLikeRepository interface {
	Create(ctx context.Context, articleRelationLike *model.ArticlesRelationLike) error
	Delete(ctx context.Context, articleRelationLike *model.ArticlesRelationLike) error
	Get(ctx context.Context, articleRelationLike *model.ArticlesRelationLike) bool
	List(ctx context.Context, articleID int64) (int64, error)
}

type articlesRelationLikeRepository struct {
	DB    *gorm.DB
	Redis CacheDB.RedisDB
}

func NewArticlesRelationLikeRepository(db *gorm.DB, redis CacheDB.RedisDB) ArticlesRelationLikeRepository {
	return &articlesRelationLikeRepository{
		DB:    db,
		Redis: redis,
	}
}

func (r *articlesRelationLikeRepository) Create(ctx context.Context, articleRelationLike *model.ArticlesRelationLike) error {
	return r.DB.Create(articleRelationLike).Error
}

func (r *articlesRelationLikeRepository) Delete(ctx context.Context, articleRelationLike *model.ArticlesRelationLike) error {
	return r.DB.Delete(articleRelationLike, "article_id = ? AND user_id = ?", articleRelationLike.ArticleID, articleRelationLike.UserID).Error
}

func (r *articlesRelationLikeRepository) Get(ctx context.Context, articleRelationLike *model.ArticlesRelationLike) bool {
	var count int64
	r.DB.Model(articleRelationLike).Where("article_id = ? AND user_id = ?", articleRelationLike.ArticleID, articleRelationLike.UserID).Count(&count)
	return count > 0
}

func (r *articlesRelationLikeRepository) List(ctx context.Context, articleID int64) (int64, error) {
	var articlesRelationLikesCount int64
	r.DB.Model(&model.ArticlesRelationLike{}).Where("article_id = ?", articleID).Count(&articlesRelationLikesCount)
	return articlesRelationLikesCount, nil
}
