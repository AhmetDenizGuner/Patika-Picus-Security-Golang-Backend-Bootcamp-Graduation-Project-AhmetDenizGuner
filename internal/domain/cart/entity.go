package cart

import (
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/cart/cart_item"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	TotalPrice float64
	UserID     uint
	Items      []cart_item.CartItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
