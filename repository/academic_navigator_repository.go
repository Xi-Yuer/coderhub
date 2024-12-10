package repository

import (
	"coderhub/model"

	"gorm.io/gorm"
)

type AcademicNavigatorRepository interface {
	AddAcademicNavigator(academicNavigator *model.AcademicNavigator) error
	GetAcademicNavigator(academicNavigator *model.AcademicNavigator) ([]*model.AcademicNavigator, int64, error)
	GetAcademicNavigatorByID(ID int64) (*model.AcademicNavigator, error)
	DeleteAcademicNavigator(ID int64) error
}

type AcademicNavigatorRepositoryImpl struct {
	DB *gorm.DB
}

func NewAcademicNavigatorRepositoryImpl(db *gorm.DB) *AcademicNavigatorRepositoryImpl {
	return &AcademicNavigatorRepositoryImpl{
		DB: db,
	}
}

func (r *AcademicNavigatorRepositoryImpl) AddAcademicNavigator(academicNavigator *model.AcademicNavigator) error {
	return r.DB.Create(academicNavigator).Error
}

func (r *AcademicNavigatorRepositoryImpl) GetAcademicNavigator(academicNavigator *model.AcademicNavigator) ([]*model.AcademicNavigator, int64, error) {
	var total int64
	var academicNavigators []*model.AcademicNavigator
	err := r.DB.Model(&model.AcademicNavigator{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = r.DB.Where("user_id = ?", academicNavigator.UserId).Find(&academicNavigators).Error
	return academicNavigators, total, err
}

func (r *AcademicNavigatorRepositoryImpl) GetAcademicNavigatorByID(ID int64) (*model.AcademicNavigator, error) {
	var academicNavigator model.AcademicNavigator
	err := r.DB.First(&academicNavigator, ID).Error
	return &academicNavigator, err
}

func (r *AcademicNavigatorRepositoryImpl) DeleteAcademicNavigator(ID int64) error {
	return r.DB.Delete(&model.AcademicNavigator{}, ID).Error
}
