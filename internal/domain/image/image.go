package image

// ImageGenerateRequest represents the payload structure for the third-party image generation API.
type ImageGenerateRequest struct {
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

// ImageData represents a single item in the "data" array from the third-party API.
type ImageData struct {
	B64JSON string `json:"b64_json"`
	URL     *string `json:"url"` // Use pointer for nullable field
}

// ImageGenerateResponse represents the expected response structure from the third-party image generation API.
type ImageGenerateResponse struct {
	Data []ImageData `json:"data"`
	ID   string      `json:"id"`
}
