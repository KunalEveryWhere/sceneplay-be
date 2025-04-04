package routes

import (
	"sceneplay-be/internal/handlers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/health", handlers.HealthHandler)
	e.POST("/create-payment", handlers.CreatePaymentHandler)
}
