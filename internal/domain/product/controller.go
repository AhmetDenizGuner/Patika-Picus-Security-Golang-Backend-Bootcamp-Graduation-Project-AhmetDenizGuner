package product

import (
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/api/types"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/pkg/pagination"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/shared"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ProductController struct {
	productService *ProductService
}

func NewProductController(service *ProductService) *ProductController {
	return &ProductController{
		productService: service,
	}
}

func (c *ProductController) ListProducts(g *gin.Context) {
	page := pagination.NewFromGinRequest(g, -1)

	products, err := c.productService.fetchProductsWithPagination(*page)

	if err != nil {
		g.JSON(http.StatusBadGateway, gin.H{
			"error_message": "DB connection problem!",
		})
		g.Abort()
		return
	}

	page.Items = products

	g.JSON(http.StatusOK, page)
}

func (c *ProductController) SearchProducts(g *gin.Context) {
	page := pagination.NewFromGinRequest(g, -1)

	searchItem := g.Query("searchKeyword")

	products, err := c.productService.searchProductsWithPagination(*page, searchItem)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": err,
		})
		g.Abort()
		return
	}

	page.Items = products

	g.JSON(http.StatusOK, page)
}

func (c *ProductController) AddProduct(g *gin.Context) {
	var requestModel types.AddProductRequest

	//check request body is correct form
	if err := g.ShouldBind(&requestModel); err != nil {
		log.Println(err.Error())
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": shared.GeneralErrorRequestBodyNotCorrect,
		})
		g.Abort()
		return
	}

	err := c.productService.addNewProduct(requestModel)

	if err != nil {
		log.Println(err.Error())
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": err.Error(),
		})
		g.Abort()
		return
	}

	log.Println("ProductName: " + requestModel.Name + " added.")
	g.JSON(http.StatusCreated, requestModel)
}

func (c *ProductController) DeleteProduct(g *gin.Context) {

	stockCode := g.PostForm("stock_code")

	err1 := c.productService.deleteProduct(stockCode)

	if err1 != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": err1.Error(),
		})
		g.Abort()
		return
	}

	g.JSON(http.StatusOK, "sucsesfully deleted")
}

func (c *ProductController) UpdateProduct(g *gin.Context) {
	var requestModel types.AddProductRequest

	//check request body is correct form
	if err := g.ShouldBind(&requestModel); err != nil {
		log.Println(err.Error())
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": shared.GeneralErrorRequestBodyNotCorrect,
		})
		g.Abort()
		return
	}

	err := c.productService.updateProduct(requestModel)

	if err != nil {
		log.Println(err.Error())
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": err.Error(),
		})
		g.Abort()
		return
	}

	log.Println("Product StockCode: " + requestModel.StockCode + "updated")
	g.JSON(http.StatusOK, "updated")
}
