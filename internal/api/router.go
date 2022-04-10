package api

import (
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/config"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/database"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/cart"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/cart/cart_item"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/category"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/order"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/order/order_item"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/product"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/role"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/user"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/middleware"
	redisHelper "github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/pkg/redis"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

var AppConfig = &config.Configuration{}

func RegisterHandlers(r *gin.Engine) {

	os.Setenv("ENV", "local")

	cfgFile := "app." + os.Getenv("ENV") + ".yaml"
	AppConfig, err := config.GetAllConfigValues(cfgFile)
	if err != nil {
		log.Fatalf("Failed to read config file. %v", err.Error())
	}

	db := database.Connect(AppConfig.DatabaseURI)

	redisClient := redisHelper.NewRedisClient(AppConfig)

	categoryRepository := category.NewCategoryRepository(db)
	categoryService := category.NewCategoryService(*categoryRepository)
	categoryController := category.NewCategoryController(categoryService)

	productRepository := product.NewProductRepository(db)
	productService := product.NewProductService(*productRepository, *categoryService)
	productController := product.NewProductController(productService)

	cartItemRepository := cart_item.NewCartItemRepository(db)

	cartRepository := cart.NewCartRepository(db)
	cartService := cart.NewCartService(*cartRepository, *productService, *cartItemRepository)
	cartController := cart.NewCartController(cartService, AppConfig)

	orderItemRepository := order_item.NewOrderItemRepository(db)

	orderRepository := order.NewOrderRepository(db)
	orderService := order.NewOrderService(*orderRepository, cartService, *orderItemRepository, productService)
	orderController := order.NewOrderController(orderService, AppConfig)

	roleRepository := role.NewRoleRepository(db)
	roleService := role.NewRoleService(*roleRepository)

	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(*userRepository, *roleRepository, *cartService)
	userController := user.NewUserController(userService, AppConfig, redisClient)

	//TODO - Create DB Schema and Insert Sample Data
	roleService.InsertSampleData()
	userService.InsertSampleData()
	categoryService.InsertSampleData()
	productService.InsertSampleData()

	orderGroup := r.Group("/order")
	orderGroup.POST("complete", middleware.UserAuthMiddleware(AppConfig.JwtSettings.SecretKey, redisClient), orderController.CompleteOrder)
	orderGroup.POST("/cancel", middleware.UserAuthMiddleware(AppConfig.JwtSettings.SecretKey, redisClient), orderController.CancelOrder)
	orderGroup.GET("/list", middleware.UserAuthMiddleware(AppConfig.JwtSettings.SecretKey, redisClient), orderController.ListOrders)

	cartGroup := r.Group("/cart")
	cartGroup.GET("/list", middleware.UserAuthMiddleware(AppConfig.JwtSettings.SecretKey, redisClient), cartController.ShowCart)
	cartGroup.POST("/add-item", middleware.UserAuthMiddleware(AppConfig.JwtSettings.SecretKey, redisClient), cartController.AddCartItem)
	cartGroup.PUT("/update-delete-item", middleware.UserAuthMiddleware(AppConfig.JwtSettings.SecretKey, redisClient), cartController.UpdateCartItem)

	productGroup := r.Group("/product")
	productGroup.GET("/list", productController.ListProducts)
	productGroup.POST("search", productController.SearchProducts)
	productGroup.POST("/add", middleware.AdminAuthMiddleware(AppConfig.JwtSettings.SecretKey, redisClient), productController.AddProduct)
	productGroup.PUT("/update", middleware.AdminAuthMiddleware(AppConfig.JwtSettings.SecretKey, redisClient), productController.UpdateProduct)
	productGroup.DELETE("/delete", middleware.AdminAuthMiddleware(AppConfig.JwtSettings.SecretKey, redisClient), productController.DeleteProduct)

	categoryGroup := r.Group("/category")
	categoryGroup.GET("/list", categoryController.CategoryList)
	categoryGroup.POST("/add-all", middleware.AdminAuthMiddleware(AppConfig.JwtSettings.SecretKey, redisClient), categoryController.AddCategoryFromCSV)

	authGroup := r.Group("/auth")
	authGroup.POST("/login", userController.SignIn)
	authGroup.POST("/signup", userController.SignUp)
	authGroup.POST("/logout", middleware.UserAuthMiddleware(AppConfig.JwtSettings.SecretKey, redisClient), userController.SignOut)

}
