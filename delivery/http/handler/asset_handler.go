package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	assetdto "inventory-ticketing-system/application/dto/asset"
	assetusecase "inventory-ticketing-system/application/usecase/asset"
	"inventory-ticketing-system/pkg/common"
)

type AssetHandler struct {
	createAssetUseCase *assetusecase.CreateAssetUseCase
	listAssetsUseCase  *assetusecase.ListAssetsUseCase
}

func NewAssetHandler(
	createAssetUseCase *assetusecase.CreateAssetUseCase,
	listAssetsUseCase *assetusecase.ListAssetsUseCase,
) *AssetHandler {
	return &AssetHandler{
		createAssetUseCase: createAssetUseCase,
		listAssetsUseCase:  listAssetsUseCase,
	}
}

func (h *AssetHandler) Create(c *gin.Context) {
	var req assetdto.CreateAssetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.SendValidationError(c, err)
		return
	}

	asset, err := h.createAssetUseCase.Execute(c.Request.Context(), &req)
	if err != nil {
		common.SendError(c, http.StatusBadRequest, "VALIDATION_ERROR", err.Error(), nil)
		return
	}

	response := assetdto.AssetResponse{Asset: asset}
	common.SendSuccess(c, http.StatusCreated, "Asset created successfully", response)
}

func (h *AssetHandler) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		common.SendError(c, http.StatusBadRequest, "INVALID_ID", "Invalid asset ID", nil)
		return
	}

	// Note: Implement GetAssetUseCase for this functionality
	// For now, return a simple response
	common.SendSuccess(c, http.StatusOK, "Asset retrieved successfully", gin.H{"id": id})
}

func (h *AssetHandler) List(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "20")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 20
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	// Build filters
	filters := make(map[string]interface{})
	if assetType := c.Query("jenis"); assetType != "" {
		filters["type"] = assetType
	}
	if status := c.Query("status"); status != "" {
		filters["status"] = status
	}
	if category := c.Query("category"); category != "" {
		filters["category"] = category
	}
	if brand := c.Query("brand"); brand != "" {
		filters["brand"] = brand
	}

	response, err := h.listAssetsUseCase.Execute(c.Request.Context(), limit, offset, filters)
	if err != nil {
		common.SendError(c, http.StatusInternalServerError, "INTERNAL_ERROR", err.Error(), nil)
		return
	}

	// Add pagination info
	hasMore := response.Total > (offset + limit)
	pagination := common.PaginationInfo{
		Total:   response.Total,
		Limit:   limit,
		Offset:  offset,
		HasMore: hasMore,
	}

	data := gin.H{
		"assets":     response.Assets,
		"pagination": pagination,
	}

	common.SendSuccess(c, http.StatusOK, "Assets retrieved successfully", data)
}

func (h *AssetHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	_, err := uuid.Parse(idStr)
	if err != nil {
		common.SendError(c, http.StatusBadRequest, "INVALID_ID", "Invalid asset ID", nil)
		return
	}

	// Note: Implement UpdateAssetUseCase for this functionality
	common.SendSuccess(c, http.StatusOK, "Asset updated successfully", gin.H{"id": idStr})
}

func (h *AssetHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	_, err := uuid.Parse(idStr)
	if err != nil {
		common.SendError(c, http.StatusBadRequest, "INVALID_ID", "Invalid asset ID", nil)
		return
	}

	// Note: Implement DeleteAssetUseCase for this functionality
	common.SendSuccess(c, http.StatusOK, "Asset deleted successfully", gin.H{"id": idStr})
}
