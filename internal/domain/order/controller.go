package order

import (
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/config"
	jwtHelper "github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/pkg/jwt"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/shared"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type OrderController struct {
	orderService *OrderService
	appConfig    *config.Configuration
}

func NewOrderController(service *OrderService, configuration *config.Configuration) *OrderController {
	return &OrderController{
		orderService: service,
		appConfig:    configuration,
	}
}

//CompleteOrder crates order from items that is in basket and clear the basket
func (c *OrderController) CompleteOrder(g *gin.Context) {
	userId := getUserIdFromAuthToken(g.GetHeader("Authorization"), c.appConfig.JwtSettings.SecretKey)

	err := c.orderService.CompleteOrderWithUserId(userId)

	if err != nil {
		log.Println(err.Error())
		g.JSON(http.StatusBadRequest, shared.ApiErrorResponse{
			IsSuccess:    false,
			ErrorMessage: err.Error(),
		})
		g.Abort()
		return
	}

	g.JSON(http.StatusCreated, shared.ApiOkResponse{
		IsSuccess: true,
		Message:   "ok",
	})
}

//CancelOrder cancel order if it is not too old
func (c *OrderController) CancelOrder(g *gin.Context) {

	userId := getUserIdFromAuthToken(g.GetHeader("Authorization"), c.appConfig.JwtSettings.SecretKey)
	deleteID := g.PostForm("order_delete_id")

	err := c.orderService.cancelOrder(userId, deleteID)

	if err != nil {
		log.Println(err.Error())
		g.JSON(http.StatusBadRequest, shared.ApiErrorResponse{
			IsSuccess:    false,
			ErrorMessage: err.Error(),
		})
		g.Abort()
		return
	}

	g.JSON(http.StatusNoContent, shared.ApiOkResponse{
		IsSuccess: true,
		Message:   "order canceled",
	})

}

//ListOrders gets the orders except canceled orders
func (c *OrderController) ListOrders(g *gin.Context) {

	userId := getUserIdFromAuthToken(g.GetHeader("Authorization"), c.appConfig.JwtSettings.SecretKey)

	orders, err := c.orderService.listOrders(userId)

	if err != nil {
		log.Println(err.Error())
		g.JSON(http.StatusBadRequest, shared.ApiErrorResponse{
			IsSuccess:    false,
			ErrorMessage: err.Error(),
		})
		g.Abort()
		return
	}

	g.JSON(http.StatusOK, shared.ApiOkResponse{
		IsSuccess: true,
		Message:   "ok",
		Data:      orders,
	})
}

//getUserIdFromAuthToken can be moved from this class
func getUserIdFromAuthToken(token, secretKey string) int {
	decodedClaims := jwtHelper.VerifyToken(token, secretKey)
	userId := decodedClaims.UserId
	userIdInt, _ := strconv.Atoi(userId)
	return userIdInt
}
