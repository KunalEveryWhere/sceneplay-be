package handlers

import (
	"net/http"

	"sceneplay-be/internal/models"
	"sceneplay-be/internal/services"

	"github.com/labstack/echo/v4"
)

func CreatePaymentHandler(c echo.Context) error {
	var req models.PaymentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	// Generate Razorpay payment link
	response, err := services.CreatePaymentLink(req)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create payment link"})
	}

	// Return payment link
	return c.JSON(http.StatusOK, response)
}
