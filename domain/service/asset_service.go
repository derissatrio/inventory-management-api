package service

import (
	"context"

	"github.com/google/uuid"
	"inventory-ticketing-system/domain/entity"
)

type AssetService interface {
	CreateAsset(ctx context.Context, asset *entity.Asset) error
	GetAsset(ctx context.Context, id uuid.UUID) (*entity.Asset, error)
	UpdateAsset(ctx context.Context, id uuid.UUID, asset *entity.Asset) error
	DeleteAsset(ctx context.Context, id uuid.UUID) error
	ListAssets(ctx context.Context, limit, offset int, filters map[string]interface{}) ([]*entity.Asset, int, error)
	UpdateAssetStatus(ctx context.Context, id uuid.UUID, status string) error
	DecreaseAssetQuantity(ctx context.Context, id uuid.UUID, qty int) error
	IncreaseAssetQuantity(ctx context.Context, id uuid.UUID, qty int) error
}