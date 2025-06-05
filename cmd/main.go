package main

import (
	"fmt"
	"log"
	"time"

	"imagyn_go/internal/application/image/usecases"
	"imagyn_go/internal/infrastructure/config"
	"imagyn_go/internal/infrastructure/thirdparty"
	"imagyn_go/internal/interfaces/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
	cfg := config.LoadConfig()

	app := fiber.New()

	// Initialize third-party client
	modelClient := thirdparty.NewModelClient(cfg.BaseURLModel, cfg.BearerTokenModel)

	// Initialize use cases
	generateImageUseCase := usecases.NewGenerateImageUseCase(modelClient)

	// Initialize handlers
	imageHandler := http.NewImageHandler(generateImageUseCase)

	// API V1 Group
	apiV1 := app.Group("/api/v1")

	// Apply rate limiting to the image generation route
	apiV1.Post("/image/generate", limiter.New(limiter.Config{
		Max:        cfg.RateLimitPerMinute,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP() // Limit by IP address
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Rate limit exceeded. Please try again in 1 minute.",
			})
		},
	}), imageHandler.GenerateImage)

	// Start server
	log.Printf("Server starting on port %s", cfg.Port)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", cfg.Port)))
}
