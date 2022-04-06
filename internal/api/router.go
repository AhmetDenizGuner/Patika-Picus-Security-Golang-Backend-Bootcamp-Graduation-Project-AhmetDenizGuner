package api

import (
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/config"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/database"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/cart"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/category"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/product"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/user"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/middleware"
	redisHelper "github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/pkg/redis"
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

	redisClient := redisHelper.NewRedisClient(AppConfig)

	db := database.Connect(AppConfig.DatabaseURI)

	categoryRepository := category.NewCategoryRepository(db)
	categoryService := category.NewCategoryService(*categoryRepository)
	categoryController := category.NewCategoryController(categoryService)

	productRepository := product.NewProductRepository(db)
	productService := product.NewProductService(*productRepository, *categoryService)
	productController := product.NewProductController(productService)

	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(*userRepository)
	userController := user.NewUserController(userService, AppConfig, redisClient)

	cartRepository := cart.NewCartRepository(db)
	cartService := cart.NewCartService(*cartRepository)
	cartController := cart.NewCartController(cartService, AppConfig)

	cartGroup := r.Group("/cart")
	cartGroup.GET("/list", middleware.UserAuthMiddleware(AppConfig.JwtSettings.SecretKey), cartController.AddCartItem)
	cartGroup.POST("/add-item", middleware.UserAuthMiddleware(AppConfig.JwtSettings.SecretKey), cartController.UpdateCartItem)
	cartGroup.PUT("/update-delete-item", middleware.UserAuthMiddleware(AppConfig.JwtSettings.SecretKey), cartController.ShowCart)

	productGroup := r.Group("/product")
	productGroup.GET("/list", productController.ListProducts)
	productGroup.POST("search", productController.SearchProducts)
	productGroup.POST("/add", middleware.AdminAuthMiddleware(AppConfig.JwtSettings.SecretKey), productController.AddProduct)
	productGroup.PUT("/update", middleware.AdminAuthMiddleware(AppConfig.JwtSettings.SecretKey), productController.UpdateProduct)
	productGroup.DELETE("/delete", middleware.AdminAuthMiddleware(AppConfig.JwtSettings.SecretKey), productController.DeleteProduct)

	categoryGroup := r.Group("/category")
	categoryGroup.GET("/list", categoryController.CategoryList)
	categoryGroup.GET("/add-all", middleware.AdminAuthMiddleware(AppConfig.JwtSettings.SecretKey), categoryController.AddCategoryFromCSV)

	authGroup := r.Group("/auth")
	authGroup.POST("/login", userController.SignIn)
	authGroup.POST("/signup", userController.SignUp)
	authGroup.POST("/logout", userController.SignOut)

}
