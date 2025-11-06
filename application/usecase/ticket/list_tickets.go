package ticket

import (
	"context"

	ticketdto "inventory-ticketing-system/application/dto/ticket"
	"inventory-ticketing-system/domain/service"
)

type ListTicketsUseCase struct {
	ticketService service.TicketService
}

func NewListTicketsUseCase(ticketService service.TicketService) *ListTicketsUseCase {
	return &ListTicketsUseCase{
		ticketService: ticketService,
	}
}

func (uc *ListTicketsUseCase) Execute(ctx context.Context, limit, offset int, filters map[string]interface{}) (*ticketdto.TicketListResponse, error) {
	tickets, total, err := uc.ticketService.ListTickets(ctx, limit, offset, filters)
	if err != nil {
		return nil, err
	}

	ticketResponses := make([]ticketdto.TicketResponse, len(tickets))
	for i, ticket := range tickets {
		ticketResponses[i] = ticketdto.NewTicketResponse(ticket)
	}

	return &ticketdto.TicketListResponse{
		Tickets: ticketResponses,
		Total:   total,
	}, nil
}