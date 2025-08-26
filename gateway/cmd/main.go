package main

import (
	"log"
	"github.com/ali-naseri-php/Askly/gateway/config"
	"github.com/ali-naseri-php/Askly/gateway/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// load .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  .env file not found, falling back to system env")
	}

	// config
	cfg := config.Load()

	// gRPC connection to Auth service
	conn, err := grpc.Dial(
		cfg.AuthServiceURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect to auth service: %v", err)
	}
	defer conn.Close()

	// echo
	e := echo.New()

	// register routes
	routes.RegisterRoutes(e, conn)

	log.Printf("🚀 Gateway running on :%s", cfg.GatewayPort)
	e.Logger.Fatal(e.Start(":" + cfg.GatewayPort))
}
