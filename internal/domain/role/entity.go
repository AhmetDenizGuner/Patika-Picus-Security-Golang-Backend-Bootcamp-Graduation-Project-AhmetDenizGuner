package role

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name string
}

func NewRole(name string) *Role {
	return &Role{
		Name: name,
	}
}
