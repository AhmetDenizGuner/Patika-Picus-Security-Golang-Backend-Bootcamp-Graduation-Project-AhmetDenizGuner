package cart

import (
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/config"
	jwtHelper "github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/pkg/jwt"
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

	stockCode := g.PostForm("stock_code")

	userId := getUserIdFromAuthToken(g.GetHeader("Authorization"), c.appConfig.JwtSettings.SecretKey)
	err := c.cartService.addItem(stockCode, userId)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": err.Error(),
		})
		g.Abort()
		return
	}

	//TODO - check response
	g.JSON(http.StatusCreated, stockCode)
}

//TODO create cart items check
func (c *CartController) UpdateCartItem(g *gin.Context) {
	stockCode := g.PostForm("stock_code")
	updateQuantity := g.PostForm("update_quantity")

	userId := getUserIdFromAuthToken(g.GetHeader("Authorization"), c.appConfig.JwtSettings.SecretKey)

	err := c.cartService.updateCartItem(stockCode, updateQuantity, userId)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": err.Error(),
		})
	}

	g.JSON(http.StatusOK, "updated")

}

func (c *CartController) ShowCart(g *gin.Context) {
	userId := getUserIdFromAuthToken(g.GetHeader("Authorization"), c.appConfig.JwtSettings.SecretKey)

	cartModel, err := c.cartService.fetchCartModelByUserID(userId)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": err.Error(),
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
