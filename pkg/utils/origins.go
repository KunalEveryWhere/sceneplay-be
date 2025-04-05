package utils

import (
	"net/url"
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
	// Allow all origins if developer mode is enabled
	if config.GetEnv("DEVELOPER_MODE", "disabled") == "enabled" {
		return true
	}

	// Get Referer and Origin headers
	referer, origin := c.Request().Header.Get("Referer"), c.Request().Header.Get("Origin")

	// Block requests missing both headers
	if referer == "" && origin == "" {
		return false
	}

	// Normalize allowed origins
	allowedSet := make(map[string]struct{}, len(allowedOrigins))
	for _, origin := range allowedOrigins {
		// Allow all origins if "*" is present
		if origin == "*" {
			return true
		}
		allowedSet[strings.TrimSuffix(origin, "/")] = struct{}{}
	}

	// Check if Referer or Origin matches allowed origins
	for _, header := range []string{referer, origin} {
		if header == "" {
			continue
		}
		if parsedURL, err := url.Parse(header); err == nil {
			if _, exists := allowedSet[strings.TrimSuffix(parsedURL.Host, "/")]; exists {
				return true
			}
		}
	}

	return false
}
