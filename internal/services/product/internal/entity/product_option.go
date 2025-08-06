package entity

import (
	"github.com/hthinh24/go-store/internal/pkg/entity"
)

type ProductOption struct {
	entity.BaseEntity
	Name string `json:"name" gorm:"column:name;type:varchar(255);not null"`
}

type ProductOptionValue struct {
	ID              int64  `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Value           string `json:"value" gorm:"column:value;type:varchar(255);not null"`
	ProductOptionID int64  `json:"product_option_id" gorm:"column:product_option_id;not null"`
}

type ProductOptionCombination struct {
	ID              int64 `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	ProductID       int64 `json:"product_id" gorm:"column:product_id;not null"`
	ProductOptionID int64 `json:"product_option_id" gorm:"column:product_option_id;not null"`
	DisplayOrder    int32 `json:"display_order" gorm:"column:display_order;not null"`
}

func (ProductOption) TableName() string {
	return "product_option"
}

func (ProductOptionValue) TableName() string {
	return "product_option_value"
}

func (ProductOptionCombination) TableName() string {
	return "product_option_combination"
}
