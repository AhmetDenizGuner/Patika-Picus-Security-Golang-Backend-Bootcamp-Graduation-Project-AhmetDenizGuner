package cart

import (
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/product"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/shared"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CartController struct {
	cartService *CartService
}

func NewCartController(service *CartService) *CartController {
	return &CartController{
		cartService: service,
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

	//TODO - get user id from auth , pass to service

	c.cartService.addItem(requestModel)

	g.JSON(http.StatusCreated, requestModel)
}