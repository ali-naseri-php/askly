package db

import (
	"time"
	"auth/internal/domain"
)

type AuthDB struct {
	ID        string     `gorm:"primaryKey;size:36"`     // UUID
	Email     string     `gorm:"size:255;uniqueIndex"`   // ایمیل یکتا
	Password  string     `gorm:"size:255"`               // هش پسورد
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`                  // optional
}

// تبدیل UserDB → Domain.User
func (u *AuthDB) ToDomain() *domain.Auth {
	return &domain.Auth{
		ID:        u.ID,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt,
	}
}

// تبدیل Domain.User → UserDB
func FromDomain(user *domain.Auth) *AuthDB {
	return &AuthDB{
		ID:        user.ID,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}
