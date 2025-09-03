package routes

import (
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

func RegisterRoutes(e *echo.Echo, conn *grpc.ClientConn) {
	// health check
	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "Gateway is up!")
	})

	RegisterAuthRoutes(e, conn)
	RegisterQuestionAnswerRoutes(e,conn)
}
