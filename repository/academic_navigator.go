package repository

import (
	"coderhub/model"

	"gorm.io/gorm"
)

type AcademicNavigatorRepository interface {
	AddAcademicNavigator(academicNavigator *model.AcademicNavigator) error
	GetAcademicNavigator(academicNavigator *model.AcademicNavigator) error
	DeleteAcademicNavigator(academicNavigator *model.AcademicNavigator) error
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

func (r *AcademicNavigatorRepositoryImpl) GetAcademicNavigator(academicNavigator *model.AcademicNavigator) error {
	return r.DB.Where("user_id = ?", academicNavigator.UserId).Find(academicNavigator).Error
}

func (r *AcademicNavigatorRepositoryImpl) DeleteAcademicNavigator(academicNavigator *model.AcademicNavigator) error {
	return r.DB.Delete(academicNavigator).Error
}
