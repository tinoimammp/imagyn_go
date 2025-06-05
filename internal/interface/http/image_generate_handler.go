package http

import (
	"log"

	"github.com/tinoimammp/imagyn_go/internal/application/image/dto"
	"github.com/tinoimammp/imagyn_go/internal/application/image/usecases"

	"github.com/gofiber/fiber/v2"
)

// ImageHandler handles HTTP requests related to image generation.
type ImageHandler struct {
	generateImageUseCase *usecases.GenerateImageUseCase
}

// NewImageHandler creates a new ImageHandler.
func NewImageHandler(guc *usecases.GenerateImageUseCase) *ImageHandler {
	return &ImageHandler{
		generateImageUseCase: guc,
	}
}

// GenerateImage handles the POST /api/v1/image/generate endpoint.
func (h *ImageHandler) GenerateImage(c *fiber.Ctx) error {
	var reqDTO dto.GenerateImageRequestDTO
	if err := c.BodyParser(&reqDTO); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Basic validation
	if reqDTO.Prompt == "" || reqDTO.Model == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Prompt and Model are required",
		})
	}
	if reqDTO.Width <= 0 || reqDTO.Height <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Width and Height must be positive integers",
		})
	}

	respDTO, err := h.generateImageUseCase.Execute(c.Context(), &reqDTO)
	if err != nil {
		log.Printf("Error generating image: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate image",
			"details": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(respDTO)
}
