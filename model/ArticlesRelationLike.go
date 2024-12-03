package model

import "gorm.io/gorm"

type ArticlesRelationLike struct {
	gorm.Model
	ArticleID int64 `gorm:"column:article_id;not null;index:idx_article_id"`
	UserID    int64 `gorm:"column:user_id;not null;index:idx_user_id"`
}
