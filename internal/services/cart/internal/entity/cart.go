package entity

import (
	"github.com/hthinh24/go-store/internal/pkg/entity"
)

type Cart struct {
	entity.BaseEntity
	UserID int64  `json:"user_id" gorm:"column:user_id;not null"`
	Status string `json:"status" gorm:"column:status;not null;default:ACTIVE"`
}

func (c Cart) TableName() string {
	return "cart"
}
