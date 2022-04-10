package order

import (
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/cart"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/order/order_item"
	"gorm.io/gorm"
	"time"
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

func (o *Order) isCancelable() bool {
	orderDate := o.CreatedAt
	now := time.Now()
	if now.Sub(orderDate) < time.Hour*24*14 {
		return true
	}
	return false
}
