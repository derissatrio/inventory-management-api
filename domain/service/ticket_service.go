package service

import (
	"context"

	"github.com/google/uuid"
	"inventory-ticketing-system/domain/entity"
)

type TicketService interface {
	CreateTicket(ctx context.Context, ticket *entity.Ticket) error
	GetTicket(ctx context.Context, id uuid.UUID) (*entity.Ticket, error)
	UpdateTicket(ctx context.Context, id uuid.UUID, ticket *entity.Ticket) error
	DeleteTicket(ctx context.Context, id uuid.UUID) error
	ListTickets(ctx context.Context, limit, offset int, filters map[string]interface{}) ([]*entity.Ticket, int, error)
	AssignTicket(ctx context.Context, ticketID, assignedTo uuid.UUID) error
	ResolveTicket(ctx context.Context, ticketID uuid.UUID, resolutionComment string) error
	CloseTicket(ctx context.Context, ticketID uuid.UUID) error
	GetTicketsByAsset(ctx context.Context, assetID uuid.UUID) ([]*entity.Ticket, error)
	GetTicketsByReporter(ctx context.Context, reporterID uuid.UUID) ([]*entity.Ticket, error)
}