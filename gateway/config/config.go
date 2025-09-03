package config

import "os"

type Config struct {
	AuthServiceURL          string
	QuestionAnswerServiceURL string
	UserServiceURL          string
	ProductServiceURL       string
	GatewayPort             string
}

func Load() *Config {
	cfg := &Config{
		AuthServiceURL:          os.Getenv("AUTH_SERVICE_URL"),
		QuestionAnswerServiceURL: os.Getenv("QUESTION_ANSWER_SERVICE_URL"),
		UserServiceURL:          os.Getenv("USER_SERVICE_URL"),
		ProductServiceURL:       os.Getenv("PRODUCT_SERVICE_URL"),
		GatewayPort:             os.Getenv("GATEWAY_PORT"),
	}

	// مقادیر پیش‌فرض
	if cfg.AuthServiceURL == "" {
		cfg.AuthServiceURL = "localhost:50051"
	}
	if cfg.QuestionAnswerServiceURL == "" {
		cfg.QuestionAnswerServiceURL = "localhost:50052"
	}
	if cfg.UserServiceURL == "" {
		cfg.UserServiceURL = "http://localhost:8001"
	}
	if cfg.ProductServiceURL == "" {
		cfg.ProductServiceURL = "http://localhost:8002"
	}
	if cfg.GatewayPort == "" {
		cfg.GatewayPort = "3000"
	}

	return cfg
}
