package response

type ProductSKUResponse struct {
	ID           int64   `json:"id"`
	SKU          string  `json:"sku"`
	SKUSignature string  `json:"sku_signature"`
	Price        float64 `json:"price"`
	Status       string  `json:"status"`
	ProductID    int64   `json:"product_id"`
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
