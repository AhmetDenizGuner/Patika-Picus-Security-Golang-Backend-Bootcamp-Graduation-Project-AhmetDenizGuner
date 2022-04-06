package cart

import (
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/cart/cart_item"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/product"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var (
	maxAllowedForBasket             = 20
	maxAllowedQtyPerProduct         = 9
	minCartAmountForOrder   float64 = 50
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

func NewCartWithItems(userId int, items []cart_item.CartItem) *Cart {
	cart := &Cart{
		UserID:     uint(userId),
		TotalPrice: 0,
		Items:      items,
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
	/*if quantity >= maxAllowedQtyPerProduct {
		return nil, errors.Errorf("You can't add more this item to your basket. Maximum allowed item count is %d", maxAllowedQtyPerProduct)
	}
	if (len(c.Items) + quantity) >= maxAllowedForBasket {
		return nil, errors.Errorf("You can't add more item to your basket. Maximum allowed basket item count is %d", maxAllowedForBasket)
	} */

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

	if index, item := c.SearchItem(itemProductId); index != -1 {

		if quantity >= maxAllowedQtyPerProduct {
			return errors.Errorf("You can't add more item. Item count can be less then %d", maxAllowedQtyPerProduct)
		}

		item.Quantity = quantity
	} else {
		return errors.Errorf("Item can not found. ItemProductId : %s", itemProductId)
	}

	return
}

func (c *Cart) RemoveItem(itemProductId int) (err error) {

	if index, _ := c.SearchItem(itemProductId); index != -1 {
		c.Items = append(c.Items[:index], c.Items[index+1:]...)
	} else {
		return ErrNotFound
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
