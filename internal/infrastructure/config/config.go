package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds all application configurations.
type Config struct {
	Port                string
	BaseURLModel        string
	BearerTokenModel    string
	RateLimitPerMinute  int
}

// LoadConfig reads configuration from environment variables.
func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file, continuing without it. Err: %v", err)
	}

	rateLimit, err := strconv.Atoi(os.Getenv("RATE_LIMIT_PER_MINUTE"))
	if err != nil {
		rateLimit = 25 // Default value if not set or invalid
	}

	return &Config{
		Port:                os.Getenv("PORT"),
		BaseURLModel:        os.Getenv("BASE_URL_MODEL_IMAGE"),
		BearerTokenModel:    os.Getenv("BEARER_TOKEN_MODEL"),
		RateLimitPerMinute:  rateLimit,
	}
}
