package model

import "gorm.io/gorm"

type CommentRelationLike struct {
	gorm.Model
	CommentID int64 `gorm:"column:comment_id;not null;index:idx_comment_id"`
	UserID    int64 `gorm:"column:user_id;not null;index:idx_user_id"`
}
