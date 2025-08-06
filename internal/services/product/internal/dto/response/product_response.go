package response

import (
	"time"
)

type ProductResponse struct {
	ID               int64      `json:"id"`
	Name             string     `json:"name"`
	ShortDescription string     `json:"short_description"`
	ImageURL         string     `json:"image_url"`
	BasePrice        float64    `json:"base_price"`
	SalePrice        *float64   `json:"sale_price,omitempty"`
	IsFeatured       bool       `json:"is_featured"`
	SaleStartDate    *time.Time `json:"sale_start_date,omitempty"`
	SaleEndDate      *time.Time `json:"sale_end_date,omitempty"`
	Status           string     `json:"status"`
	BrandID          int64      `json:"brand_id"`
	CategoryID       int64      `json:"category_id"`
	UserID           int64      `json:"user_id"`
}

type ProductDetailResponse struct {
	ID               int64                                  `json:"id"`
	Name             string                                 `json:"name"`
	Description      string                                 `json:"description"`
	ShortDescription string                                 `json:"short_description"`
	ImageURL         string                                 `json:"image_url"`
	Slug             string                                 `json:"slug"`
	BasePrice        float64                                `json:"base_price"`
	SalePrice        *float64                               `json:"sale_price,omitempty"`
	IsFeatured       bool                                   `json:"is_featured"`
	SaleStartDate    *time.Time                             `json:"sale_start_date,omitempty"`
	SaleEndDate      *time.Time                             `json:"sale_end_date,omitempty"`
	Status           string                                 `json:"status"`
	BrandID          int64                                  `json:"brand_id"`
	CategoryID       int64                                  `json:"category_id"`
	UserID           int64                                  `json:"user_id"`
	Version          int32                                  `json:"version"`
	AttributeValues  *[]*ProductWithAttributeValuesResponse `json:"attribute_values"`
	ProductSKUs      *[]*ProductSKUWithInventoryResponse    `json:"product_skus"`
	OptionValues     *[]*ProductWithOptionValuesResponse    `json:"option_values"`
}

type ProductWithAttributeValuesResponse struct {
	ID              int64  `json:"id"`
	AttributeName   string `json:"name"`
	AttributeValues string `json:"attribute_values"`
}

type ProductWithOptionValuesResponse struct {
	ID           int64  `json:"id"`
	OptionNames  string `json:"option_name"`
	OptionValues string `json:"option_values"`
}
