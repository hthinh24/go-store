package response

type ProductResponse struct {
	ID                int64                `json:"id"`
	Name              string               `json:"name"`
	Description       string               `json:"description"`
	ShortDescription  string               `json:"short_description"`
	ImageURL          string               `json:"image_url"`
	Slug              string               `json:"slug"`
	BasePrice         float64              `json:"base_price"`
	SalePrice         *float64             `json:"sale_price,omitempty"`
	IsFeatured        bool                 `json:"is_featured"`
	SaleStartDate     *string              `json:"sale_start_date,omitempty"`
	SaleEndDate       *string              `json:"sale_end_date,omitempty"`
	Status            string               `json:"status"`
	BrandID           int64                `json:"brand_id"`
	CategoryID        int64                `json:"category_id"`
	UserID            int64                `json:"user_id"`
	Version           int32                `json:"version"`
	ProductAttributes map[string][]string  `json:"product_attributes"`
	ProductSKUs       []ProductSKUResponse `json:"product_skus"`
}

type ProductSKUResponse struct {
	s
	SKU          string  `json:"sku"`
	SKUSignature string  `json:"sku_signature"`
	Price        float64 `json:"price"`
	Status       string  `json:"status"`
	ProductID    int64   `json:"product_id"`
}
