package cart

import (
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/config"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/cart/cart_item"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/product"
	jwtHelper "github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/pkg/jwt"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/shared"
	"github.com/gin-gonic/gin"
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

func (c *CartController) AddCartItem(g *gin.Context) {
	var requestModel product.ProductModel

	//check request body is correct form
	if err := g.ShouldBind(&requestModel); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": shared.GeneralErrorRequestBodyNotCorrect,
		})
	}

	userId := getUserIdFromAuthToken(g.GetHeader("Authorization"), c.appConfig.JwtSettings.SecretKey)

	err := c.cartService.addItem(requestModel, userId)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": err,
		})
	}

	//TODO - check response
	g.JSON(http.StatusCreated, requestModel)
}

func (c *CartController) UpdateCartItem(g *gin.Context) {
	var requestModel cart_item.CartItemModel

	//check request body is correct form
	if err := g.ShouldBind(&requestModel); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": shared.GeneralErrorRequestBodyNotCorrect,
		})
	}

	userId := getUserIdFromAuthToken(g.GetHeader("Authorization"), c.appConfig.JwtSettings.SecretKey)

	err := c.cartService.updateCartItem(requestModel, userId)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": err,
		})
	}

	g.JSON(http.StatusOK, requestModel)

}

func (c *CartController) ShowCart(g *gin.Context) {
	userId := getUserIdFromAuthToken(g.GetHeader("Authorization"), c.appConfig.JwtSettings.SecretKey)

	cartModel, err := c.cartService.fetchCartModelByUserID(userId)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": err,
		})
	}

	g.JSON(http.StatusOK, cartModel)

}

func getUserIdFromAuthToken(token, secretKey string) int {
	decodedClaims := jwtHelper.VerifyToken(token, secretKey)
	userId := decodedClaims.UserId
	userIdInt, _ := strconv.Atoi(userId)
	return userIdInt
}
