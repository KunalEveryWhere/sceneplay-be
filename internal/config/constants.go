package config

// CORS Allowed Origins
// Read allowed origins from ENV or use defaults
var AllowedOrigins = getAllowedOrigins()

// RazerPay API URL
const RazorpayBaseURL = "https://api.razorpay.com/v1/payment_links"
