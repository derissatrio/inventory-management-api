package asset

import (
	"context"

	"github.com/google/uuid"
	assetdto "inventory-ticketing-system/application/dto/asset"
	"inventory-ticketing-system/domain/entity"
	"inventory-ticketing-system/domain/service"
)

type CreateAssetUseCase struct {
	assetService service.AssetService
}

func NewCreateAssetUseCase(assetService service.AssetService) *CreateAssetUseCase {
	return &CreateAssetUseCase{
		assetService: assetService,
	}
}

func (uc *CreateAssetUseCase) Execute(ctx context.Context, req *assetdto.CreateAssetRequest) (*entity.Asset, error) {
	asset := &entity.Asset{
		ID:            uuid.New(),
		UniqueID:      req.UniqueID,
		Name:          req.Name,
		Comment:       req.Comment,
		Detail:        req.Detail,
		Qty:           req.Qty,
		Brand:         req.Brand,
		Type:          req.Type,
		Status:        req.Status,
		Category:      req.Category,
		LocationID:    req.GetLocationID(),
		LocationLabel: req.LocationLabel,
	}

	err := uc.assetService.CreateAsset(ctx, asset)
	if err != nil {
		return nil, err
	}

	return asset, nil
}