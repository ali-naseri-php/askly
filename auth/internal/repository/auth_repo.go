package repository

import (
	"context"
	"auth/internal/domain"
	"auth/internal/db"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.Auth) error
	FindByEmail(ctx context.Context, email string) (*domain.Auth, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, user *domain.Auth) error {
	// تبدیل Domain → DB
	userDB := db.FromDomain(user)
	return r.db.WithContext(ctx).Create(userDB).Error
}

func (r *userRepo) FindByEmail(ctx context.Context, email string) (*domain.Auth, error) {
	var userDB db.AuthDB
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&userDB).Error
	if err != nil {
		return nil, err
	}
	// تبدیل DB → Domain
	return userDB.ToDomain(), nil
}
