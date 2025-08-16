package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CartClient interface {
	CreateCart(ctx context.Context, userID int64) error
}

type cartClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewCartClient(baseURL string) CartClient {
	return &cartClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

type CreateCartRequest struct {
	UserID int64 `json:"user_id"`
}

func (c *cartClient) CreateCart(ctx context.Context, userID int64) error {
	req := CreateCartRequest{
		UserID: userID,
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal cart request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", c.baseURL+"/api/v1/cart/register", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("failed to call cart service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("cart service returned status: %d", resp.StatusCode)
	}

	return nil
}
