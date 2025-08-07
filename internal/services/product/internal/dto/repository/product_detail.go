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

type ProductSKUDetail struct {
	ID            int64    `json:"id"`
	SKU           string   `json:"sku"`
	SKUSignature  string   `json:"sku_signature"`
	ExtraPrice    float64  `json:"extra_price"`
	SaleType      *string  `json:"sale_type,omitempty"`       // "Percentage" or "Fixed"
	SaleValue     *float64 `json:"sale_value,omitempty"`      // Discounted price
	SaleStartDate *string  `json:"sale_start_date,omitempty"` // Sale start date in ISO 8601 format
	SaleEndDate   *string  `json:"sale_end_date,omitempty"`   // Sale end date in ISO 8601 format
	Stock         int32    `json:"stock"`
	Status        string   `json:"status"`
	ProductID     int64    `json:"product_id"`
}
