package repository

import (
	"coderhub/model"
	"context"
	"errors"

	"gorm.io/gorm"
)

// ImageRepository 图片仓储接口
type ImageRepository interface {
	// Create 创建图片元数据记录
	Create(ctx context.Context, image *model.Image) error
	// BatchCreate 批量创建图片元数据记录
	BatchCreate(ctx context.Context, images []*model.Image) error
	// GetByID 获取图片元数据
	GetByID(ctx context.Context, id int64) (*model.Image, error)
	// Delete 删除图片元数据（软删除）
	Delete(ctx context.Context, id int64) error
	// ListByEntityID 获取关联实体的图片列表
	ListByEntityID(ctx context.Context, entityID int64, entityType string) ([]model.Image, int64, error)
}

var (
	ErrImageNotFound = errors.New("图片未找到")
)

type imageRepository struct {
	DB *gorm.DB
}

func NewImageRepository(db *gorm.DB) ImageRepository {
	return &imageRepository{
		DB: db,
	}
}

func (r *imageRepository) Create(ctx context.Context, image *model.Image) error {
	return r.DB.WithContext(ctx).Create(image).Error
}

func (r *imageRepository) BatchCreate(ctx context.Context, images []*model.Image) error {
	if len(images) == 0 {
		return nil
	}
	return r.DB.WithContext(ctx).Create(&images).Error
}

func (r *imageRepository) GetByID(ctx context.Context, id int64) (*model.Image, error) {
	var image model.Image
	err := r.DB.WithContext(ctx).First(&image, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrImageNotFound
		}
		return nil, err
	}
	return &image, nil
}

func (r *imageRepository) Delete(ctx context.Context, id int64) error {
	result := r.DB.WithContext(ctx).Delete(&model.Image{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return ErrImageNotFound
	}
	return result.Error
}

func (r *imageRepository) ListByEntityID(ctx context.Context, entityID int64, entityType string) ([]model.Image, int64, error) {
	var images []model.Image
	var total int64
	err := r.DB.WithContext(ctx).
		Select("images.*").
		Joins("RIGHT JOIN image_relations ON images.id = image_relations.image_id").
		Where("image_relations.entity_id = ? AND image_relations.entity_type = ?", entityID, entityType).
		Order("image_relations.sort ASC").
		Find(&images).Error
	if err != nil {
		return nil, 0, err
	}
	return images, total, nil
}
