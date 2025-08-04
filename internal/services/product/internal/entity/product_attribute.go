package entity

import (
	"github.com/hthinh24/go-store/internal/pkg/entity"
)

type ProductAttribute struct {
	entity.BaseEntity
	Name string `json:"name" gorm:"column:name;type:varchar(255);not null"`
}

type ProductAttributeCategory struct {
	ID                 int32 `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	ProductAttributeID int32 `json:"product_attribute_id" gorm:"column:product_attribute_id;not null"`
	CategoryID         int32 `json:"category_id" gorm:"column:category_id;not null"`
}

type ProductAttributeValue struct {
	ID                 int64  `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Value              string `json:"value" gorm:"column:value;type:varchar(255);not null"`
	ProductAttributeID int64  `json:"product_attribute_id" gorm:"column:product_attribute_id;not null"`
}

type ProductProductAttributeValue struct {
	ID                      int64 `json:"id" gorm:"primaryKey;autoIncrement;column:ID"`
	ProductID               int64 `json:"product_id" gorm:"column:product_id;not null"`
	ProductAttributeValueID int64 `json:"product_attribute_value_id" gorm:"column:product_attribute_value_id;not null"`
	DisplayOrder            int32 `json:"display_order" gorm:"column:display_order;not null"`
}

func (ProductAttribute) TableName() string {
	return "product_attribute"
}

func (ProductAttributeCategory) TableName() string {
	return "product_attribute_category"
}

func (ProductAttributeValue) TableName() string {
	return "product_attribute_value"
}

func (ProductProductAttributeValue) TableName() string {
	return "product_product_attribute_value"
}
