package cart

import (
	"errors"
	"fmt"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/cart/cart_item"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/product"
	"gorm.io/gorm"
)

var (
	maxAllowedQtyPerProduct = 9
)

type Cart struct {
	gorm.Model
	TotalPrice float64
	UserID     uint
	Items      []cart_item.CartItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewCart(userId int) *Cart {
	cart := &Cart{
		UserID:     uint(userId),
		TotalPrice: 0,
	}
	cart.CalculateTotalPrice()
	return cart
}

func (c *Cart) CalculateTotalPrice() {

	c.TotalPrice = 0

	for _, item := range c.Items {
		c.TotalPrice += float64(item.Quantity) * item.Product.Price
	}

}

func (c *Cart) AddItem(product product.Product) (*cart_item.CartItem, error) {

	_, item := c.SearchItem(int(product.ID))
	if item != nil {
		return item, errors.New("Service: Item already added")
	}
	item = &cart_item.CartItem{
		ProductId: product.ID,
		CartId:    c.ID,
		Quantity:  1,
	}

	c.Items = append(c.Items, *item)
	c.CalculateTotalPrice()

	return item, nil
}

func (c *Cart) UpdateItem(itemProductId int, quantity int) (err error) {

	if index, _ := c.SearchItem(itemProductId); index != -1 {

		if quantity >= maxAllowedQtyPerProduct {
			return errors.New(fmt.Sprintf("You can't add more item. Item count can be less then %d", maxAllowedQtyPerProduct))
		}

		c.Items[index].Quantity = quantity
		c.CalculateTotalPrice()
	} else {
		return errors.New(fmt.Sprintf("Item can not found. ItemProductId : %s", itemProductId))
	}

	return
}

func (c *Cart) SearchItem(itemId int) (int, *cart_item.CartItem) {

	for i, n := range c.Items {
		if int(n.Product.ID) == itemId {
			return i, &n
		}
	}
	return -1, nil
}
