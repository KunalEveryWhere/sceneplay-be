package main

import (
	"fmt"
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

	allowedOrigins := config.AllowedOrigins

	// Remove wildcard allowance even in Developer Mode
	if config.GetEnv("DEVELOPER_MODE", "disabled") == "enabled" {
		allowedOrigins = []string{"*"} // Replace with specific allowed URLs
		fmt.Println("Developer Mode Enabled")
	}

	// Middleware Setup
	e.Use(
		middleware.Logger(),  // Logs requests
		middleware.Recover(), // Recovers from panics
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     allowedOrigins,
			AllowMethods:     []string{http.MethodGet, http.MethodPost},
			AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
			AllowCredentials: true,
		}),
		utils.EnforceAllowedOrigins(config.AllowedOrigins), // Use custom middleware to enforce allowed origins
	)

	// Setup Routes
	routes.SetupRoutes(e)

	// Start the server
	log.Printf("Server running on port %s...", port)
	e.Logger.Fatal(e.Start(":" + port))
}
