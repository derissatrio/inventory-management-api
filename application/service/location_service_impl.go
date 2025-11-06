package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"inventory-ticketing-system/domain/entity"
	"inventory-ticketing-system/domain/repository"
	"inventory-ticketing-system/domain/service"
)

type LocationServiceImpl struct {
	locationRepo repository.LocationRepository
}

func NewLocationService(locationRepo repository.LocationRepository) service.LocationService {
	return &LocationServiceImpl{
		locationRepo: locationRepo,
	}
}

func (s *LocationServiceImpl) CreateLocation(ctx context.Context, location *entity.Location) error {
	if location.Capacity <= 0 {
		location.Capacity = 0 // Default capacity
	}

	existingLocation, err := s.locationRepo.GetByName(ctx, location.Name)
	if err == nil && existingLocation != nil {
		return errors.New("location with this name already exists")
	}

	return s.locationRepo.Create(ctx, location)
}

func (s *LocationServiceImpl) GetLocation(ctx context.Context, id uuid.UUID) (*entity.Location, error) {
	return s.locationRepo.GetByID(ctx, id)
}

func (s *LocationServiceImpl) UpdateLocation(ctx context.Context, id uuid.UUID, location *entity.Location) error {
	existingLocation, err := s.locationRepo.GetByID(ctx, id)
	if err != nil {
		return errors.New("location not found")
	}

	location.ID = id
	location.CreatedAt = existingLocation.CreatedAt

	return s.locationRepo.Update(ctx, location)
}

func (s *LocationServiceImpl) DeleteLocation(ctx context.Context, id uuid.UUID) error {
	_, err := s.locationRepo.GetByID(ctx, id)
	if err != nil {
		return errors.New("location not found")
	}

	return s.locationRepo.Delete(ctx, id)
}

func (s *LocationServiceImpl) ListLocations(ctx context.Context, limit, offset int) ([]*entity.Location, int, error) {
	return s.locationRepo.List(ctx, limit, offset)
}

func (s *LocationServiceImpl) GetLocationByName(ctx context.Context, name string) (*entity.Location, error) {
	return s.locationRepo.GetByName(ctx, name)
}