package routes

import (
	"context"
	"net/http"
	"time"

	authpb "github.com/ali-naseri-php/Askly/proto/auth"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

func RegisterAuthRoutes(e *echo.Echo, conn *grpc.ClientConn) {
	client := authpb.NewAuthServiceClient(conn)

	e.POST("/auth/login", func(c echo.Context) error {
		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
		}

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		resp, err := client.Login(ctx, &authpb.LoginRequest{
			Email:    req.Email,
			Password: req.Password,
		})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, map[string]string{"token": resp.Token})
	})
}
