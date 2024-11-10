package repository

import (
	"coderhub/model"
	"coderhub/shared/cacheDB"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUserByName(name string) (*model.User, error)
	GetUserByID(id int64) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id int64) error
}

func NewUserRepositoryImpl(db *gorm.DB, rdb cacheDB.RedisDB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		DB:    db,
		Redis: rdb,
	}
}

type UserRepositoryImpl struct {
	DB    *gorm.DB
	Redis cacheDB.RedisDB
}

func (r *UserRepositoryImpl) CreateUser(user *model.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepositoryImpl) GetUserByName(name string) (*model.User, error) {
	var user model.User
	// 先从 Redis 获取
	key := fmt.Sprintf("user:%s", name)
	data, err := r.Redis.Get(key)
	if err == nil {
		if err := json.Unmarshal([]byte(data), &user); err == nil {
			return &user, nil
		}
	}
	if err := r.DB.Where("user_name = ?", name).First(&user).Error; err != nil {
		return nil, err
	}
	// 写入 Redis 缓存
	if data, err := json.Marshal(user); err == nil {
		err := r.Redis.Set(key, string(data))
		if err != nil {
			return nil, err
		}
	}
	return &user, nil
}

func (r *UserRepositoryImpl) GetUserByID(id int64) (*model.User, error) {
	// 先从 Redis 获取
	key := fmt.Sprintf("user:%d", id)
	data, err := r.Redis.Get(key)
	if err == nil {
		var user model.User
		if err := json.Unmarshal([]byte(data), &user); err == nil {
			return &user, nil
		}
	}

	// Redis 未命中,从数据库获取
	var user model.User
	if err := r.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	// 写入 Redis 缓存
	if data, err := json.Marshal(user); err == nil {
		err := r.Redis.Set(key, string(data))
		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (r *UserRepositoryImpl) UpdateUser(user *model.User) error {
	return r.DB.Model(user).Updates(user).Error
}

func (r *UserRepositoryImpl) DeleteUser(id int64) error {
	return r.DB.Delete(&model.User{}, id).Error
}
