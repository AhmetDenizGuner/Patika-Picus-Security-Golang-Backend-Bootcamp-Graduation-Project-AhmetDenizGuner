package category

import (
	"fmt"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/pkg/pagination"
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

func (c *CategoryController) CategoryList(g *gin.Context) {

	page := pagination.NewFromGinRequest(g, -1)
	categories, err := c.categoryService.fetchCategoriesWithPagination(*page)

	if err != nil {
		g.JSON(http.StatusInsufficientStorage, gin.H{
			"error_message": ErrCategoryDataNotFound,
		})
		g.Abort()
		return
	}

	page.Items = categories

	g.JSON(http.StatusOK, page)
}

func (c *CategoryController) AddCategoryFromCSV(g *gin.Context) {

	fmt.Println("Heloooooooooooooooooooooooooooooooooooooo11")

	file, err := g.FormFile("file")
	if err != nil {
		g.JSON(http.StatusUnsupportedMediaType, gin.H{
			"error_message": ErrUploadDataNotFoundOrNotSupported,
		})
		log.Fatal(err)
	}

	fmt.Println("Heloooooooooooooooooooooooooooooooooooooo22")

	if strings.Compare(string(file.Filename[len(file.Filename)-4:]), ".csv") != 0 {
		g.JSON(http.StatusUnsupportedMediaType, gin.H{
			"error_message": ErrUploadDataNotFoundOrNotSupported,
		})
		log.Fatal(err)
	}

	fmt.Println("Heloooooooooooooooooooooooooooooooooooooo33")

	err = g.SaveUploadedFile(file, ""+file.Filename)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error_message": "File cannot saved",
		})
		log.Fatal(err)
	}

	fmt.Println("Heloooooooooooooooooooooooooooooooooooooo44")

	c.categoryService.AddBulkCategory("" + file.Filename)

	g.JSON(http.StatusCreated, fmt.Sprintf("'%s' uploaded!", file.Filename))

}
