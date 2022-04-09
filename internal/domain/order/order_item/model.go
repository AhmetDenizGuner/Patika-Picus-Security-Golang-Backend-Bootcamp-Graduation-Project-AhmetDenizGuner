package order_item

import (
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/product"
)

type OrderItemModel struct {
	Id       int                  `json:"id"`
	Product  product.ProductModel `json:"product"`
	Quantity int                  `json:"quantity"`
}

func NewOrderItemModel(item OrderItem) *OrderItemModel {
	return &OrderItemModel{
		Id:       int(item.ID),
		Quantity: item.Quantity,
		Product:  *product.NewProductModel(item.Product),
	}
}
