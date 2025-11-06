package asset

import "inventory-ticketing-system/domain/entity"

type AssetResponse struct {
	*entity.Asset
}

type AssetListResponse struct {
	Assets []AssetResponse `json:"assets"`
	Total  int             `json:"total"`
}

type PaginationInfo struct {
	Limit  int  `json:"limit"`
	Offset int  `json:"offset"`
	HasMore bool `json:"hasMore"`
}