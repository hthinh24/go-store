package entity

import (
	"github.com/hthinh24/go-store/internal/pkg/entity"
	"time"
)

type User struct {
	entity.BaseEntity
	Email        string    `json:"email" gorm:"column:email;unique;not null"`
	Password     string    `json:"password" gorm:"column:password"`
	ProviderID   string    `json:"provider_id" gorm:"column:provider_id;not null"`
	ProviderName string    `json:"provider_name" gorm:"column:provider_name;not null"`
	LastName     string    `json:"last_name" gorm:"column:last_name;not null"`
	FirstName    string    `json:"first_name" gorm:"column:first_name;not null"`
	Avatar       string    `json:"avatar" gorm:"column:avatar"`
	Gender       string    `json:"gender" gorm:"column:gender;not null"`
	PhoneNumber  string    `json:"phone_number" gorm:"column:phone_number;not null"`
	DateOfBirth  time.Time `json:"date_of_birth" gorm:"column:date_of_birth;not null"`
	Status       string    `json:"status" gorm:"column:status;not null"`
}

type UserRoles struct {
	ID     int64 `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID int64 `json:"user_id" gorm:"column:user_id;not null"`
	RoleID int64 `json:"role_id" gorm:"column:role_id;not null"`
}

func (u User) TableName() string {
	return "users"
}
