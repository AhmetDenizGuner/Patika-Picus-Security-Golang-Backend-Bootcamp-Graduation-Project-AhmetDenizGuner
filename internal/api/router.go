package api

import (
	"fmt"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/config"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/database"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/user"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

var AppConfig = &config.Configuration{}

func RegisterHandlers(r *gin.Engine) {
	cfgFile := "./config/app" + os.Getenv("ENV") + ".yaml"
	AppConfig, err := config.GetAllConfigValues(cfgFile)
	if err != nil {
		log.Fatalf("Failed to read config file. %v", err.Error())
	}

	db := database.Connect(AppConfig.DatabaseURI)
	fmt.Println(db)

	productGroup := r.Group("/product")
	productGroup.GET("/list")

	categoryGroup := r.Group("/category")
	categoryGroup.GET("/list")

	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(*userRepository)
	userController := user.NewUserController(userService)

	authGroup := r.Group("/auth")
	authGroup.POST("/login")
	authGroup.POST("/signup", userController.SignUp)
	authGroup.POST("/logout")

}
