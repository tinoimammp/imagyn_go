package http

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/imagyn_go/internal/application/example/dto"
	"gitlab.com/imagyn_go/internal/application/example/usecases"
)

type ExampleHandler struct {
	CreateExampleCase *usecases.CreateUseCase
}

func NewExampleHandler(app fiber.Router, createUC *usecases.CreateUseCase) {
	handler := &ExampleHandler{
		CreateExampleCase: createUC,
	}

	app.Post("/create", handler.Create)
}

func (h *ExampleHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateExampleRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":    400,
			"message": err.Error(),
		})
	}

	if err := h.CreateExampleCase.Excecute(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":    400,
			"message": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"code":    201,
		"message": "Create example success",
	})
}
