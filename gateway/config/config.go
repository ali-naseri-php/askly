package config

import "os"

type Config struct {
	AuthServiceURL string
	GatewayPort    string
}

func Load() *Config {
	cfg := &Config{
		AuthServiceURL: os.Getenv("AUTH_SERVICE_URL"),
		GatewayPort:    os.Getenv("GATEWAY_PORT"),
	}

	if cfg.AuthServiceURL == "" {
		cfg.AuthServiceURL = "localhost:50051"
	}
	if cfg.GatewayPort == "" {
		cfg.GatewayPort = "8080"
	}
	return cfg
}
