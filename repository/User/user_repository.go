package repository

import (
	"coderhub/model"
	"coderhub/shared/CacheDB"
	"encoding/json"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUserByName(name string) (*model.User, error)
	GetUserByID(id int64) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id int64) error
}

func NewUserRepositoryImpl(db *gorm.DB, rdb CacheDB.RedisDB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		DB:    db,
		Redis: rdb,
	}
}

type UserRepositoryImpl struct {
	DB    *gorm.DB
	Redis CacheDB.RedisDB
}

func (r *UserRepositoryImpl) CreateUser(user *model.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepositoryImpl) GetUserByName(name string) (*model.User, error) {
	var user model.User
	key := user.CacheKeyByName(name)

	// 尝试从缓存获取
	if cached, err := r.getCache(key); err == nil {
		return cached, nil
	}

	// 从数据库获取
	if err := r.DB.Where("user_name = ?", name).First(&user).Error; err != nil {
		return nil, err
	}

	// 设置缓存
	if err := r.setCache(key, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryImpl) GetUserByID(id int64) (*model.User, error) {
	var user model.User
	key := user.CacheKeyByID(id)

	// 尝试从缓存获取
	if cached, err := r.getCache(key); err == nil {
		return cached, nil
	}

	// 从数据库获取
	if err := r.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	// 设置缓存
	if err := r.setCache(key, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryImpl) UpdateUser(user *model.User) error {
	// 更新 Redis 缓存
	key := user.CacheKeyByID(user.ID)
	if err := r.Redis.Del(key); err != nil {
		return err
	}
	return r.DB.Model(user).Where("id = ?", user.ID).Updates(user).Error
}

func (r *UserRepositoryImpl) DeleteUser(id int64) error {
	var user model.User
	UserInfo, err := r.GetUserByID(id)
	if err != nil {
		return err
	}
	// 先从 Redis 删除
	CacheIdKey := user.CacheKeyByID(id)
	CacheUserNameIdKey := user.CacheKeyByName(UserInfo.UserName)
	// 根据 id 删除
	if err := r.delCache(CacheIdKey); err != nil {
		return err
	}
	// 根据用户名删除
	if err := r.delCache(CacheUserNameIdKey); err != nil {
		return err
	}
	return r.DB.Where("id = ?", id).Delete(&model.User{}).Error
}

func (r *UserRepositoryImpl) getCache(key string) (*model.User, error) {
	var user model.User
	data, err := r.Redis.Get(key)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(data), &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) setCache(key string, user *model.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return r.Redis.Set(key, string(data))
}

func (r *UserRepositoryImpl) delCache(key string) error {
	return r.Redis.Del(key)
}
