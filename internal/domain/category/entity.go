package category

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name     string
	ParentID uint `gorm:"TYPE:integer REFERENCES table_category"`
	Parent   *Category
}

func NewCategory(name string) *Category {
	return &Category{
		Name: name,
	}
}

func NewCategoryWithParent(name string, parentID int) *Category {
	return &Category{
		Name:     name,
		ParentID: uint(parentID),
	}
}
