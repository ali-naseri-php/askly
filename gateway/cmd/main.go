package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	authpb "github.com/ali-naseri-php/Askly/proto/auth"
	"github.com/joho/godotenv" 
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

func main() {
	// load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  .env file not found, falling back to system env")
	}

	authServiceAddr := os.Getenv("AUTH_SERVICE_URL")
	if authServiceAddr == "" {
		authServiceAddr = "localhost:50051"
	}

	gatewayPort := os.Getenv("GATEWAY_PORT")
	if gatewayPort == "" {
		gatewayPort = "8080"
	}

	// connect to gRPC Auth service
	conn, err := grpc.Dial(authServiceAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to auth service: %v", err)
	}
	defer conn.Close()

	authClient := authpb.NewAuthServiceClient(conn)

	e := echo.New()

	e.POST("/login", func(c echo.Context) error {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		resp, err := authClient.Login(ctx, &authpb.LoginRequest{
			Email:    "test@test.com",
			Password: "123",
		})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, map[string]string{"token": resp.Token})
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "Gateway is up!")
	})

	log.Printf("🚀 Gateway running on :%s", gatewayPort)
	e.Logger.Fatal(e.Start(":" + gatewayPort))
}
