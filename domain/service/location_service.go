package service

import (
	"context"

	"github.com/google/uuid"
	"inventory-ticketing-system/domain/entity"
)

type LocationService interface {
	CreateLocation(ctx context.Context, location *entity.Location) error
	GetLocation(ctx context.Context, id uuid.UUID) (*entity.Location, error)
	UpdateLocation(ctx context.Context, id uuid.UUID, location *entity.Location) error
	DeleteLocation(ctx context.Context, id uuid.UUID) error
	ListLocations(ctx context.Context, limit, offset int) ([]*entity.Location, int, error)
	GetLocationByName(ctx context.Context, name string) (*entity.Location, error)
}