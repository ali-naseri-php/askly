package config

import (
	"os"
)

// Config struct برای نگهداری تنظیمات پروژه
type Config struct {
	DB_DSN      string
	ServicePort string
	DevMode     bool
}

// LoadConfig خواندن مقادیر ENV و ساخت Config
func LoadConfig() *Config {
	dev := false
	if os.Getenv("DEV") == "true" {
		dev = true
	}

	port := os.Getenv("SERVICE_PORT")
	if port == "" {
		port = "50051"
	}

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=qa_db port=5432 sslmode=disable"
	}

	return &Config{
		DB_DSN:      dsn,
		ServicePort: port,
		DevMode:     dev,
	}
}
	