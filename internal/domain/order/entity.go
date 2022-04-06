package order

import (
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/cart"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/order/order_item"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	TotalPrice float64
	UserID     uint
	Items      []order_item.OrderItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewOrder(cart cart.Cart) *Order {

	order := &Order{
		TotalPrice: cart.TotalPrice,
		UserID:     cart.UserID,
		Items:      []order_item.OrderItem{},
	}

	for _, item := range cart.Items {
		orderItem := order_item.NewOrderItem(item)
		order.Items = append(order.Items, *orderItem)
	}

	return order

}
