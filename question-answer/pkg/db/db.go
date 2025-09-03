package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB ایجاد connection به دیتابیس Postgres
func InitDB() *gorm.DB {
	// خواندن DSN از ENV
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=qa_db port=5432 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	fmt.Println("✅ Database connected")
	return db
}
