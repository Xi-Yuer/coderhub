package repository

import (
	"coderhub/model"
	"context"

	"gorm.io/gorm"
)

type ImageRelationRepository interface {
	Create(ctx context.Context, imageRelation *model.ImageRelation) error
	BatchCreate(ctx context.Context, imageRelations []*model.ImageRelation) error
	Delete(ctx context.Context, imageID int64, entityID int64, entityType string) error
	ListByEntityID(ctx context.Context, entityID int64, entityType string) ([]model.ImageRelation, error)
	ListByImageID(ctx context.Context, imageID int64) ([]model.ImageRelation, error)
}

type imageRelationRepository struct {
	DB *gorm.DB
}

func NewImageRelationRepository(db *gorm.DB) ImageRelationRepository {
	return &imageRelationRepository{DB: db}
}

func (r *imageRelationRepository) Create(ctx context.Context, imageRelation *model.ImageRelation) error {
	return r.DB.WithContext(ctx).Create(imageRelation).Error
}

func (r *imageRelationRepository) BatchCreate(ctx context.Context, imageRelations []*model.ImageRelation) error {
	return r.DB.WithContext(ctx).Create(&imageRelations).Error
}

func (r *imageRelationRepository) Delete(ctx context.Context, imageID int64, entityID int64, entityType string) error {
	return r.DB.WithContext(ctx).Where("image_id = ? AND entity_id = ? AND entity_type = ?", imageID, entityID, entityType).Delete(&model.ImageRelation{}).Error
}

func (r *imageRelationRepository) ListByEntityID(ctx context.Context, entityID int64, entityType string) ([]model.ImageRelation, error) {
	var imageRelations []model.ImageRelation
	return imageRelations, r.DB.WithContext(ctx).Where("entity_id = ? AND entity_type = ?", entityID, entityType).Find(&imageRelations).Error
}

func (r *imageRelationRepository) ListByImageID(ctx context.Context, imageID int64) ([]model.ImageRelation, error) {
	var imageRelations []model.ImageRelation
	return imageRelations, r.DB.WithContext(ctx).Where("image_id = ?", imageID).Find(&imageRelations).Error
}