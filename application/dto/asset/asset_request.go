package asset

import "github.com/google/uuid"

type CreateAssetRequest struct {
	UniqueID      string     `json:"uniqueId" binding:"required"`
	Name          string     `json:"name" binding:"required"`
	Comment       string     `json:"comment"`
	Detail        string     `json:"detail"`
	Qty           int        `json:"qty" binding:"min=1"`
	Brand         string     `json:"brand"`
	Type          string     `json:"type" binding:"required,oneof=it non_it"`
	Status        string     `json:"status" binding:"omitempty,oneof=available booked broken repair"`
	Category      string     `json:"category"`
	LocationID    string     `json:"locationId"` // Accept string, will be validated and converted to UUID
	LocationLabel string     `json:"locationLabel"`
}

// GetLocationID returns the LocationID as UUID or nil if empty
func (r *CreateAssetRequest) GetLocationID() *uuid.UUID {
	if r.LocationID == "" || r.LocationID == "null" || r.LocationID == "undefined" {
		return nil
	}
	parsed, err := uuid.Parse(r.LocationID)
	if err != nil {
		return nil
	}
	return &parsed
}

type UpdateAssetRequest struct {
	Name          string `json:"name,omitempty"`
	Comment       string `json:"comment,omitempty"`
	Detail        string `json:"detail,omitempty"`
	Qty           *int   `json:"qty,omitempty"`
	Brand         string `json:"brand,omitempty"`
	Type          string `json:"type,omitempty" binding:"omitempty,oneof=it non_it"`
	Status        string `json:"status,omitempty" binding:"omitempty,oneof=available booked broken repair"`
	Category      string `json:"category,omitempty"`
	LocationID    string `json:"locationId,omitempty"` // Accept string, will be validated and converted to UUID
	LocationLabel string `json:"locationLabel,omitempty"`
}

// GetLocationID returns the LocationID as UUID or nil if empty
func (r *UpdateAssetRequest) GetLocationID() *uuid.UUID {
	if r.LocationID == "" || r.LocationID == "null" || r.LocationID == "undefined" {
		return nil
	}
	parsed, err := uuid.Parse(r.LocationID)
	if err != nil {
		return nil
	}
	return &parsed
}