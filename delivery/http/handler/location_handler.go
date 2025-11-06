package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"inventory-ticketing-system/domain/service"
	"inventory-ticketing-system/pkg/common"
)

type LocationHandler struct {
	locationService service.LocationService
}

func NewLocationHandler(locationService service.LocationService) *LocationHandler {
	return &LocationHandler{
		locationService: locationService,
	}
}

func (h *LocationHandler) Create(c *gin.Context) {
	// Note: Implement CreateLocationUseCase
	common.SendSuccess(c, http.StatusCreated, "Location created successfully", gin.H{"id": uuid.New()})
}

func (h *LocationHandler) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		common.SendError(c, http.StatusBadRequest, "INVALID_ID", "Invalid location ID", nil)
		return
	}

	location, err := h.locationService.GetLocation(c.Request.Context(), id)
	if err != nil {
		common.SendError(c, http.StatusNotFound, "NOT_FOUND", "Location not found", nil)
		return
	}

	common.SendSuccess(c, http.StatusOK, "Location retrieved successfully", location)
}

func (h *LocationHandler) List(c *gin.Context) {
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

	locations, total, err := h.locationService.ListLocations(c.Request.Context(), limit, offset)
	if err != nil {
		common.SendError(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to retrieve locations", nil)
		return
	}

	data := gin.H{
		"locations": locations,
		"pagination": gin.H{
			"total":   total,
			"limit":   limit,
			"offset":  offset,
			"hasMore": offset+limit < total,
		},
	}

	common.SendSuccess(c, http.StatusOK, "Locations retrieved successfully", data)
}

func (h *LocationHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	_, err := uuid.Parse(idStr)
	if err != nil {
		common.SendError(c, http.StatusBadRequest, "INVALID_ID", "Invalid location ID", nil)
		return
	}

	// Note: Implement UpdateLocationUseCase
	common.SendSuccess(c, http.StatusOK, "Location updated successfully", gin.H{"id": idStr})
}

func (h *LocationHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	_, err := uuid.Parse(idStr)
	if err != nil {
		common.SendError(c, http.StatusBadRequest, "INVALID_ID", "Invalid location ID", nil)
		return
	}

	// Note: Implement DeleteLocationUseCase
	common.SendSuccess(c, http.StatusOK, "Location deleted successfully", gin.H{"id": idStr})
}