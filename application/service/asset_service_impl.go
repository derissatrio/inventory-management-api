package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"inventory-ticketing-system/domain/entity"
	"inventory-ticketing-system/domain/repository"
	"inventory-ticketing-system/domain/service"
)

type AssetServiceImpl struct {
	assetRepo repository.AssetRepository
}

func NewAssetService(assetRepo repository.AssetRepository) service.AssetService {
	return &AssetServiceImpl{
		assetRepo: assetRepo,
	}
}

func (s *AssetServiceImpl) CreateAsset(ctx context.Context, asset *entity.Asset) error {
	if asset.Type == "" {
		asset.Type = "it" // Default type
	}
	if asset.Status == "" {
		asset.Status = "available" // Default status
	}
	if asset.Qty <= 0 {
		asset.Qty = 1 // Default quantity
	}

	existingAsset, err := s.assetRepo.GetByUniqueID(ctx, asset.UniqueID)
	if err == nil && existingAsset != nil {
		return errors.New("asset with this unique ID already exists")
	}

	return s.assetRepo.Create(ctx, asset)
}

func (s *AssetServiceImpl) GetAsset(ctx context.Context, id uuid.UUID) (*entity.Asset, error) {
	return s.assetRepo.GetByID(ctx, id)
}

func (s *AssetServiceImpl) UpdateAsset(ctx context.Context, id uuid.UUID, asset *entity.Asset) error {
	existingAsset, err := s.assetRepo.GetByID(ctx, id)
	if err != nil {
		return errors.New("asset not found")
	}

	asset.ID = id
	asset.CreatedAt = existingAsset.CreatedAt
	asset.UpdatedAt = time.Now()

	return s.assetRepo.Update(ctx, asset)
}

func (s *AssetServiceImpl) DeleteAsset(ctx context.Context, id uuid.UUID) error {
	_, err := s.assetRepo.GetByID(ctx, id)
	if err != nil {
		return errors.New("asset not found")
	}

	return s.assetRepo.Delete(ctx, id)
}

func (s *AssetServiceImpl) ListAssets(ctx context.Context, limit, offset int, filters map[string]interface{}) ([]*entity.Asset, int, error) {
	return s.assetRepo.List(ctx, limit, offset, filters)
}

func (s *AssetServiceImpl) UpdateAssetStatus(ctx context.Context, id uuid.UUID, status string) error {
	asset, err := s.assetRepo.GetByID(ctx, id)
	if err != nil {
		return errors.New("asset not found")
	}

	asset.Status = status
	return s.assetRepo.Update(ctx, asset)
}

func (s *AssetServiceImpl) DecreaseAssetQuantity(ctx context.Context, id uuid.UUID, qty int) error {
	asset, err := s.assetRepo.GetByID(ctx, id)
	if err != nil {
		return errors.New("asset not found")
	}

	if asset.Qty < qty {
		return errors.New("insufficient quantity")
	}

	asset.Qty -= qty
	return s.assetRepo.Update(ctx, asset)
}

func (s *AssetServiceImpl) IncreaseAssetQuantity(ctx context.Context, id uuid.UUID, qty int) error {
	asset, err := s.assetRepo.GetByID(ctx, id)
	if err != nil {
		return errors.New("asset not found")
	}

	asset.Qty += qty
	return s.assetRepo.Update(ctx, asset)
}