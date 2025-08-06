package repository

type ProductWithAttributeValues struct {
	ID              int64  `json:"id"`
	AttributeName   string `json:"name"`
	AttributeValues string `json:"attribute_values"`
}

type ProductWithOptionValues struct {
	ID           int64  `json:"id"`
	OptionNames  string `json:"option_name"`
	OptionValues string `json:"option_values"`
}

type ProductSKUWithInventory struct {
	ID           int64   `json:"id"`
	SKU          string  `json:"sku"`
	SKUSignature string  `json:"sku_signature"`
	Price        float64 `json:"price"`
	Stock        int32   `json:"stock"`
	Status       string  `json:"status"`
	ProductID    int64   `json:"product_id"`
}
