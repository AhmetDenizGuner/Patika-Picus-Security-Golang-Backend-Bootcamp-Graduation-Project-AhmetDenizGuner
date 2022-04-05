package cart_item

import (
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/product"
)

type CartItemModel struct {
	product.ProductModel
	Quantity int
}
