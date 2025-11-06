package repository

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"inventory-ticketing-system/domain/entity"
	"inventory-ticketing-system/domain/repository"
)

type TicketRepositoryImpl struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) repository.TicketRepository {
	return &TicketRepositoryImpl{
		db: db,
	}
}

func (r *TicketRepositoryImpl) Create(ctx context.Context, ticket *entity.Ticket) error {
	return r.db.WithContext(ctx).Create(ticket).Error
}

func (r *TicketRepositoryImpl) GetByID(ctx context.Context, id uuid.UUID) (*entity.Ticket, error) {
	var ticket entity.Ticket
	err := r.db.WithContext(ctx).
		Preload("Asset").
		Where("id = ?", id).
		First(&ticket).Error
	if err != nil {
		return nil, err
	}
	return &ticket, nil
}

func (r *TicketRepositoryImpl) Update(ctx context.Context, ticket *entity.Ticket) error {
	return r.db.WithContext(ctx).Save(ticket).Error
}

func (r *TicketRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&entity.Ticket{}, "id = ?", id).Error
}

func (r *TicketRepositoryImpl) List(ctx context.Context, limit, offset int, filters map[string]interface{}) ([]*entity.Ticket, int, error) {
	var tickets []*entity.Ticket
	var total int64

	query := r.db.WithContext(ctx).Model(&entity.Ticket{})

	for key, value := range filters {
		switch key {
		case "status":
			query = query.Where("status = ?", value)
		case "asset_id":
			query = query.Where("asset_id = ?", value)
		case "severity":
			query = query.Where("severity = ?", value)
		case "category":
			query = query.Where("category ILIKE ?", "%"+value.(string)+"%")
		}
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&tickets).Error
	if err != nil {
		return nil, 0, err
	}

	return tickets, int(total), nil
}

func (r *TicketRepositoryImpl) GetByAssetID(ctx context.Context, assetID uuid.UUID) ([]*entity.Ticket, error) {
	var tickets []*entity.Ticket
	err := r.db.WithContext(ctx).
		Preload("Asset").
		Where("asset_id = ?", assetID).
		Order("created_at DESC").
		Find(&tickets).Error
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (r *TicketRepositoryImpl) GetByReporter(ctx context.Context, reporterID uuid.UUID) ([]*entity.Ticket, error) {
	var tickets []*entity.Ticket
	err := r.db.WithContext(ctx).
		Preload("Asset").
		Where("reporting = ?", reporterID).
		Order("created_at DESC").
		Find(&tickets).Error
	if err != nil {
		return nil, err
	}
	return tickets, nil
}