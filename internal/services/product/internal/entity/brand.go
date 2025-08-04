package entity

import (
	"github.com/hthinh24/go-store/internal/pkg/entity"
)

type Brand struct {
	entity.BaseEntity
	Name        string `json:"name" gorm:"column:name;type:varchar(255);not null;uniqueIndex"`
	Description string `json:"description,omitempty" gorm:"column:description;type:text"`
}

func (Brand) TableName() string {
	return "brand"
}
