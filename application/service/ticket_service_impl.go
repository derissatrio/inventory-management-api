package service

import (
	"context"
	"errors"
	"time"

	"inventory-ticketing-system/domain/entity"
	"inventory-ticketing-system/domain/repository"
	"inventory-ticketing-system/domain/service"

	"github.com/google/uuid"
)

type TicketServiceImpl struct {
	ticketRepo repository.TicketRepository
	assetRepo  repository.AssetRepository
}

func NewTicketService(ticketRepo repository.TicketRepository, assetRepo repository.AssetRepository) service.TicketService {
	return &TicketServiceImpl{
		ticketRepo: ticketRepo,
		assetRepo:  assetRepo,
	}
}

func (s *TicketServiceImpl) CreateTicket(ctx context.Context, ticket *entity.Ticket) error {
	_, err := s.assetRepo.GetByID(ctx, ticket.AssetID)
	if err != nil {
		return errors.New("asset not found")
	}

	ticket.Status = "open"
	ticket.Duration = s.calculateDuration(ticket.Severity)
	ticket.DueDate = time.Now().Add(time.Duration(ticket.Duration) * time.Hour)

	return s.ticketRepo.Create(ctx, ticket)
}

func (s *TicketServiceImpl) GetTicket(ctx context.Context, id uuid.UUID) (*entity.Ticket, error) {
	return s.ticketRepo.GetByID(ctx, id)
}

func (s *TicketServiceImpl) UpdateTicket(ctx context.Context, id uuid.UUID, ticket *entity.Ticket) error {
	existingTicket, err := s.ticketRepo.GetByID(ctx, id)
	if err != nil {
		return errors.New("ticket not found")
	}

	ticket.ID = id
	ticket.CreatedAt = existingTicket.CreatedAt
	ticket.UpdatedAt = time.Now()

	return s.ticketRepo.Update(ctx, ticket)
}

func (s *TicketServiceImpl) DeleteTicket(ctx context.Context, id uuid.UUID) error {
	_, err := s.ticketRepo.GetByID(ctx, id)
	if err != nil {
		return errors.New("ticket not found")
	}

	return s.ticketRepo.Delete(ctx, id)
}

func (s *TicketServiceImpl) ListTickets(ctx context.Context, limit, offset int, filters map[string]interface{}) ([]*entity.Ticket, int, error) {
	return s.ticketRepo.List(ctx, limit, offset, filters)
}

func (s *TicketServiceImpl) AssignTicket(ctx context.Context, ticketID, assignedTo uuid.UUID) error {
	ticket, err := s.ticketRepo.GetByID(ctx, ticketID)
	if err != nil {
		return errors.New("ticket not found")
	}

	ticket.AssignedTo = &assignedTo
	ticket.Status = "in_progress"

	return s.ticketRepo.Update(ctx, ticket)
}

func (s *TicketServiceImpl) ResolveTicket(ctx context.Context, ticketID uuid.UUID, resolutionComment string) error {
	ticket, err := s.ticketRepo.GetByID(ctx, ticketID)
	if err != nil {
		return errors.New("ticket not found")
	}

	ticket.Status = "resolved"
	ticket.ResolutionComment = resolutionComment

	return s.ticketRepo.Update(ctx, ticket)
}

func (s *TicketServiceImpl) CloseTicket(ctx context.Context, ticketID uuid.UUID) error {
	ticket, err := s.ticketRepo.GetByID(ctx, ticketID)
	if err != nil {
		return errors.New("ticket not found")
	}

	ticket.Status = "closed"

	return s.ticketRepo.Update(ctx, ticket)
}

func (s *TicketServiceImpl) GetTicketsByAsset(ctx context.Context, assetID uuid.UUID) ([]*entity.Ticket, error) {
	return s.ticketRepo.GetByAssetID(ctx, assetID)
}

func (s *TicketServiceImpl) GetTicketsByReporter(ctx context.Context, reporterID uuid.UUID) ([]*entity.Ticket, error) {
	return s.ticketRepo.GetByReporter(ctx, reporterID)
}

func (s *TicketServiceImpl) calculateDuration(severity string) int {
	switch severity {
	case "low":
		return 72
	case "medium":
		return 48
	case "high":
		return 24
	case "critical":
		return 4
	default:
		return 24
	}
}
