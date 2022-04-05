package cart_item

import (
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/product"
	"gorm.io/gorm"
)

type CartItem struct {
	gorm.Model
	CartId    uint
	ProductId uint
	Product   product.Product `gorm:"foreignKey:ProductId"`
	Quantity  int
}
