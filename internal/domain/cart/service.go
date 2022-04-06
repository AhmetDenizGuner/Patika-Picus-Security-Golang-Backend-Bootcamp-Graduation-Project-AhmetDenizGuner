package cart

import (
	"errors"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/cart/cart_item"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/product"
	"gorm.io/gorm"
)

type CartService struct {
	repository     CartRepository
	productService product.ProductService
}

//NewCartService is constructor of CategoryService
func NewCartService(r CartRepository) *CartService {
	return &CartService{
		repository: r,
	}
}

func (service *CartService) addItem(productModel product.ProductModel, userID int) error {

	cart, err := service.repository.FindByID(userID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		newCart := NewCart(userID)
		service.repository.Create(*newCart)
		cart = *newCart
	}

	product, err1 := service.productService.FetchBySKU(productModel.StockCode)

	if err1 != nil {
		return err1
	}

	_, err2 := cart.AddItem(product)

	if err2 != nil {
		return err2
	}

	err3 := service.repository.Update(cart)

	if err3 != nil {
		return err3
	}

	//TODO - check items created at DB

	return nil

}

//updateCartItem update quantity of cart item or delete cart item according to parameters
func (service *CartService) updateCartItem(cartItemModel cart_item.CartItemModel, userID int) error {

	cart, err := service.repository.FindByID(userID)

	//TODO create cart when user signup
	if errors.Is(err, gorm.ErrRecordNotFound) {
		newCart := NewCart(userID)
		service.repository.Create(*newCart)
		cart = *newCart
	}

	product, err1 := service.productService.FetchBySKU(cartItemModel.Product.StockCode)

	if err1 != nil {
		return err1
	}

	if cartItemModel.Quantity == 0 { //DELETE
		err2 := cart.RemoveItem(int(product.ID))
		if err2 != nil {
			return err2
		}
	} else if cartItemModel.Quantity > 0 { // UPDATE CART QUANTITY
		err2 := cart.UpdateItem(int(product.ID), cartItemModel.Quantity)
		if err2 != nil {
			return err2
		}
	} else {
		return ErrCartItemQuantityNegative
	}

	err3 := service.repository.Update(cart)

	if err3 != nil {
		return err3
	}

	return nil
}
func (service *CartService) fetchCartModelByUserID(userID int) (CartModel, error) {

	cart, err := service.repository.FindByUserId(userID)

	if err != nil {
		return CartModel{}, err
	}

	cartModel := NewCartModel(cart)

	return *cartModel, nil
}

func (service *CartService) FetchCartByUserId(UserID int) (Cart, error) {
	cart, err := service.repository.FindByUserId(UserID)

	if err != nil {
		return Cart{}, err
	}

	return cart, nil

}
