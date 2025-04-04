package models

import "time"

// PaymentRequest represents the incoming donation request
type PaymentRequest struct {
	Amount   int    `json:"amount"`   // Amount in cents / paisa (1 USD = 100 cents)
	Currency string `json:"currency"` // Either USD, INR, others
}

// PaymentResponse represents the Razorpay Payment Link response
type PaymentResponse struct {
	Amount      int       `json:"amount"`
	CreatedAt   time.Time `json:"created-at"`
	Currency    string    `json:"currency"`
	PaymentLink string    `json:"payment-link"`
	ExpireBy    time.Time `json:"expire-by"`
}
