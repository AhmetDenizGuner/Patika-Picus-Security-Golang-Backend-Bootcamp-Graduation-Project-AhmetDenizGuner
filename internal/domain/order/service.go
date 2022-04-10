package order

import (
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

//CompleteOrderWithUserId crates order from items that is in basket and clear the basket
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
	//prepare update product input service
	var productList []product.Product = []product.Product{}
	var orderAmountList []int = []int{}
	for _, item := range cart.Items {
		productList = append(productList, item.Product)
		orderAmountList = append(orderAmountList, item.Quantity)
	}
	//update products quantity which are in basket
	errUpdQuant := service.productService.UpdateProductQuantityForOrder(productList, orderAmountList)
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

//listOrders gets the orders except canceled orders
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

//cancelOrder cancel order if it is not too old
func (service *OrderService) cancelOrder(userId int, deleteID string) error {
	//check parameter is ok
	deleteIDInt, err := strconv.Atoi(deleteID)
	if err != nil {
		return err
	}
	//find order
	order, err1 := service.repository.FindByID(deleteIDInt)

	if err1 != nil {
		return err1
	}

	//check cancel day
	if !order.isCancelable() {
		return ErrOrderCannotBeCanceled
	}
	//extra validation for order belongs to user --> it will be unnecesary
	if int(order.UserID) != userId {
		return ErrUserNotAuth
	}
	//cancel order
	err2 := service.repository.DeleteById(order.ID)

	if err2 != nil {
		return err2
	}
	//prepare update product input service
	var productList []product.Product = []product.Product{}
	var orderAmountList []int = []int{}
	for _, item := range order.Items {
		productList = append(productList, item.Product)
		orderAmountList = append(orderAmountList, item.Quantity)
	}
	//update products quantity which are in canceled order
	errUpdQuant := service.productService.UpdateProductQuantityForOrder(productList, orderAmountList)
	if errUpdQuant != nil {
		return errUpdQuant
	}
	//delete order items
	for _, item := range order.Items {
		service.orderItemRepository.DeleteById(item.ID)
	}

	return nil
}
