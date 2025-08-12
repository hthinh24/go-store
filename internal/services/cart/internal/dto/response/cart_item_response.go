package response

type CartItemResponse struct {
	ID           int64   `json:"id"`
	CartID       int64   `json:"cart_id"`
	ProductID    int64   `json:"product_id"`
	ProductSKUID int64   `json:"product_sku_id"`
	Quantity     int     `json:"quantity"`
	UnitPrice    float64 `json:"unit_price"`
	TotalPrice   float64 `json:"total_price"`
	Status       string  `json:"status"`
}
