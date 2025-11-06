package ticket

import (
	"context"

	"github.com/google/uuid"
	ticketdto "inventory-ticketing-system/application/dto/ticket"
	"inventory-ticketing-system/domain/entity"
	"inventory-ticketing-system/domain/service"
)

type CreateTicketUseCase struct {
	ticketService service.TicketService
}

func NewCreateTicketUseCase(ticketService service.TicketService) *CreateTicketUseCase {
	return &CreateTicketUseCase{
		ticketService: ticketService,
	}
}

func (uc *CreateTicketUseCase) Execute(ctx context.Context, req *ticketdto.CreateTicketRequest, reporterID uuid.UUID) (*entity.Ticket, error) {
	ticket := &entity.Ticket{
		ID:        uuid.New(),
		AssetID:   req.AssetID,
		Category:  req.Category,
		Severity:  req.Severity,
		Comment:   req.Comment,
		Reporting: reporterID,
	}

	err := uc.ticketService.CreateTicket(ctx, ticket)
	if err != nil {
		return nil, err
	}

	return ticket, nil
}