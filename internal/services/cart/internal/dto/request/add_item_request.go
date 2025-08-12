package request

type AddItemRequest struct {
	ProductID    int64 `json:"product_id"`
	ProductSKUID int64 `json:"product_sku_id"`
}
