package model

import (
	"gorm.io/gorm"
)

type UserFollow struct {
	gorm.Model
	FollowerID int64 `gorm:"index"` // 关注者ID
	FollowedID int64 `gorm:"index"` // 被关注者ID

	// 联合唯一键
	UNIQUEKEY UserFollowUniqueKey `gorm:"uniqueIndex:idx_user_follow_unique,composite:follower_followed"`
}

type UserFollowUniqueKey struct {
	FollowerID int64
	FollowedID int64
}
