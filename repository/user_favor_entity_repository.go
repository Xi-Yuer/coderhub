package repository

import (
	"coderhub/model"
	"coderhub/shared/storage"
	"context"
	"gorm.io/gorm"
)

type UserFavorEntityRepository interface {
	Create(ctx context.Context, userFavorEntity *model.UserFavor) error
	Delete(ctx context.Context, userFavorEntity *model.UserFavor) error
	GetList(ctx context.Context, userFavorEntity *model.UserFavor, page, pageSize int) ([]*model.UserFavor, int64, error)
	GetFavorByEntityID(ctx context.Context, id int64, entityType string, id2 int64) (*model.UserFavor, error)
}

type UserFavorEntityRepositoryImpl struct {
	DB    *gorm.DB
	Redis storage.RedisDB
}

func NewUserFavorEntityRepository(db *gorm.DB, rdb storage.RedisDB) *UserFavorEntityRepositoryImpl {
	return &UserFavorEntityRepositoryImpl{DB: db, Redis: rdb}
}

func (r *UserFavorEntityRepositoryImpl) Create(ctx context.Context, userFavorEntity *model.UserFavor) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		// 更新收藏夹数量
		update := tx.Model(&model.UserFavorFolder{}).Where("id = ?", userFavorEntity.FavorFoldId).Update("favor_num", gorm.Expr("favor_num + ?", 1))
		if update.Error != nil {
			return update.Error
		}
		// 创建收藏记录
		return tx.WithContext(ctx).Create(userFavorEntity).Error
	})
}

func (r *UserFavorEntityRepositoryImpl) Delete(ctx context.Context, userFavorEntity *model.UserFavor) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&model.UserFavorFolder{}).Where("id = ?", userFavorEntity.FavorFoldId).Update("favor_num", gorm.Expr("favor_num - ?", 1)).Error
		if err != nil {
			return err
		}
		return r.DB.WithContext(ctx).Where("id = ? user_id = ? AND favor_fold_id = ? AND entity_id = ? AND entity_type = ?", userFavorEntity.ID, userFavorEntity.UserId, userFavorEntity.FavorFoldId, userFavorEntity.EntityId, userFavorEntity.EntityType).Delete(userFavorEntity).Error
	})
}

func (r *UserFavorEntityRepositoryImpl) GetFavorByEntityID(ctx context.Context, id int64, entityType string, id2 int64) (*model.UserFavor, error) {
	var userFavorEntity model.UserFavor
	r.DB.WithContext(ctx).Where("user_id = ? AND favor_fold_id = ? AND entity_id = ? AND entity_type = ?", id, id2, id2, entityType).First(&userFavorEntity)
	return &userFavorEntity, nil
}

func (r *UserFavorEntityRepositoryImpl) GetList(ctx context.Context, userFavorEntity *model.UserFavor, page, pageSize int) ([]*model.UserFavor, int64, error) {
	var userFavorEntities []*model.UserFavor
	var total int64
	err := r.DB.WithContext(ctx).Model(userFavorEntity).Where("user_id = ? AND favor_fold_id = ? AND entity_type = ?", userFavorEntity.UserId, userFavorEntity.FavorFoldId, userFavorEntity.EntityType).Limit(pageSize).Offset((page - 1) * pageSize).Count(&total).Find(&userFavorEntities).Error
	return userFavorEntities, total, err
}
