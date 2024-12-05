package repository

import (
	"coderhub/model"
	"coderhub/shared/CacheDB"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

type UserFollowRepository interface {
	// 创建用户关注关系
	CreateUserFollow(userFollow *model.UserFollow) error
	// 删除用户关注关系
	DeleteUserFollow(userFollow *model.UserFollow) error
	// 查询用户关注的所有用户
	GetUserFollows(followerID int64, page int32, pageSize int32) ([]*model.UserFollow, error)
	// 批量查询用户关注的所有用户
	BatchGetUserFollows(followerID int64, page int32, pageSize int32) ([]*model.UserFollow, error)
	// 查询某用户的粉丝列表
	GetUserFans(followedID int64, page int32, pageSize int32) ([]*model.UserFollow, error)
	// 批量查询某用户的粉丝列表
	BatchGetUserFans(followedID int64, page int32, pageSize int32) ([]*model.UserFollow, error)
	// 判断两个用户是否存在关注关系
	IsUserFollowed(followerID int64, followedID int64) (bool, error)
	// 查询互相关注的用户
	GetMutualFollows(userID int64, page int32, pageSize int32) ([]*model.UserFollow, error)
}

func NewUserFollowRepositoryImpl(db *gorm.DB, rdb CacheDB.RedisDB) *UserFollowRepositoryImpl {
	return &UserFollowRepositoryImpl{
		DB:    db,
		Redis: rdb,
	}
}

type UserFollowRepositoryImpl struct {
	DB    *gorm.DB
	Redis CacheDB.RedisDB
}

// 创建用户关注关系
func (r *UserFollowRepositoryImpl) CreateUserFollow(userFollow *model.UserFollow) error {
	return r.DB.Create(userFollow).Error
}

// 删除用户关注关系
func (r *UserFollowRepositoryImpl) DeleteUserFollow(userFollow *model.UserFollow) error {
	return r.DB.Delete(userFollow).Where("follower_id = ? AND followed_id = ?", userFollow.FollowerID, userFollow.FollowedID).Error
}

// 查询用户关注的所有用户
func (r *UserFollowRepositoryImpl) GetUserFollows(followerID int64, page int32, pageSize int32) ([]*model.UserFollow, error) {
	var userFollows []*model.UserFollow
	return userFollows, r.DB.Where("follower_id = ?", followerID).Offset((int(page) - 1) * int(pageSize)).Limit(int(pageSize)).Find(&userFollows).Error
}

// 批量查询用户关注的所有用户
func (r *UserFollowRepositoryImpl) BatchGetUserFollows(followerID int64, page int32, pageSize int32) ([]*model.UserFollow, error) {
	var userFollows []*model.UserFollow
	return userFollows, r.DB.Where("follower_id = ?", followerID).Offset((int(page) - 1) * int(pageSize)).Limit(int(pageSize)).Find(&userFollows).Error
}

// 查询某用户的粉丝列表(热点用户（如明星）可能有数千万粉丝，查询时可能导致数据库压力大。可以将热点用户的粉丝列表缓存到 Redis 中。)
func (r *UserFollowRepositoryImpl) GetUserFans(followedID int64, page int32, pageSize int32) ([]*model.UserFollow, error) {
	var userFollows []*model.UserFollow
	// 先从 Redis 中查询
	cacheKey := fmt.Sprintf("user_fans:%d", followedID)
	cacheData, err := r.Redis.Get(cacheKey)
	if err == nil {
		// 反序列化
		err = json.Unmarshal([]byte(cacheData), &userFollows)
		if err == nil {
			return userFollows, nil
		}
	}
	// 如果 Redis 中没有数据，则从数据库中查询
	err = r.DB.Where("followed_id = ?", followedID).Offset((int(page) - 1) * int(pageSize)).Limit(int(pageSize)).Find(&userFollows).Error
	if err == nil {
		// 将查询结果序列化并缓存到 Redis 中
		bytes, _ := json.Marshal(userFollows)
		cacheData = string(bytes)
		r.Redis.Set(cacheKey, cacheData)
	}
	return userFollows, err
}

// 批量查询某用户的粉丝列表
func (r *UserFollowRepositoryImpl) BatchGetUserFans(followedID int64, page int32, pageSize int32) ([]*model.UserFollow, error) {
	var userFollows []*model.UserFollow
	return userFollows, r.DB.Where("followed_id = ?", followedID).Offset((int(page) - 1) * int(pageSize)).Limit(int(pageSize)).Find(&userFollows).Error
}

// 判断两个用户是否存在关注关系
func (r *UserFollowRepositoryImpl) IsUserFollowed(followerID int64, followedID int64) (bool, error) {
	var userFollow model.UserFollow
	return r.DB.Where("follower_id = ? AND followed_id = ?", followerID, followedID).First(&userFollow).RowsAffected > 0, nil
}

// 查询互相关注的用户
func (r *UserFollowRepositoryImpl) GetMutualFollows(userID int64, page int32, pageSize int32) ([]*model.UserFollow, error) {
	var userFollows []*model.UserFollow
	return userFollows, r.DB.Where("follower_id = ? AND followed_id = ?", userID, userID).Offset((int(page) - 1) * int(pageSize)).Limit(int(pageSize)).Find(&userFollows).Error
}
