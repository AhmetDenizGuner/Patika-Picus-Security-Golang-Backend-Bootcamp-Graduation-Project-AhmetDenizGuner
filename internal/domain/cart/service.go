package cart

import (
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/cart/cart_item"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/product"
	"strconv"
)

type CartService struct {
	repository         CartRepository
	productService     product.ProductService
	cartItemRepository cart_item.CartItemRepository
}

//NewCartService is constructor of CategoryService
func NewCartService(r CartRepository, productService product.ProductService, cartItemRepository cart_item.CartItemRepository) *CartService {
	return &CartService{
		repository:         r,
		productService:     productService,
		cartItemRepository: cartItemRepository,
	}
}

//addItem adds to new item to cart
func (service *CartService) addItem(stockCode string, userID int) error {

	cart, err := service.repository.FindByUserId(userID)

	if err != nil {
		return err
	}

	product, err1 := service.productService.FetchBySKU(stockCode)

	if err1 != nil {
		return err1
	}

	_, err2 := cart.AddItem(product)

	if err2 != nil {
		return err2
	}

	err3 := service.repository.Update(&cart)

	if err3 != nil {
		return err3
	}

	return nil
}

//updateCartItem update quantity of cart item or delete cart item according to parameters
func (service *CartService) updateCartItem(stockCode, stockQuantity string, userID int) error {

	//get user cart
	cart, err := service.repository.FindByUserId(userID)

	if err != nil {
		return err
	}
	//check is integer
	stockQuantityInt, err4 := strconv.Atoi(stockQuantity)

	if err4 != nil {
		return err4
	}

	product, err1 := service.productService.FetchBySKU(stockCode)

	if err1 != nil {
		return err1
	}

	if stockQuantityInt >= 0 {
		err2 := cart.UpdateItem(int(product.ID), stockQuantityInt)
		if err2 != nil {
			return err2
		}
	} else {
		return ErrCartItemQuantityNegative
	}
	//update general cart info
	err3 := service.repository.Update(&cart)

	if err3 != nil {
		return err3
	}
	//update or delete item
	for _, item := range cart.Items {
		if item.Quantity == 0 {
			service.cartItemRepository.DeleteById(item.ID)
		} else {
			service.cartItemRepository.Update(&item)
		}
	}

	return nil
}

//fetchCartModelByUserID it returns cart model for show
func (service *CartService) fetchCartModelByUserID(userID int) (CartModel, error) {

	cart, err := service.repository.FindByUserId(userID)

	if err != nil {
		return CartModel{}, err
	}

	cartModel := NewCartModel(cart)

	return *cartModel, nil
}

//FetchCartByUserId it returns cart model for complete
func (service *CartService) FetchCartByUserId(UserID int) (Cart, error) {
	cart, err := service.repository.FindByUserId(UserID)

	if err != nil {
		return Cart{}, err
	}

	return cart, nil

}

func (service *CartService) CreateDbSchema() {

	tableExist := service.repository.db.Migrator().HasTable(&Cart{})

	if !tableExist {
		service.repository.MigrateTable()
	}

}

func (service *CartService) CreateUserCart(id int) {
	newCart := NewCart(id)
	service.repository.Create(newCart)
}

func (service *CartService) ClearBasket(cart *Cart) {
	for _, item := range cart.Items {
		service.cartItemRepository.DeleteById(item.ID)
	}
}
