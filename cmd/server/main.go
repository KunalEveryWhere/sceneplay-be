package main

import (
	"log"
	"net/http"

	"sceneplay-be/internal/config"
	"sceneplay-be/internal/routes"
	"sceneplay-be/pkg/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Set the server port from environment
	port := "8080"

	// Create a new Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())  // Logs requests
	e.Use(middleware.Recover()) // Recovers from panics
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: config.AllowedOrigins,
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))

	// Use custom middleware to enforce allowed origins
	e.Use(utils.EnforceAllowedOrigins(config.AllowedOrigins))

	// Setup Routes
	routes.SetupRoutes(e)

	// Start the server
	log.Printf("Server running on port %s...", port)
	e.Logger.Fatal(e.Start(":" + port))
}
