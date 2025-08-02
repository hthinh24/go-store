package entity

type Permission struct {
	ID          int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name" gorm:"not null"`
	Description string ` json:"description" gorm:"not null"`
}

func (p Permission) TableName() string {
	return "permissions"
}
