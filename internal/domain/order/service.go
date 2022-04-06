package order

import "github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/cart"

type OrderService struct {
	repository  OrderRepository
	cartService *cart.CartService
}

//NewOrderService is constructor of OrderService
func NewOrderService(r OrderRepository, cartService *cart.CartService) *OrderService {
	return &OrderService{
		repository:  r,
		cartService: cartService,
	}
}

func (service *OrderService) CompleteOrderWithUserId(userId int) error {

	cart, err := service.cartService.FetchCartByUserId(userId)

	if err != nil {
		return err
	}

	order := NewOrder(cart)

	err1 := service.repository.Create(*order)

	if err1 != nil {
		return err1
	}

	return nil
}
