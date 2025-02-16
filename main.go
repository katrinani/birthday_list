package main

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"log/slog"
	"net/http"

	_ "baseToDo/docs"
	"baseToDo/handlers"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/swagger/*", echoSwagger.WrapHandler) // swagger
	e.POST("/gifts", handlers.CreateGift)
	e.PUT("/gifts/:id/reserve", handlers.ReserveGift)

	// Start server
	if err := e.Start(":1323"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
	}
}
