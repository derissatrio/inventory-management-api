package location

import "inventory-ticketing-system/domain/entity"

type LocationResponse struct {
	*entity.Location
}

type LocationListResponse struct {
	Locations []LocationResponse `json:"locations"`
	Total     int                `json:"total"`
}