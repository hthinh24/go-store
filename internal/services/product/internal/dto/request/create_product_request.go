package request

import "time"

type CreateProductRequest struct {
	// TODO - Add validation tags
	Name              string                    `json:"name" binding:"required"`               // Product name
	Description       string                    `json:"description"`                           // Product description
	ShortDescription  string                    `json:"short_description"`                     // Short description of the product
	ImageURL          string                    `json:"image_url" binding:"required"`          // Product image URL
	Slug              string                    `json:"slug" binding:"required"`               // Unique slug for the product
	BasePrice         float64                   `json:"base_price" binding:"required"`         // Product price
	SalePrice         *float64                  `json:"sale_price"`                            // Discounted price
	IsFeatured        bool                      `json:"is_featured"`                           // Whether the product is featured
	SaleStartDate     *time.Time                `json:"sale_start_date"`                       // Sale start date in ISO 8601 format
	SaleEndDate       *time.Time                `json:"sale_end_date"`                         // Sale end date in ISO 8601 format
	Status            string                    `json:"status" binding:"required"`             // Product status (e.g., active, inactive)
	BrandID           int64                     `json:"brand_id" binding:"required"`           // Brand ID
	CategoryID        int64                     `json:"category_id" binding:"required"`        // Category ID
	UserID            int64                     `json:"user_id" binding:"required"`            // User ID of the product creator
	ProductAttributes map[int64][]string        `json:"product_attributes" binding:"required"` // Product attributes as key-value pairs
	ProductSKUs       []CreateProductSKURequest `json:"product_skus" binding:"required"`       // List of product SKUs
	OptionValues      map[int64][]string        `json:"option_values"`                         // Product option values as key-value pairs
}

type CreateProductSKURequest struct {
	SKU   string  `json:"sku" binding:"required"`   // Stock Keeping Unit
	Price float64 `json:"price" binding:"required"` // Price of the SKU
	Stock int32   `json:"stock" binding:"required"` // Stock quantity
}

type CreateProductWithoutSKURequest struct {
	// TODO - Add validation tags
	Name              string             `json:"name" binding:"required"`               // Product name
	Description       string             `json:"description"`                           // Product description
	ShortDescription  string             `json:"short_description"`                     // Short description of the product
	ImageURL          string             `json:"image_url" binding:"required"`          // Product image URL
	Slug              string             `json:"slug" binding:"required"`               // Unique slug for the product
	BasePrice         float64            `json:"base_price" binding:"required"`         // Product price
	SalePrice         *float64           `json:"sale_price"`                            // Discounted price
	IsFeatured        bool               `json:"is_featured"`                           // Whether the product is featured
	SaleStartDate     *time.Time         `json:"sale_start_date"`                       // Sale start date in ISO 8601 format
	SaleEndDate       *time.Time         `json:"sale_end_date"`                         // Sale end date in ISO 8601 format
	BrandID           int64              `json:"brand_id" binding:"required"`           // Brand ID
	CategoryID        int64              `json:"category_id" binding:"required"`        // Category ID
	UserID            int64              `json:"user_id" binding:"required"`            // User ID of the product creator
	ProductAttributes map[int64][]string `json:"product_attributes" binding:"required"` // Product attributes as key-value pairs
	OptionValues      map[int64][]string `json:"option_values"`                         // Product option values as key-value pairs
}
