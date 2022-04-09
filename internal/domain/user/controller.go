package user

import (
	"fmt"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/api/types"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/config"
	jwtHelper "github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/pkg/jwt"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/pkg/redis"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/shared"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type UserController struct {
	userService *UserService
	appConfig   *config.Configuration
	redisClient *redis.RedisClient
}

func NewUserController(service *UserService, appConfig *config.Configuration, redisClient *redis.RedisClient) *UserController {
	return &UserController{
		userService: service,
		appConfig:   appConfig,
		redisClient: redisClient,
	}
}

func (c *UserController) SignUp(g *gin.Context) {
	var requestModel types.SignupRequest

	//check request body is correct form
	if err := g.ShouldBind(&requestModel); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": shared.GeneralErrorRequestBodyNotCorrect.Error(),
		})
		return
	}

	//service call
	err2 := c.userService.SignupService(requestModel)

	if err2 != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Check your form of inputs. Error: " + err2.Error(),
		})
		return
	}

	g.JSON(http.StatusCreated, requestModel)

}

func (c *UserController) SignIn(g *gin.Context) {
	var requestModel types.SigninRequest

	//check request body is correct form
	if err := g.ShouldBindJSON(&requestModel); err != nil {
		fmt.Println(err)
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": shared.GeneralErrorRequestBodyNotCorrect.Error(),
		})
		g.Abort()
		return
	}

	//service call
	user, err2 := c.userService.SignInService(requestModel)

	//check user is exist
	if err2 != nil {
		g.JSON(http.StatusNotFound, gin.H{
			"error_message": "User not found or password is not correct!",
		})
		g.Abort()
		return

	}

	//generate jwt token
	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": strconv.Itoa(int(user.ID)),
		"email":  user.Email,
		"iat":    time.Now().Unix(),
		"iss":    os.Getenv("ENV"),
		"exp":    time.Now().Add(time.Hour).Unix(),
		"role":   user.Role,
	})
	token := jwtHelper.GenerateToken(jwtClaims, c.appConfig.JwtSettings.SecretKey)

	//save the token to redis
	err3 := c.redisClient.SetKey(user.Email, token, time.Hour)

	if err3 != nil {
		log.Fatalln("Redis is unreachable!")
		g.JSON(http.StatusBadGateway, gin.H{
			"error_message": "Inmemory cache is unreachable!",
		})
		g.Abort()
		return
	}

	//ok request
	log.Println(fmt.Sprintf("User id: %d, email: %s is login", user.ID, user.Email))
	//TODO
	log.Println(token)
	log.Println(jwtHelper.VerifyToken(token, c.appConfig.SecretKey))
	log.Println(jwtHelper.VerifyToken(token, c.appConfig.SecretKey).Role.Name)

	g.JSON(http.StatusOK, token)

}

func (c *UserController) SignOut(g *gin.Context) {
	var requestModel types.SignoutRequest

	//check request body is correct form
	if err := g.ShouldBind(&requestModel); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": shared.GeneralErrorRequestBodyNotCorrect,
		})
		return
	}

	token := g.GetHeader("Authorization")

	var decodedToken jwtHelper.DecodedToken
	decodedToken = *jwtHelper.VerifyToken(token, c.appConfig.JwtSettings.SecretKey)

	//service call
	c.userService.SignOutService(token, decodedToken.Email, *c.redisClient)
}
