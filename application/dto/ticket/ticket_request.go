package ticket

import "github.com/google/uuid"

type CreateTicketRequest struct {
	AssetID  uuid.UUID `json:"assetId" binding:"required"`
	Category string    `json:"kategori" binding:"required"`
	Severity string    `json:"severity" binding:"required,oneof=low medium high critical"`
	Comment  string    `json:"comment" binding:"required"`
}

type UpdateTicketRequest struct {
	Status              string     `json:"status,omitempty" binding:"omitempty,oneof=open in_progress resolved closed"`
	AssignedTo          *uuid.UUID `json:"assignedTo,omitempty"`
	ResolutionComment   string     `json:"resolutionComment,omitempty"`
}

type TicketListRequest struct {
	Limit   int    `form:"limit,default=20" binding:"min=1,max=100"`
	Offset  int    `form:"offset,default=0" binding:"min=0"`
	Status  string `form:"status"`
	AssetID string `form:"assetId"`
	SortBy  string `form:"sortBy,default=created_at"`
	Order   string `form:"order,default=desc" binding:"omitempty,oneof=asc desc"`
}