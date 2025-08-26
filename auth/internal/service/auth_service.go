package service

import (
	"auth/internal/domain"
	"auth/internal/repository"
	"auth/pkg/utils"
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type AuthService struct {
	userRepo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) *AuthService {
	return &AuthService{userRepo: repo}
}

func (s *AuthService) Register(ctx context.Context, email, password string) (string, error) {
	hashed, err := utils.HashPassword(password)
	if err != nil {
		return "", err
	}

	user := &domain.Auth{
		ID:        uuid.NewString(),
		Email:     email,
		Password:  hashed,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return "", err
	}

	return utils.GenerateJWT(user.ID)
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return "", errors.New("user not found")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	return utils.GenerateJWT(user.ID)
}
