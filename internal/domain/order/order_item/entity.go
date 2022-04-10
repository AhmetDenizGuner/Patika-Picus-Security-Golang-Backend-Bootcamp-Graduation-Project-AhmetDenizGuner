package order_item

import (
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/cart/cart_item"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/product"
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	OrderId   uint
	ProductId uint
	Product   product.Product `gorm:"foreignKey:ProductId"`
	Quantity  int
}

func NewOrderItem(cartItem cart_item.CartItem) *OrderItem {
	item := &OrderItem{
		ProductId: cartItem.ProductId,
		Quantity:  cartItem.Quantity,
	}
	return item
}
