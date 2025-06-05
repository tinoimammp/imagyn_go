package image

import "context"

// ImageService
type ImageService interface {
	GenerateImage(ctx context.Context, req *ImageGenerateRequest) (*ImageGenerateResponse, error)
}
