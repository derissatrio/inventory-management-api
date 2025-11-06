package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"inventory-ticketing-system/domain/entity"
	"inventory-ticketing-system/domain/repository"
	"inventory-ticketing-system/domain/service"
	"inventory-ticketing-system/infrastructure/jwt"
)

type AuthServiceImpl struct {
	userRepo    repository.UserRepository
	jwtManager  *jwt.JWTManager
	tokenExpiry time.Duration
}

func NewAuthService(userRepo repository.UserRepository, jwtManager *jwt.JWTManager) service.AuthService {
	return &AuthServiceImpl{
		userRepo:    userRepo,
		jwtManager:  jwtManager,
		tokenExpiry: 24 * time.Hour, // 24 hours
	}
}

func (s *AuthServiceImpl) Login(ctx context.Context, email, password string) (string, *entity.User, error) {
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	err = s.CheckPassword(user.PasswordHash, password)
	if err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	token, err := s.jwtManager.GenerateToken(user.ID, user.Role, s.tokenExpiry)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

func (s *AuthServiceImpl) Register(ctx context.Context, user *entity.User) error {
	existingUser, err := s.userRepo.GetByEmail(ctx, user.Email)
	if err == nil && existingUser != nil {
		return errors.New("user with this email already exists")
	}

	hashedPassword, err := s.HashPassword(user.PasswordHash)
	if err != nil {
		return err
	}

	user.PasswordHash = hashedPassword
	user.Role = "employee" // Default role for registration

	return s.userRepo.Create(ctx, user)
}

func (s *AuthServiceImpl) ValidateToken(ctx context.Context, token string) (uuid.UUID, string, error) {
	return s.jwtManager.ValidateToken(token)
}

func (s *AuthServiceImpl) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (s *AuthServiceImpl) CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}