package repository

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"inventory-ticketing-system/domain/entity"
	"inventory-ticketing-system/domain/repository"
)

type LocationRepositoryImpl struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) repository.LocationRepository {
	return &LocationRepositoryImpl{
		db: db,
	}
}

func (r *LocationRepositoryImpl) Create(ctx context.Context, location *entity.Location) error {
	return r.db.WithContext(ctx).Create(location).Error
}

func (r *LocationRepositoryImpl) GetByID(ctx context.Context, id uuid.UUID) (*entity.Location, error) {
	var location entity.Location
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&location).Error
	if err != nil {
		return nil, err
	}
	return &location, nil
}

func (r *LocationRepositoryImpl) Update(ctx context.Context, location *entity.Location) error {
	return r.db.WithContext(ctx).Save(location).Error
}

func (r *LocationRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&entity.Location{}, "id = ?", id).Error
}

func (r *LocationRepositoryImpl) List(ctx context.Context, limit, offset int) ([]*entity.Location, int, error) {
	var locations []*entity.Location
	var total int64

	query := r.db.WithContext(ctx).Model(&entity.Location{})

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("name ASC").Limit(limit).Offset(offset).Find(&locations).Error
	if err != nil {
		return nil, 0, err
	}

	return locations, int(total), nil
}

func (r *LocationRepositoryImpl) GetByName(ctx context.Context, name string) (*entity.Location, error) {
	var location entity.Location
	err := r.db.WithContext(ctx).Where("name = ?", name).First(&location).Error
	if err != nil {
		return nil, err
	}
	return &location, nil
}