package user

import (
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/api/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	userService *UserService
}

func NewUserController(service *UserService) *UserController {
	return &UserController{
		userService: service,
	}
}

func (c *UserController) SignUp(g *gin.Context) {
	var requestModel types.SignupRequest

	//check request body is correct form
	if err := g.ShouldBind(&requestModel); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Check your request body.",
		})
	}

	//service call
	err2 := c.userService.SignupService(requestModel)

	if err2 != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Check your form of inputs. Error: " + err2.Error(),
		})
	}

	g.JSON(http.StatusCreated, requestModel)

}

func (c *UserController) SignIn(g *gin.Context) {

}
