package category

import (
	"fmt"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/pkg/pagination"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/shared"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

type CategoryController struct {
	categoryService *CategoryService
}

func NewCategoryController(service *CategoryService) *CategoryController {
	return &CategoryController{
		categoryService: service,
	}
}

//CategoryList get the list of categories
func (c *CategoryController) CategoryList(g *gin.Context) {

	//prepare page and get categories
	page := pagination.NewFromGinRequest(g, -1)
	categories, err := c.categoryService.fetchCategoriesWithPagination(*page)

	if err != nil {
		log.Println(err.Error())
		g.JSON(http.StatusInsufficientStorage, shared.ApiErrorResponse{
			IsSuccess:    false,
			ErrorMessage: ErrCategoryDataNotFound.Error(),
		})
		g.Abort()
		return
	}

	//assign page
	page.Items = categories

	g.JSON(http.StatusOK, shared.ApiOkResponse{
		IsSuccess: true,
		Message:   "ok",
		Data:      page,
	})
}

//AddCategoryFromCSV get CSV from form-data
func (c *CategoryController) AddCategoryFromCSV(g *gin.Context) {

	//get form data
	file, err := g.FormFile("file")
	if err != nil {
		log.Println(err.Error())
		g.JSON(http.StatusUnsupportedMediaType, shared.ApiErrorResponse{
			IsSuccess:    false,
			ErrorMessage: ErrUploadDataNotFoundOrNotSupported.Error()})
		g.Abort()
		return

	}
	//check file format is csv
	if strings.Compare(string(file.Filename[len(file.Filename)-4:]), ".csv") != 0 {
		log.Println(ErrUploadDataNotFoundOrNotSupported.Error())
		g.JSON(http.StatusUnsupportedMediaType, shared.ApiErrorResponse{
			IsSuccess:    false,
			ErrorMessage: ErrUploadDataNotFoundOrNotSupported.Error()})
		g.Abort()
		return
	}
	//save file
	err = g.SaveUploadedFile(file, "../../resources/uploaded"+file.Filename)
	if err != nil {
		log.Println(err.Error())
		g.JSON(http.StatusInternalServerError, shared.ApiErrorResponse{
			IsSuccess:    false,
			ErrorMessage: "File cannot be saved"})
		g.Abort()
		return
	}
	//add categories
	err2 := c.categoryService.AddBulkCategory("../../resources/uploaded" + file.Filename)
	if err2 != nil {
		log.Println(err2.Error())
		g.JSON(http.StatusBadRequest, shared.ApiErrorResponse{
			IsSuccess:    false,
			ErrorMessage: err2.Error()})
		g.Abort()
		return
	}
	//ok
	g.JSON(http.StatusCreated, shared.ApiOkResponse{
		IsSuccess: true,
		Message:   fmt.Sprintf("'%s' uploaded!", file.Filename),
	})

}
