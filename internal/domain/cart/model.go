package cart

import "github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/cart/cart_item"

type CartModel struct {
	TotalPrice float64
	ItemModels []cart_item.CartItemModel
}

func NewCartModel(cart Cart) *CartModel {
	model := &CartModel{
		TotalPrice: cart.TotalPrice,
		ItemModels: []cart_item.CartItemModel{},
	}

	for _, item := range cart.Items {
		itemModel := cart_item.NewCartItemModel(item)
		model.ItemModels = append(model.ItemModels, *itemModel)
	}

	return model
}
