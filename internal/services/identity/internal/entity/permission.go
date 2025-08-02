package entity

type Permission struct {
	ID          int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"not null" json:"name"`
	Description string `gorm:"not null" json:"description"`
}

func (p Permission) TableName() string {
	return "permissions"
}
