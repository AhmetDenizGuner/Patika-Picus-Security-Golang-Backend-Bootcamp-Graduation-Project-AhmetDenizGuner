package user

import (
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/api/types"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/config"
	jwtHelper "github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/pkg/jwt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

type UserController struct {
	userService *UserService
	appConfig   *config.Configuration
}

func NewUserController(service *UserService, appConfig *config.Configuration) *UserController {
	return &UserController{
		userService: service,
		appConfig:   appConfig,
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
	var requestModel types.SigninRequest

	//check request body is correct form
	if err := g.ShouldBind(&requestModel); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Check your request body.",
		})
	}

	//service call
	user, err2 := c.userService.SignInService(requestModel)

	if err2 != nil {
		g.JSON(http.StatusNotFound, gin.H{
			"error_message": "User not found or password is not correct!",
		})
	}

	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"email":  user.Email,
		"iat":    time.Now().Unix(),
		"iss":    os.Getenv("ENV"),
		"exp":    time.Now().Add(time.Hour).Unix(),
		"roles":  user.Roles,
	})
	token := jwtHelper.GenerateToken(jwtClaims, c.appConfig.JwtSettings.SecretKey)
	g.JSON(http.StatusOK, token)

}
