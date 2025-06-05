package thirdparty

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/tinoimammp/imagyn_go/internal/domain/image"
)

// ModelClient interacts with the third-party image generation model API.
type ModelClient struct {
	baseURL     string
	bearerToken string
	httpClient  *http.Client
}

// NewModelClient creates a new ModelClient.
func NewModelClient(baseURL, bearerToken string) *ModelClient {
	return &ModelClient{
		baseURL:     baseURL,
		bearerToken: bearerToken,
		httpClient: &http.Client{
			Timeout: 30 * time.Second, // Set a timeout for the HTTP request
		},
	}
}

// Generate sends a request to the third-party model to generate an image.
func (c *ModelClient) Generate(ctx context.Context, req *image.ImageGenerateRequest) (*image.ImageGenerateResponse, error) {
	payloadBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request payload: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", c.baseURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+c.bearerToken)

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request to model: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("model API returned non-OK status: %s, body: %s", resp.Status, string(bodyBytes))
	}

	var modelResp image.ImageGenerateResponse
	if err := json.NewDecoder(resp.Body).Decode(&modelResp); err != nil {
		return nil, fmt.Errorf("failed to decode model API response: %w", err)
	}

	return &modelResp, nil
}
