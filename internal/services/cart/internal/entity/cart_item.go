package entity

import (
	"github.com/hthinh24/go-store/internal/pkg/entity"
)

type CartItem struct {
	entity.BaseEntity
	CartID       int64   `json:"cart_id" gorm:"column:cart_id;not null"`
	ProductID    int64   `json:"product_id" gorm:"column:product_id;not null"`
	ProductSKUID int64   `json:"product_sku_id" gorm:"column:product_sku_id;not null"`
	Quantity     int     `json:"quantity" gorm:"column:quantity;not null;default:1"`
	UnitPrice    float64 `json:"unit_price" gorm:"column:unit_price;not null;type:decimal(10,2)"`
	TotalPrice   float64 `json:"total_price" gorm:"column:total_price;type:decimal(10,2);<-:false"`
	Status       string  `json:"status" gorm:"column:status;not null;default:ACTIVE"`
}

func (ci CartItem) TableName() string {
	return "cart_item"
}
