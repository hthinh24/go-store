package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type ProductClient interface {
	GetProductSKUByID(ctx context.Context, productSKUID int64) (*ProductSKUDetailResponse, error)
}

type productClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewProductClient(baseURL string) ProductClient {
	return &productClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

type ProductSKUDetailResponse struct {
	ID            int64    `json:"id"`
	SKU           string   `json:"sku"`
	SKUSignature  string   `json:"sku_signature"`
	Price         float64  `json:"price"`
	SaleType      *string  `json:"sale_type"` // "Percentage" or "Fixed"
	SalePrice     *float64 `json:"sale_price"`
	SaleStartDate *string  `json:"sale_start_date"`
	SaleEndDate   *string  `json:"sale_end_date"`
	Stock         int32    `json:"stock"`
	Status        string   `json:"status"`
	ProductID     int64    `json:"product_id"`
}

func (c *productClient) GetProductSKUByID(ctx context.Context, productSKUID int64) (*ProductSKUDetailResponse, error) {
	url := fmt.Sprintf("%s/api/v1/products/skus/%s", c.baseURL, strconv.FormatInt(productSKUID, 10))

	httpReq, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to call product service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("product SKU not found")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("product service returned status: %d", resp.StatusCode)
	}

	var productSKU ProductSKUDetailResponse
	if err := json.NewDecoder(resp.Body).Decode(&productSKU); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	fmt.Printf("Product SKU Details: %+v\n", productSKU)

	return &productSKU, nil
}
