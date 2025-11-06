package ticket

import (
	"time"

	"github.com/google/uuid"
	"inventory-ticketing-system/domain/entity"
)

type TicketResponse struct {
	ID                uuid.UUID        `json:"id"`
	AssetID           uuid.UUID        `json:"assetId"`
	Category          string           `json:"kategori"`
	Severity          string           `json:"severity"`
	Status            string           `json:"status"`
	Comment           string           `json:"comment"`
	Duration          int              `json:"duration"`
	DueDate           time.Time        `json:"dueDate"`
	Reporting         uuid.UUID        `json:"reporting"`
	AssignedTo        *uuid.UUID       `json:"assignedTo"`
	ResolutionComment string           `json:"resolutionComment"`
	CreatedAt         time.Time        `json:"createdAt"`
	UpdatedAt         time.Time        `json:"updatedAt"`
}

type TicketListResponse struct {
	Tickets []TicketResponse `json:"tickets"`
	Total   int              `json:"total"`
}

func NewTicketResponse(ticket *entity.Ticket) TicketResponse {
	return TicketResponse{
		ID:                ticket.ID,
		AssetID:           ticket.AssetID,
		Category:          ticket.Category,
		Severity:          ticket.Severity,
		Status:            ticket.Status,
		Comment:           ticket.Comment,
		Duration:          ticket.Duration,
		DueDate:           ticket.DueDate,
		Reporting:         ticket.Reporting,
		AssignedTo:        ticket.AssignedTo,
		ResolutionComment: ticket.ResolutionComment,
		CreatedAt:         ticket.CreatedAt,
		UpdatedAt:         ticket.UpdatedAt,
	}
}