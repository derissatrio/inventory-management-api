package repository

import (
	"context"

	"github.com/google/uuid"
	"inventory-ticketing-system/domain/entity"
)

type AssetRepository interface {
	Create(ctx context.Context, asset *entity.Asset) error
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Asset, error)
	Update(ctx context.Context, asset *entity.Asset) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, limit, offset int, filters map[string]interface{}) ([]*entity.Asset, int, error)
	GetByUniqueID(ctx context.Context, uniqueID string) (*entity.Asset, error)
}