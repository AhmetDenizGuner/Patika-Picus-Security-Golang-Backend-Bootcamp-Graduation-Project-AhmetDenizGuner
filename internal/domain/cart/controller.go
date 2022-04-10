package cart

import (
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/config"
	jwtHelper "github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/pkg/jwt"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/shared"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type CartController struct {
	cartService *CartService
	appConfig   *config.Configuration
}

func NewCartController(service *CartService, configuration *config.Configuration) *CartController {
	return &CartController{
		cartService: service,
		appConfig:   configuration,
	}
}

//AddCartItem adds new item to cart with 1 qauntity
func (c *CartController) AddCartItem(g *gin.Context) {

	stockCode := g.PostForm("stock_code")

	userId := getUserIdFromAuthToken(g.GetHeader("Authorization"), c.appConfig.JwtSettings.SecretKey)
	err := c.cartService.addItem(stockCode, userId)

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
		Message:   "item added",
		Data:      stockCode})
}

//UpdateCartItem update or delete item from cart
func (c *CartController) UpdateCartItem(g *gin.Context) {
	stockCode := g.PostForm("stock_code")
	updateQuantity := g.PostForm("update_quantity")

	userId := getUserIdFromAuthToken(g.GetHeader("Authorization"), c.appConfig.JwtSettings.SecretKey)

	err := c.cartService.updateCartItem(stockCode, updateQuantity, userId)

	if err != nil {
		log.Println(err.Error())
		g.JSON(http.StatusBadRequest, shared.ApiErrorResponse{
			IsSuccess:    false,
			ErrorMessage: err.Error(),
		})
	}

	g.JSON(http.StatusNoContent, shared.ApiOkResponse{
		IsSuccess: true,
		Message:   "item added",
		Data:      stockCode})

}

//ShowCart gets active cart for user
func (c *CartController) ShowCart(g *gin.Context) {
	userId := getUserIdFromAuthToken(g.GetHeader("Authorization"), c.appConfig.JwtSettings.SecretKey)

	cartModel, err := c.cartService.fetchCartModelByUserID(userId)

	if err != nil {
		log.Println(err.Error())
		g.JSON(http.StatusBadRequest, shared.ApiErrorResponse{
			IsSuccess:    false,
			ErrorMessage: err.Error(),
		})
	}

	g.JSON(http.StatusOK, shared.ApiOkResponse{
		IsSuccess: true,
		Message:   "item added",
		Data:      cartModel})

}

//getUserIdFromAuthToken can be moved from this class
func getUserIdFromAuthToken(token, secretKey string) int {
	decodedClaims := jwtHelper.VerifyToken(token, secretKey)
	userId := decodedClaims.UserId
	userIdInt, _ := strconv.Atoi(userId)
	return userIdInt
}
