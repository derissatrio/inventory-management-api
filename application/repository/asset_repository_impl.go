package repository

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"inventory-ticketing-system/domain/entity"
	"inventory-ticketing-system/domain/repository"
)

type AssetRepositoryImpl struct {
	db *gorm.DB
}

func NewAssetRepository(db *gorm.DB) repository.AssetRepository {
	return &AssetRepositoryImpl{
		db: db,
	}
}

func (r *AssetRepositoryImpl) Create(ctx context.Context, asset *entity.Asset) error {
	return r.db.WithContext(ctx).Create(asset).Error
}

func (r *AssetRepositoryImpl) GetByID(ctx context.Context, id uuid.UUID) (*entity.Asset, error) {
	var asset entity.Asset
	err := r.db.WithContext(ctx).Preload("Location").Where("id = ?", id).First(&asset).Error
	if err != nil {
		return nil, err
	}
	return &asset, nil
}

func (r *AssetRepositoryImpl) Update(ctx context.Context, asset *entity.Asset) error {
	return r.db.WithContext(ctx).Save(asset).Error
}

func (r *AssetRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&entity.Asset{}, "id = ?", id).Error
}

func (r *AssetRepositoryImpl) List(ctx context.Context, limit, offset int, filters map[string]interface{}) ([]*entity.Asset, int, error) {
	var assets []*entity.Asset
	var total int64

	query := r.db.WithContext(ctx).Model(&entity.Asset{}).Preload("Location")

	for key, value := range filters {
		switch key {
		case "type":
			query = query.Where("type = ?", value)
		case "status":
			query = query.Where("status = ?", value)
		case "category":
			query = query.Where("category ILIKE ?", "%"+value.(string)+"%")
		case "brand":
			query = query.Where("brand ILIKE ?", "%"+value.(string)+"%")
		}
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Limit(limit).Offset(offset).Find(&assets).Error
	if err != nil {
		return nil, 0, err
	}

	return assets, int(total), nil
}

func (r *AssetRepositoryImpl) GetByUniqueID(ctx context.Context, uniqueID string) (*entity.Asset, error) {
	var asset entity.Asset
	err := r.db.WithContext(ctx).Preload("Location").Where("unique_id = ?", uniqueID).First(&asset).Error
	if err != nil {
		return nil, err
	}
	return &asset, nil
}