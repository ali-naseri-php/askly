package main

import (
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
    "os"

    "github.com/joho/godotenv"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

// reverse proxy handler
func reverseProxy(target string) echo.HandlerFunc {
    targetURL, _ := url.Parse(target)
    proxy := httputil.NewSingleHostReverseProxy(targetURL)

    return func(c echo.Context) error {
        c.Request().Header.Set("X-Forwarded-Host", c.Request().Host)
        c.Request().Header.Set("X-Origin-Host", targetURL.Host)
        proxy.ServeHTTP(c.Response(), c.Request())
        return nil
    }
}

func main() {
    // load .env file
    if err := godotenv.Load(); err != nil {
        log.Println("⚠️  .env file not found, falling back to system env")
    }

    // read env vars
    authService := os.Getenv("AUTH_SERVICE_URL")
    userService := os.Getenv("USER_SERVICE_URL")
    gatewayPort := os.Getenv("GATEWAY_PORT")
    if gatewayPort == "" {
        gatewayPort = "8080" // default
    }

    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // routes
    e.Any("/auth/*", reverseProxy(authService))
    e.Any("/user/*", reverseProxy(userService))

    e.GET("/ping", func(c echo.Context) error {
        return c.String(http.StatusOK, "Gateway is up!")
    })

    log.Printf("🚀 Gateway running on :%s", gatewayPort)
    e.Logger.Fatal(e.Start(":" + gatewayPort))
}
