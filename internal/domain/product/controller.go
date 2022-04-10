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

//ListProducts godoc
//@Summary This endpoint used for getting product list with pagination
//@Accept  json
//@Tags Product
//@Param page query int false "Page Index"
//@Param pageSize query int false "Page Size"
//@Success 200
//Failure 502 shared.ApiErrorResponse
//@Router /product/list [get]
//ListProducts get the products list with pagination
func (c *ProductController) ListProducts(g *gin.Context) {
	//prepare page and get products
	page := pagination.NewFromGinRequest(g, -1)
	products, err := c.productService.fetchProductsWithPagination(*page)

	if err != nil {
		log.Println(err.Error())
		g.JSON(http.StatusBadGateway, shared.ApiErrorResponse{
			IsSuccess:    false,
			ErrorMessage: "DB connection problem!",
		})
		g.Abort()
		return
	}
	//assign products
	page.Items = products

	g.JSON(http.StatusOK, shared.ApiOkResponse{
		IsSuccess: true,
		Message:   "ok",
		Data:      page})
}

//SearchProducts godoc
//@Summary This endpoint used for searching product with pagination
//@Accept  json
//@Tags Product
//@Param page query int false "Page Index"
//@Param pageSize query int false "Page Size"
//@Param searchKeyword query string false "word will be searched"
//@Success 200
//Failure 400 shared.ApiErrorResponse
//@Router /product/search{searchKeyword} [post]
//SearchProducts search the products with pagination
func (c *ProductController) SearchProducts(g *gin.Context) {
	//prepare page
	page := pagination.NewFromGinRequest(g, -1)
	//get search keyword from query
	searchItem := g.Query("searchKeyword")

	//search product
	products, err := c.productService.searchProductsWithPagination(*page, searchItem)

	if err != nil {
		log.Println(err.Error())
		g.JSON(http.StatusBadRequest, shared.ApiErrorResponse{
			IsSuccess:    false,
			ErrorMessage: err.Error(),
		})
		g.Abort()
		return
	}
	//assign products to page
	page.Items = products

	g.JSON(http.StatusOK, shared.ApiOkResponse{
		IsSuccess: true,
		Message:   "ok",
		Data:      page})
}

//AddProduct godoc
//@Summary This endpoint used for creating new product
//@Accept  json
//@Tags Product
//@Param requestModel body types.AddProductRequest "it is a new product model"
//@Success 201
//Failure 400 shared.ApiErrorResponse
//@Router /product/add [post]
//@Security ApiKeyAuth
//@param Authorization header string true "Authorization"
//AddProduct creates new product
func (c *ProductController) AddProduct(g *gin.Context) {
	var requestModel types.AddProductRequest

	//check request body is correct form
	if err := g.ShouldBind(&requestModel); err != nil {
		log.Println(err.Error())
		g.JSON(http.StatusBadRequest, shared.ApiErrorResponse{
			IsSuccess:    false,
			ErrorMessage: shared.GeneralErrorRequestBodyNotCorrect.Error(),
		})
		g.Abort()
		return
	}

	err := c.productService.addNewProduct(requestModel)

	if err != nil {
		log.Println(err.Error())
		g.JSON(http.StatusBadRequest, shared.ApiErrorResponse{
			IsSuccess:    false,
			ErrorMessage: err.Error(),
		})
		g.Abort()
		return
	}

	log.Println("ProductName: " + requestModel.Name + " added.")
	g.JSON(http.StatusCreated, shared.ApiOkResponse{
		IsSuccess: true,
		Message:   "ok",
		Data:      requestModel,
	})
}

//DeleteProduct godoc
//@Summary This endpoint used for remove the product fromDB
//@Accept  json
//@Tags Product
//@Param stock_code formData string true "stock code belongs product will be deleted"
//@Success 204
//Failure 400 shared.ApiErrorResponse
//@Router /product/delete [delete]
//@Security ApiKeyAuth
//@param Authorization header string true "Authorization"
//DeleteProduct remove the product fromDB
func (c *ProductController) DeleteProduct(g *gin.Context) {

	stockCode := g.PostForm("stock_code")

	err1 := c.productService.deleteProduct(stockCode)

	if err1 != nil {
		log.Println(err1.Error())
		g.JSON(http.StatusBadRequest, shared.ApiErrorResponse{
			IsSuccess:    false,
			ErrorMessage: err1.Error(),
		})
		g.Abort()
		return
	}

	g.JSON(http.StatusNoContent, shared.ApiOkResponse{
		IsSuccess: true,
		Message:   "deleted",
	})
}

//UpdateProduct godoc
//@Summary This endpoint used for updates product in DB
//@Accept  json
//@Tags Product
//@Param requestModel body types.AddProductRequest "it is an updated product model"
//@Success 204
//Failure 400 shared.ApiErrorResponse
//@Router /product/update [put]
//@Security ApiKeyAuth
//@param Authorization header string true "Authorization"
//UpdateProduct updates product in DB
func (c *ProductController) UpdateProduct(g *gin.Context) {
	var requestModel types.AddProductRequest

	//check request body is correct form
	if err := g.ShouldBind(&requestModel); err != nil {
		log.Println(err.Error())
		g.JSON(http.StatusBadRequest, shared.ApiErrorResponse{
			IsSuccess:    false,
			ErrorMessage: shared.GeneralErrorRequestBodyNotCorrect.Error(),
		})
		g.Abort()
		return
	}

	err := c.productService.updateProduct(requestModel)

	if err != nil {
		log.Println(err.Error())
		g.JSON(http.StatusBadRequest, shared.ApiErrorResponse{
			IsSuccess:    false,
			ErrorMessage: err.Error(),
		})
		g.Abort()
		return
	}

	log.Println("Product StockCode: " + requestModel.StockCode + "updated")
	g.JSON(http.StatusOK, shared.ApiOkResponse{
		IsSuccess: true,
		Message:   "updated",
		Data:      requestModel,
	})
}
