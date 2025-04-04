package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// To Enforce allowed origins
func EnforceAllowedOrigins(allowedOrigins []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !isAllowedOrigin(c, allowedOrigins) {
				return c.JSON(http.StatusForbidden, map[string]string{
					"message": "Request blocked: Origin not allowed",
				})
			}

			return next(c)
		}
	}
}
