package utils

import (
	"strings"

	"sceneplay-be/internal/config"

	"github.com/labstack/echo/v4"
)

// IsAllowedOrigin checks if the Referer or Origin header of the given request
// matches any of the allowed origins.
// Returns:
//   - bool: True if the request's Referer or Origin header matches any of the
//     allowed origins, false otherwise.
func isAllowedOrigin(c echo.Context, allowedOrigins []string) bool {
	// Get the DeveloperMode environment variable
	// If the value is 'enabled', then tools like Postman are allowed to communicate with the server
	DeveloperMode := config.GetEnv("DEVELOPER_MODE", "disabled")

	// Get the Referer and Origin headers from the request
	referer := c.Request().Header.Get("Referer")
	origin := c.Request().Header.Get("Origin")

	// Block requests missing both Referer and Origin headers, unless Developer Mode is enabled.
	if referer == "" && origin == "" {
		return DeveloperMode == "enabled"
	}

	// Iterate over each allowed origin
	for _, allowedOrigin := range allowedOrigins {
		// If the allowed origin is "*", return true
		if allowedOrigin == "*" {
			return true
		}

		// Check if the Referer or Origin header contains the allowed origin
		// Make sure referer and origin are not empty before checking
		if (referer != "" && strings.Contains(referer, allowedOrigin)) ||
			(origin != "" && strings.Contains(origin, allowedOrigin)) {
			// If a match is found, return true
			return true
		}
	}

	// If no match is found, return false
	return false
}
