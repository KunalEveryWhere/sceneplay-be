package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found")
	}
}

func GetEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}

func getAllowedOrigins() []string {
	origins := GetEnv("ALLOWED_ORIGINS", "all")
	return strings.Split(origins, ",") // Convert CSV string to slice
}
