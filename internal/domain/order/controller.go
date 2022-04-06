package order

import (
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/config"
	jwtHelper "github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/pkg/jwt"
	"github.com/gin-gonic/gin"
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

func (c *OrderController) CompleteOrder(g *gin.Context) {
	userId := getUserIdFromAuthToken(g.GetHeader("Authorization"), c.appConfig.JwtSettings.SecretKey)

	err := c.orderService.CompleteOrderWithUserId(userId)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": err,
		})
	}

	//TODO response
	g.JSON(http.StatusCreated, "ok")
}

func (c *OrderController) CancelOrder(g *gin.Context) {

}

func (c *OrderController) ListOrders(g *gin.Context) {

}

func getUserIdFromAuthToken(token, secretKey string) int {
	decodedClaims := jwtHelper.VerifyToken(token, secretKey)
	userId := decodedClaims.UserId
	userIdInt, _ := strconv.Atoi(userId)
	return userIdInt
}
