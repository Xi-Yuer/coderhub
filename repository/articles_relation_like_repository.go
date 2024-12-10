package repository

import (
	"coderhub/model"
	"coderhub/shared/storage"
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
	Redis storage.RedisDB
}

func NewArticlesRelationLikeRepository(db *gorm.DB, redis storage.RedisDB) ArticlesRelationLikeRepository {
	return &articlesRelationLikeRepository{
		DB:    db,
		Redis: redis,
	}
}

// 创建文章点赞
func (r *articlesRelationLikeRepository) Create(ctx context.Context, articleRelationLike *model.ArticlesRelationLike) error {
	return r.DB.Create(articleRelationLike).Error
}

// 删除文章点赞
func (r *articlesRelationLikeRepository) Delete(ctx context.Context, articleRelationLike *model.ArticlesRelationLike) error {
	return r.DB.Delete(articleRelationLike, "article_id = ? AND user_id = ?", articleRelationLike.ArticleID, articleRelationLike.UserID).Error
}

// 获取文章是否被用户点赞
func (r *articlesRelationLikeRepository) Get(ctx context.Context, articleRelationLike *model.ArticlesRelationLike) bool {
	var count int64
	r.DB.Model(articleRelationLike).Where("article_id = ? AND user_id = ?", articleRelationLike.ArticleID, articleRelationLike.UserID).Count(&count)
	return count > 0
}

// 获取文章点赞数
func (r *articlesRelationLikeRepository) List(ctx context.Context, articleID int64) (int64, error) {
	var articlesRelationLikesCount int64
	r.DB.Model(&model.ArticlesRelationLike{}).Where("article_id = ?", articleID).Count(&articlesRelationLikesCount)
	return articlesRelationLikesCount, nil
}
