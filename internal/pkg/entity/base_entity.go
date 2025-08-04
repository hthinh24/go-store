package entity

import "time"

type BaseEntity struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedBy string    `json:"created_by" gorm:"column:created_by;not null;<-:create"`
	UpdatedBy string    `json:"updated_by" gorm:"column:updated_by;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
	Version   int32     `json:"version" gorm:"column:version;not null;default:1"`
}
