package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	authdto "inventory-ticketing-system/application/dto/auth"
	authusecase "inventory-ticketing-system/application/usecase/auth"
	"inventory-ticketing-system/pkg/common"
)

type AuthHandler struct {
	loginUseCase *authusecase.LoginUseCase
}

func NewAuthHandler(loginUseCase *authusecase.LoginUseCase) *AuthHandler {
	return &AuthHandler{
		loginUseCase: loginUseCase,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req authdto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.SendValidationError(c, err)
		return
	}

	response, err := h.loginUseCase.Execute(c.Request.Context(), &req)
	if err != nil {
		common.SendError(c, http.StatusUnauthorized, "UNAUTHORIZED", err.Error(), nil)
		return
	}

	common.SendSuccess(c, http.StatusOK, "Login successful", response)
}