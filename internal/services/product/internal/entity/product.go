package entity

import (
	"github.com/hthinh24/go-store/internal/pkg/entity"
	"time"
)

type Product struct {
	entity.BaseEntity
	Name             string     `json:"name" gorm:"column:name;type:varchar(255);not null"`
	Description      string     `json:"description" gorm:"column:description;type:text;not null"`
	ShortDescription string     `json:"short_description" gorm:"column:short_description;type:varchar(255);not null"`
	ImageURL         string     `json:"image_url" gorm:"column:image_url;type:varchar(500);not null"`
	Slug             string     `json:"slug" gorm:"column:slug;type:varchar(255);not null;uniqueIndex"`
	BasePrice        float64    `json:"base_price" gorm:"column:base_price;type:decimal(10,2);not null;default:0.00"`
	SalePrice        *float64   `json:"sale_price,omitempty" gorm:"column:sale_price;type:decimal(10,2)"`
	IsFeatured       bool       `json:"is_featured" gorm:"column:is_featured;default:false"`
	SaleStartDate    *time.Time `json:"sale_start_date,omitempty" gorm:"column:sale_start_date"`
	SaleEndDate      *time.Time `json:"sale_end_date,omitempty" gorm:"column:sale_end_date"`
	Status           string     `json:"status" gorm:"column:status;type:varchar(255);not null"`
	BrandID          int64      `json:"brand_id" gorm:"column:brand_id;not null"`
	CategoryID       int64      `json:"category_id" gorm:"column:category_id;not null"`
	UserID           int64      `json:"user_id" gorm:"column:user_id;not null"`
	Version          int32      `json:"version" gorm:"column:version;not null;default:1"`
}

type ProductSKU struct {
	entity.BaseEntity
	SKU          string  `json:"sku" gorm:"column:sku;type:varchar(255);not null"`
	SKUSignature string  `json:"sku_signature" gorm:"column:sku_signature;type:varchar(255);not null;uniqueIndex"`
	Price        float64 `json:"price" gorm:"column:price;type:decimal(10,2);not null;default:0.00"`
	Status       string  `json:"status" gorm:"column:status;type:varchar(255);not null"`
	ProductID    int64   `json:"product_id" gorm:"column:product_id;not null"`
	Version      int32   `json:"version" gorm:"column:version;not null;default:1"`
}

type ProductSKUValue struct {
	ID                   int64 `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	ProductSKUID         int64 `json:"product_sku_id" gorm:"column:productSKU_id;not null"`
	ProductOptionValueID int64 `json:"product_option_value_id" gorm:"column:product_option_value_id;not null"`
}

type ProductInventory struct {
	entity.BaseEntity
	ProductID      int64 `json:"product_id,omitempty" gorm:"column:product_id;not null"`
	ProductSKUID   int64 `json:"product_sku_id,omitempty" gorm:"column:product_sku_id;not null"`
	AvailableStock int32 `json:"available_stock" gorm:"column:available_stock;not null;default:0"`
	ReservedStock  int32 `json:"reserved_stock" gorm:"column:reserved_stock;not null;default:0"`
	DamagedStock   int32 `json:"damaged_stock" gorm:"column:damaged_stock;not null;default:0"`
	TotalStock     int32 `json:"total_stock" gorm:"column:total_stock;-"`
}

type ProductReview struct {
	entity.BaseEntity
	ProductID          int64  `json:"product_id" gorm:"column:product_id;not null"`
	UserID             int64  `json:"user_id" gorm:"column:user_id;not null"`
	Rating             int32  `json:"rating" gorm:"column:rating;not null"`
	Title              string `json:"title,omitempty" gorm:"column:title;type:varchar(255)"`
	ReviewText         string `json:"review_text,omitempty" gorm:"column:review_text;type:text"`
	IsVerifiedPurchase bool   `json:"is_verified_purchase" gorm:"column:is_verified_purchase;default:false"`
	ReviewerName       string `json:"reviewer_name,omitempty" gorm:"column:reviewer_name;type:varchar(255)"`
	ReviewerEmail      string `json:"reviewer_email,omitempty" gorm:"column:reviewer_email;type:varchar(255)"`
}

func (Product) TableName() string {
	return "product"
}

func (ProductSKU) TableName() string {
	return "product_sku"
}

func (ProductSKUValue) TableName() string {
	return "product_sku_value"
}

func (ProductInventory) TableName() string {
	return "product_inventory"
}

func (ProductReview) TableName() string {
	return "product_review"
}
