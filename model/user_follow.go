package model

import (
	"gorm.io/gorm"
)

type UserFollow struct {
	gorm.Model
	FollowerID int64 `gorm:"index;uniqueIndex:idx_follower_followed"` // 关注者ID
	FollowedID int64 `gorm:"index;uniqueIndex:idx_follower_followed"` // 被关注者ID

	// 联合唯一索引
	_ struct {
		FollowerID int64
		FollowedID int64
	} `gorm:"uniqueIndex:idx_follower_followed"`
}
