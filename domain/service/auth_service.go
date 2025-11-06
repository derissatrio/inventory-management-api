package service

import (
	"context"

	"github.com/google/uuid"
	"inventory-ticketing-system/domain/entity"
)

type AuthService interface {
	Login(ctx context.Context, email, password string) (string, *entity.User, error)
	Register(ctx context.Context, user *entity.User) error
	ValidateToken(ctx context.Context, token string) (uuid.UUID, string, error)
	HashPassword(password string) (string, error)
	CheckPassword(hashedPassword, password string) error
}