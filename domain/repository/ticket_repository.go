package repository

import (
	"context"

	"github.com/google/uuid"
	"inventory-ticketing-system/domain/entity"
)

type TicketRepository interface {
	Create(ctx context.Context, ticket *entity.Ticket) error
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Ticket, error)
	Update(ctx context.Context, ticket *entity.Ticket) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, limit, offset int, filters map[string]interface{}) ([]*entity.Ticket, int, error)
	GetByAssetID(ctx context.Context, assetID uuid.UUID) ([]*entity.Ticket, error)
	GetByReporter(ctx context.Context, reporterID uuid.UUID) ([]*entity.Ticket, error)
}