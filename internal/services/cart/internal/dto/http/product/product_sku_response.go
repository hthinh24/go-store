package product

type ProductSKUDetailResponse struct {
	ID            int64    `json:"id"`
	SKU           string   `json:"sku"`
	SKUSignature  string   `json:"sku_signature"`
	Price         float64  `json:"price"`
	SaleType      string   `json:"sale_type,omitempty"` // "Percentage" or "Fixed"
	SalePrice     *float64 `json:"sale_price,omitempty"`
	SaleStartDate *string  `json:"sale_start_date,omitempty"`
	SaleEndDate   *string  `json:"sale_end_date,omitempty"`
	Stock         int32    `json:"stock"`
	Status        string   `json:"status"`
	ProductID     int64    `json:"product_id"`
}
