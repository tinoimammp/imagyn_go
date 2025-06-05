package main

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/imagyn_go/internal/application/example/usecases"
	"gitlab.com/imagyn_go/internal/domain/example"
	"gitlab.com/imagyn_go/internal/infrastructure/config"
	"gitlab.com/imagyn_go/internal/infrastructure/persistence/postgres"
	"gitlab.com/imagyn_go/internal/interface/http"
	"log"
)

func main() {
	dbConfig := config.NewDatabaseConfig()
	dbProvider, err := postgres.NewDatabaseProvider(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}
	app := fiber.New()

	// Repositories
	exampleRepo := postgres.NewExampleRepository(dbProvider.GetDB())

	// Services
	exampleService := example.NewExampleService()

	// Use cases
	createUC := usecases.NewExampleUseCase(exampleRepo, exampleService)

	// HTTP handlers
	api := app.Group("/api/v1")
	http.NewExampleHandler(api, &createUC)

	log.Fatal(app.Listen(":3000"))

}
