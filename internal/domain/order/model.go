package order

import (
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/order/order_item"
)

type OrderModel struct {
	ID         int
	TotalPrice float64
	Items      []order_item.OrderItemModel
}

func NewOrderModel(order Order) *OrderModel {
	model := &OrderModel{
		ID:         int(order.ID),
		TotalPrice: order.TotalPrice,
		Items:      []order_item.OrderItemModel{},
	}

	for _, item := range order.Items {
		ord_item := order_item.NewOrderItemModel(item)
		model.Items = append(model.Items, *ord_item)
	}

	return model
}
