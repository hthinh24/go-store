package entity

type Category struct {
	ID               int64  `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Name             string `json:"name" gorm:"column:name;type:varchar(255);not null;uniqueIndex"`
	Description      string `json:"description" gorm:"column:description;type:varchar(255);not null"`
	ParentCategoryID *int64 `json:"parent_category_id,omitempty" gorm:"column:parent_category_id"`
}

func (Category) TableName() string {
	return "category"
}
