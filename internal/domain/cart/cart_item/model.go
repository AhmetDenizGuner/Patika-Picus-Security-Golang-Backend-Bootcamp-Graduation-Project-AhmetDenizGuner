package cart_item

import (
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/product"
)

type CartItemModel struct {
	Product  product.ProductModel
	Quantity int
}

func NewCartItemModel(item CartItem) *CartItemModel {
	model := &CartItemModel{
		Quantity: item.Quantity,
	}
	model.Product = *product.NewProductModel(item.Product)

	return model
}
