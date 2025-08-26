package domain

import "time"

// User = موجودیت دامین، مستقل از دیتابیس و gRPC
type Auth struct {
	ID        string     // UUID
	Email     string
	Password  string     // هش شده
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time // optional
}
