package repository

import (
	"context"

	"github.com/google/uuid"
	"inventory-ticketing-system/domain/entity"
)

type LocationRepository interface {
	Create(ctx context.Context, location *entity.Location) error
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Location, error)
	Update(ctx context.Context, location *entity.Location) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, limit, offset int) ([]*entity.Location, int, error)
	GetByName(ctx context.Context, name string) (*entity.Location, error)
}