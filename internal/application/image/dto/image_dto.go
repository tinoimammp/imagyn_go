package dto

// GenerateImageRequestDTO : HTTP Request body image generation.
type GenerateImageRequestDTO struct {
	Model             string  `json:"model"`
	Prompt            string  `json:"prompt"`
	Width             int     `json:"width"`
	Height            int     `json:"height"`
	NumInferenceSteps int     `json:"num_inference_steps"`
	Seed              int     `json:"seed"`
	GuidanceScale     float64 `json:"guidance_scale"`
	NegativePrompt    string  `json:"negative_prompt"`
	ResponseExtension string  `json:"response_extension"`
	ResponseFormat    string  `json:"response_format"`
}

// GenerateImageResponseDTO : HTTP Response body image generation.
type GenerateImageResponseDTO struct {
	ImageBase64 string `json:"image_b64"`
}
