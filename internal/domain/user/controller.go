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

//SignUp crates new user with unique email
func (c *UserController) SignUp(g *gin.Context) {
	var requestModel types.SignupRequest

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

	//service call
	err2 := c.userService.SignupService(requestModel)

	if err2 != nil {
		log.Println(err2.Error())
		g.JSON(http.StatusBadRequest, shared.ApiErrorResponse{
			IsSuccess:    true,
			ErrorMessage: ErrUserCheckFormInputs.Error() + " " + err2.Error(),
		})
		g.Abort()
		return
	}

	log.Println(fmt.Sprintf("%s registered"), requestModel.Email)
	g.JSON(http.StatusCreated, shared.ApiOkResponse{
		IsSuccess: true,
		Message:   "Register completed, please login",
		Data:      requestModel.Email},
	)
}

//SignIn generate token for user
func (c *UserController) SignIn(g *gin.Context) {
	var requestModel types.SigninRequest

	//check request body is correct form
	if err := g.ShouldBindJSON(&requestModel); err != nil {
		fmt.Println(err)
		g.JSON(http.StatusBadRequest, shared.ApiErrorResponse{
			IsSuccess:    false,
			ErrorMessage: shared.GeneralErrorRequestBodyNotCorrect.Error(),
		})
		g.Abort()
		return
	}

	//service call
	user, err2 := c.userService.SignInService(requestModel)

	//check user is exist
	if err2 != nil {
		g.JSON(http.StatusNotFound, shared.ApiErrorResponse{
			IsSuccess:    false,
			ErrorMessage: ErrUserCredentialsNotCorrect.Error(),
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
		g.JSON(http.StatusInsufficientStorage, shared.ApiErrorResponse{
			IsSuccess:    false,
			ErrorMessage: shared.GeneralServerError.Error(),
		})
		g.Abort()
		log.Fatalln("Redis is unreachable!: " + err3.Error())
		return
	}

	//ok request
	log.Println(fmt.Sprintf("User id: %d, email: %s is login", user.ID, user.Email))

	ok := shared.ApiOkResponse{
		IsSuccess: true,
		Message:   "You are login.",
		Data:      token}
	g.JSON(http.StatusOK, ok)

}

//SignOut delete the logged-in user token in redis
func (c *UserController) SignOut(g *gin.Context) {
	var requestModel types.SignoutRequest

	//check request body is correct form
	if err := g.ShouldBind(&requestModel); err != nil {
		g.JSON(http.StatusBadRequest, shared.ApiErrorResponse{
			IsSuccess:    false,
			ErrorMessage: shared.GeneralErrorRequestBodyNotCorrect.Error(),
		})
		g.Abort()
		return
	}

	token := g.GetHeader("Authorization")

	var decodedToken jwtHelper.DecodedToken
	decodedToken = *jwtHelper.VerifyToken(token, c.appConfig.JwtSettings.SecretKey)

	//service call
	err1 := c.userService.SignOutService(token, decodedToken.Email, *c.redisClient)

	if err1 != nil {
		log.Println(err1.Error())
		g.JSON(http.StatusBadRequest, shared.ApiErrorResponse{
			IsSuccess:    false,
			ErrorMessage: err1.Error(),
		})
		g.Abort()
		return
	}

	ok := shared.ApiOkResponse{
		IsSuccess: true,
		Message:   "You are logout."}
	g.JSON(http.StatusOK, ok)
}
