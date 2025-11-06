package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}

type ErrorResponse struct {
	Success bool `json:"success"`
	Error   struct {
		Code    string                 `json:"code"`
		Message string                 `json:"message"`
		Details []ValidationDetail     `json:"details,omitempty"`
	} `json:"error"`
}

type ValidationDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type PaginationInfo struct {
	Total   int  `json:"total"`
	Limit   int  `json:"limit"`
	Offset  int  `json:"offset"`
	HasMore bool `json:"hasMore"`
}

func SendSuccess(c *gin.Context, statusCode int, message string, data interface{}) {
	response := SuccessResponse{
		Success: true,
		Data:    data,
		Message: message,
	}
	c.JSON(statusCode, response)
}

func SendError(c *gin.Context, statusCode int, code string, message string, details []ValidationDetail) {
	response := ErrorResponse{}
	response.Success = false
	response.Error.Code = code
	response.Error.Message = message
	response.Error.Details = details
	c.JSON(statusCode, response)
}

func SendValidationError(c *gin.Context, err error) {
	details := []ValidationDetail{
		{
			Field:   "validation",
			Message: err.Error(),
		},
	}
	SendError(c, http.StatusBadRequest, "VALIDATION_ERROR", "The provided data is invalid.", details)
}