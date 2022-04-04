package category

import (
	"fmt"
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

	//responseModel is category root of category tree that keeps child category and this goes recursively
	responseModel, err := c.categoryService.FetchCategoryRootForList()

	if err != nil {
		g.JSON(http.StatusInsufficientStorage, gin.H{
			"error_message": ErrCategoryDataNotFound,
		})
	}

	g.JSON(http.StatusOK, responseModel)
}

func (c *CategoryController) AddCategoryFromCSV(g *gin.Context) {

	file, err := g.FormFile("file")
	if err != nil {
		g.JSON(http.StatusUnsupportedMediaType, gin.H{
			"error_message": ErrUploadDataNotFoundOrNotSupported,
		})
		log.Fatal(err)
	}

	if strings.Compare(string(file.Filename[len(file.Filename)-4:]), ".csv") != 0 {
		g.JSON(http.StatusUnsupportedMediaType, gin.H{
			"error_message": ErrUploadDataNotFoundOrNotSupported,
		})
		log.Fatal(err)
	}

	err = g.SaveUploadedFile(file, "saved/"+file.Filename)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error_message": "File cannot saved",
		})
		log.Fatal(err)
	}

	c.categoryService.AddBulkCategory()

	g.JSON(http.StatusCreated, fmt.Sprintf("'%s' uploaded!", file.Filename))

}
