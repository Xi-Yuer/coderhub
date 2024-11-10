package repository

import (
	"coderhub/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUserByName(name string) (*model.User, error)
	GetUserByID(id int64) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id int64) error
}

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		DB: db,
	}
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func (r *UserRepositoryImpl) CreateUser(user *model.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepositoryImpl) GetUserByName(name string) (*model.User, error) {
	var user model.User
	if err := r.DB.Where("user_name = ?", name).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) GetUserByID(id int64) (*model.User, error) {
	var user model.User
	if err := r.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) UpdateUser(user *model.User) error {
	return r.DB.Model(user).Updates(user).Error
}

func (r *UserRepositoryImpl) DeleteUser(id int64) error {
	return r.DB.Delete(&model.User{}, id).Error
}
