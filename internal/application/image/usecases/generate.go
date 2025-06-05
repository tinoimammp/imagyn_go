package usecases

import (
	"context"
	"fmt"

	"github.com/tinoimammp/imagyn_go/internal/application/image/dto"
	"github.com/tinoimammp/imagyn_go/internal/domain/image"
	"github.com/tinoimammp/imagyn_go/internal/infrastructure/thirdparty"
)

// GenerateImageUseCase handles the business logic for generating images.
type GenerateImageUseCase struct {
	modelClient *thirdparty.ModelClient
}

// NewGenerateImageUseCase creates a new GenerateImageUseCase.
func NewGenerateImageUseCase(mc *thirdparty.ModelClient) *GenerateImageUseCase {
	return &GenerateImageUseCase{
		modelClient: mc,
	}
}

// Execute processes the image generation request by calling the third-party model.
func (uc *GenerateImageUseCase) Execute(ctx context.Context, reqDTO *dto.GenerateImageRequestDTO) (*dto.GenerateImageResponseDTO, error) {
	domainReq := &image.ImageGenerateRequest{
		Model:             reqDTO.Model,
		Prompt:            reqDTO.Prompt,
		Width:             reqDTO.Width,
		Height:            reqDTO.Height,
		NumInferenceSteps: reqDTO.NumInferenceSteps,
		Seed:              reqDTO.Seed,
		GuidanceScale:     reqDTO.GuidanceScale,
		NegativePrompt:    reqDTO.NegativePrompt,
		ResponseExtension: reqDTO.ResponseExtension,
		ResponseFormat:    reqDTO.ResponseFormat,
	}

	domainResp, err := uc.modelClient.Generate(ctx, domainReq)
	if err != nil {
		return nil, fmt.Errorf("failed to generate image from model: %w", err)
	}

	// Correctly extract the b64_json from the first item in the Data array
	respDTO := &dto.GenerateImageResponseDTO{
		ImageBase64: domainResp.Data[0].B64JSON, // FIX APPLIED HERE
	}

	return respDTO, nil
}
