package asset

import (
	"context"

	assetdto "inventory-ticketing-system/application/dto/asset"
	"inventory-ticketing-system/domain/service"
)

type ListAssetsUseCase struct {
	assetService service.AssetService
}

func NewListAssetsUseCase(assetService service.AssetService) *ListAssetsUseCase {
	return &ListAssetsUseCase{
		assetService: assetService,
	}
}

func (uc *ListAssetsUseCase) Execute(ctx context.Context, limit, offset int, filters map[string]interface{}) (*assetdto.AssetListResponse, error) {
	assets, total, err := uc.assetService.ListAssets(ctx, limit, offset, filters)
	if err != nil {
		return nil, err
	}

	assetResponses := make([]assetdto.AssetResponse, len(assets))
	for i, asset := range assets {
		assetResponses[i] = assetdto.AssetResponse{Asset: asset}
	}

	return &assetdto.AssetListResponse{
		Assets: assetResponses,
		Total:  total,
	}, nil
}