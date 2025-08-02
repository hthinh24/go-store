package entity

type Role struct {
	ID          int64  `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name        string `json:"name" gorm:"column:name;not null"`
	Description string `json:"description" gorm:"column:description;not null"`
}

type RoleHasPermission struct {
	ID           int64 `gorm:"primaryKey;autoIncrement" json:"id"`
	PermissionID int64 `gorm:"column:permissions_id;not null" json:"permissions_id"`
	RoleID       int64 `gorm:"column:roles_id;not null" json:"roles_id"`
}

func (r Role) TableName() string {
	return "roles"
}
