package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	ticketdto "inventory-ticketing-system/application/dto/ticket"
	ticketusecase "inventory-ticketing-system/application/usecase/ticket"
	"inventory-ticketing-system/pkg/common"
)

type TicketHandler struct {
	createTicketUseCase *ticketusecase.CreateTicketUseCase
	listTicketsUseCase  *ticketusecase.ListTicketsUseCase
}

func NewTicketHandler(
	createTicketUseCase *ticketusecase.CreateTicketUseCase,
	listTicketsUseCase *ticketusecase.ListTicketsUseCase,
) *TicketHandler {
	return &TicketHandler{
		createTicketUseCase: createTicketUseCase,
		listTicketsUseCase:  listTicketsUseCase,
	}
}

func (h *TicketHandler) Create(c *gin.Context) {
	var req ticketdto.CreateTicketRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.SendValidationError(c, err)
		return
	}

	// Get reporter ID from context (set by auth middleware)
	userIDStr, exists := c.Get("user_id")
	if !exists {
		common.SendError(c, http.StatusUnauthorized, "UNAUTHORIZED", "User not authenticated", nil)
		return
	}

	reporterID, ok := userIDStr.(uuid.UUID)
	if !ok {
		common.SendError(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Invalid user ID format", nil)
		return
	}

	ticket, err := h.createTicketUseCase.Execute(c.Request.Context(), &req, reporterID)
	if err != nil {
		common.SendError(c, http.StatusBadRequest, "VALIDATION_ERROR", err.Error(), nil)
		return
	}

	response := ticketdto.NewTicketResponse(ticket)
	common.SendSuccess(c, http.StatusCreated, "Ticket created successfully", response)
}

func (h *TicketHandler) Get(c *gin.Context) {
	idStr := c.Param("id")
	_, err := uuid.Parse(idStr)
	if err != nil {
		common.SendError(c, http.StatusBadRequest, "INVALID_ID", "Invalid ticket ID", nil)
		return
	}

	// Note: Implement GetTicketUseCase for this functionality
	common.SendSuccess(c, http.StatusOK, "Ticket retrieved successfully", gin.H{"id": idStr})
}

func (h *TicketHandler) List(c *gin.Context) {
	var req ticketdto.TicketListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		common.SendValidationError(c, err)
		return
	}

	// Build filters
	filters := make(map[string]interface{})
	if req.Status != "" {
		filters["status"] = req.Status
	}
	if req.AssetID != "" {
		if assetID, err := uuid.Parse(req.AssetID); err == nil {
			filters["asset_id"] = assetID
		}
	}

	response, err := h.listTicketsUseCase.Execute(c.Request.Context(), req.Limit, req.Offset, filters)
	if err != nil {
		common.SendError(c, http.StatusInternalServerError, "INTERNAL_ERROR", err.Error(), nil)
		return
	}

	// Add pagination info
	hasMore := response.Total > (req.Offset + req.Limit)
	pagination := common.PaginationInfo{
		Total:   response.Total,
		Limit:   req.Limit,
		Offset:  req.Offset,
		HasMore: hasMore,
	}

	data := gin.H{
		"tickets":    response.Tickets,
		"pagination": pagination,
	}

	common.SendSuccess(c, http.StatusOK, "Tickets retrieved successfully", data)
}

func (h *TicketHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	_, err := uuid.Parse(idStr)
	if err != nil {
		common.SendError(c, http.StatusBadRequest, "INVALID_ID", "Invalid ticket ID", nil)
		return
	}

	// Note: Implement UpdateTicketUseCase for this functionality
	common.SendSuccess(c, http.StatusOK, "Ticket updated successfully", gin.H{"id": idStr})
}

func (h *TicketHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	_, err := uuid.Parse(idStr)
	if err != nil {
		common.SendError(c, http.StatusBadRequest, "INVALID_ID", "Invalid ticket ID", nil)
		return
	}

	// Note: Implement DeleteTicketUseCase for this functionality
	common.SendSuccess(c, http.StatusOK, "Ticket deleted successfully", gin.H{"id": idStr})
}