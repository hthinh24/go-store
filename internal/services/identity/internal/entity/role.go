package entity

type Role struct {
	ID          int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name" gorm:"column:name;not null"`
	Description string `json:"description" gorm:"column:description;not null"`
}

type RolePermissions struct {
	ID           int64 `json:"id" gorm:"primaryKey;autoIncrement" `
	PermissionID int64 `json:"permissions_id" gorm:"column:permission_id;not null" `
	RoleID       int64 `json:"role_id" gorm:"column:role_id;not null"`
}

func (r Role) TableName() string {
	return "roles"
}
