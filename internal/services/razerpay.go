package services

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"sceneplay-be/internal/config"
	"sceneplay-be/internal/models"

	"github.com/go-resty/resty/v2"
)

func CreatePaymentLink(req models.PaymentRequest) (models.PaymentResponse, error) {
	client := resty.New()

	// Load RazerPay Creds & Callback URLs
	key := config.GetEnv("RAZORPAY_KEY_ID", "")
	secret := config.GetEnv("RAZORPAY_KEY_SECRET", "")
	callbackURL := config.GetEnv("RAZORPAY_CALLBACK_URL", "")

	if key == "" || secret == "" || callbackURL == "" {
		return models.PaymentResponse{}, fmt.Errorf("missing Razorpay credentials or Callback URL")
	}

	// Convert amount to cents / paisa
	amountInCents := req.Amount * 100

	// Prepare request payload
	payload := map[string]interface{}{
		"amount":                   amountInCents,
		"currency":                 req.Currency,
		"accept_partial":           false,
		"first_min_partial_amount": 0,
		"notify": map[string]bool{
			"sms":   false,
			"email": false,
		},
		"expire_by":       time.Now().Add(72 * time.Hour).Unix(), // 3 days from the date of request
		"reminder_enable": false,
		"description":     "Conjunction - Film Donation Payment",
		"callback_url":    callbackURL,
		"callback_method": "get",
	}

	// Make API request
	resp, err := client.R().
		SetBasicAuth(key, secret).
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post(config.RazorpayBaseURL)
	if err != nil {
		log.Println("Error creating payment link:", err)
		return models.PaymentResponse{}, err
	}

	// Parse response
	var responseData map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &responseData); err != nil {
		return models.PaymentResponse{}, err
	}

	// Extract necessary fields
	paymentLink, ok := responseData["short_url"].(string)
	if !ok {
		return models.PaymentResponse{}, fmt.Errorf("failed to extract payment link")
	}

	createdAtUnix, _ := responseData["created_at"].(float64)
	expireByUnix, _ := responseData["expire_by"].(float64)

	createdAt := time.Unix(int64(createdAtUnix), 0)
	expireBy := time.Unix(int64(expireByUnix), 0)

	// Return formatted response
	return models.PaymentResponse{
		Amount:      req.Amount,
		CreatedAt:   createdAt,
		Currency:    req.Currency,
		PaymentLink: paymentLink,
		ExpireBy:    expireBy,
	}, nil
}
