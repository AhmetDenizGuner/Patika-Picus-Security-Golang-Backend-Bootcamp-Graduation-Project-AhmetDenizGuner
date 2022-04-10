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

//@BasePath /cart

func NewCartController(service *CartService, configuration *config.Configuration) *CartController {
	return &CartController{
		cartService: service,
		appConfig:   configuration,
	}
}

//AddCartItem godoc
//@Summary This endpoint used for adding new element to user cart
//@Accept  json
//@Tags Cart
//@Success 201
//@Param stock_code formData string true "stock code of adding element"
//Failure 400 shared.ApiErrorResponse
//@Router /cart/add-item [post]
//@Security ApiKeyAuth
//@param Authorization header string true "Authorization"
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

//UpdateCartItem godoc
//@Summary This endpoint used for deleting or update the item that is in cart already
//@Accept  json
//@Tags Cart
//@Success 204
//@Param stock_code formData string true "stock code of update element"
//@Param update_quantity formData int true "new cart quantity of element"
//Failure 400 shared.ApiErrorResponse
//@Router /cart/update-delete-item [put]
//@Security ApiKeyAuth
//@param Authorization header string true "Authorization"
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

//ShowCart godoc
//@Summary This endpoint used for see current cart
//@Accept  json
//@Tags Cart
//@Success 200
//Failure 400 shared.ApiErrorResponse
//@Router /cart/list [get]
//@Security ApiKeyAuth
//@param Authorization header string true "Authorization"
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
