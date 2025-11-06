package auth

import (
	"context"

	authdto "inventory-ticketing-system/application/dto/auth"
	"inventory-ticketing-system/domain/service"
)

type LoginUseCase struct {
	authService service.AuthService
}

func NewLoginUseCase(authService service.AuthService) *LoginUseCase {
	return &LoginUseCase{
		authService: authService,
	}
}

func (uc *LoginUseCase) Execute(ctx context.Context, req *authdto.LoginRequest) (*authdto.LoginResponse, error) {
	token, user, err := uc.authService.Login(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &authdto.LoginResponse{
		Token: token,
		User:  user,
	}, nil
}