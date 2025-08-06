package response

type ProductSKUResponse struct {
	ID           int64   `json:"id"`
	SKU          string  `json:"sku"`
	SKUSignature string  `json:"sku_signature"`
	Price        float64 `json:"price"`
	Status       string  `json:"status"`
	ProductID    int64   `json:"product_id"`
}

type ProductSKUWithInventoryResponse struct {
	ID           int64   `json:"id"`
	SKU          string  `json:"sku"`
	SKUSignature string  `json:"sku_signature"`
	Price        float64 `json:"price"`
	Stock        int32   `json:"stock"`
	Status       string  `json:"status"`
	ProductID    int64   `json:"product_id"`
}
