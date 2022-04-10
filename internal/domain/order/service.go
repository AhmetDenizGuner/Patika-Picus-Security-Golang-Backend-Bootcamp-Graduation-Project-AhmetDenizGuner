package order

import (
	"errors"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/cart"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/order/order_item"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/product"
	"strconv"
)

type OrderService struct {
	repository          OrderRepository
	cartService         *cart.CartService
	orderItemRepository order_item.OrderItemRepository
	productService      *product.ProductService
}

//NewOrderService is constructor of OrderService
func NewOrderService(r OrderRepository, cartService *cart.CartService, orderItemRepository order_item.OrderItemRepository, productService *product.ProductService) *OrderService {
	return &OrderService{
		repository:          r,
		cartService:         cartService,
		orderItemRepository: orderItemRepository,
		productService:      productService,
	}
}

func (service *OrderService) CompleteOrderWithUserId(userId int) error {
	//get user cart
	cart, err := service.cartService.FetchCartByUserId(userId)
	if err != nil {
		return err
	}
	//check cart is empty
	if len(cart.Items) < 1 {
		return ErrOrderBasketEmpty
	}
	//update products quantity which are in basket
	errUpdQuant := service.productService.UpdateProductQuantityForOrder(cart.Items)
	if errUpdQuant != nil {
		return errUpdQuant
	}

	//create order
	order := NewOrder(cart)
	err1 := service.repository.Create(*order)

	if err1 != nil {
		return err1
	}

	//CLEAR BASKET
	service.cartService.ClearBasket(&cart)

	return nil
}

func (service *OrderService) listOrders(userId int) ([]OrderModel, error) {

	orders, err := service.repository.FindAllByUser(userId)
	orderModels := []OrderModel{}

	for _, order := range orders {
		orderModel := NewOrderModel(order)
		orderModels = append(orderModels, *orderModel)
	}

	if err != nil {
		return nil, err
	}

	return orderModels, nil
}

func (service *OrderService) cancelOrder(userId int, deleteID string) error {

	deleteIDInt, err := strconv.Atoi(deleteID)

	if err != nil {
		return err
	}

	order, err1 := service.repository.FindByID(deleteIDInt)

	if err1 != nil {
		return err1
	}

	if int(order.UserID) != userId {
		return errors.New("Credintials are not matched!")
	}

	err2 := service.repository.DeleteById(order.ID)

	if err2 != nil {
		return err2
	}

	for _, item := range order.Items {
		service.orderItemRepository.DeleteById(item.ID)
	}

	return nil
}
