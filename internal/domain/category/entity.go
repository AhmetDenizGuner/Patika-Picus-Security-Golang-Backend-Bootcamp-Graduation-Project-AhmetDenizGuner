package category

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name          string
	ParentID      uint
	Parent        *Category  `gorm:"foreignKey:ParentID"`
	SubCategories []Category `gorm:"many2many:manual_auto"`
}
