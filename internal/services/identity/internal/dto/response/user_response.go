package response

import "time"

type UserResponse struct {
	ID          int64     `json:"id"`
	Email       string    `json:"email"`
	LastName    string    `json:"last_name"`
	FirstName   string    `json:"first_name"`
	Avatar      string    `json:"avatar"`
	Gender      string    `json:"gender" gorm:"column:gender;not null"`
	PhoneNumber string    `json:"phone_number" gorm:"column:phone_number;not null"`
	DateOfBirth time.Time `json:"date_of_birth" gorm:"column:date_of_birth;not null"`
	Status      string    `json:"status" gorm:"column:status;not null"`
}
