package entity

import "github.com/hthinh24/go-store/internal/pkg/entity"

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
	ProductSKUID         int64 `json:"product_sku_id" gorm:"column:product_sku_id;not null"`
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
	Version        int32 `json:"version" gorm:"column:version;not null;default:1"`
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
