package entity

import "time"

type User struct {
	ID           int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Email        string    `json:"email" gorm:"column:email;unique;not null"`
	Password     string    `json:"password" gorm:"column:password"`
	ProviderID   string    `json:"provider_id" gorm:"column:providerID;not null"`
	ProviderName string    `json:"provider_name" gorm:"column:providerName;not null"`
	LastName     string    `json:"last_name" gorm:"column:last_name;not null"`
	FirstName    string    `json:"first_name" gorm:"column:first_name;not null"`
	Avatar       string    `json:"avatar" gorm:"column:avatar"`
	Gender       string    `json:"gender" gorm:"column:gender;not null"`
	PhoneNumber  string    `json:"phone_number" gorm:"column:phone_number;not null"`
	DateOfBirth  time.Time `json:"date_of_birth" gorm:"column:date_of_birth;not null"`
	Status       string    `json:"status" gorm:"column:status;not null"`
}

type UserHasRole struct {
	ID     int64 `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID int64 `gorm:"column:users_id;not null" json:"users_id"`
	RoleID int64 `gorm:"column:roles_id;not null" json:"roles_id"`
}

func (u User) TableName() string {
	return "users"
}
