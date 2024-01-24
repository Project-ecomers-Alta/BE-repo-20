package routes

import (
	"BE-REPO-20/app/middlewares"
	_dataAuth "BE-REPO-20/features/auth/data"
	_handlerAuth "BE-REPO-20/features/auth/handler"
	_serviceAuth "BE-REPO-20/features/auth/service"
	"log"

	_dataAdmin "BE-REPO-20/features/admin/data"
	_handlerAdmin "BE-REPO-20/features/admin/handler"
	_serviceAdmin "BE-REPO-20/features/admin/service"

	_dataUser "BE-REPO-20/features/user/data"
	_handlerUser "BE-REPO-20/features/user/handler"
	_serviceUser "BE-REPO-20/features/user/service"

	_dataProduct "BE-REPO-20/features/product/data"
	_handlerProduct "BE-REPO-20/features/product/handler"
	_serviceProduct "BE-REPO-20/features/product/service"

	"BE-REPO-20/utils/aws"
	"BE-REPO-20/utils/encrypts"
	"BE-REPO-20/utils/uploads"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	hashService := encrypts.NewHashService()
	uploadService := uploads.NewCloudService()
	awsService, err := aws.NewSession()
	if err != nil {
		log.Fatal("Failed to create AWS session:", err)
	}

	authData := _dataAuth.NewAuth(db)
	authService := _serviceAuth.NewAuth(authData, hashService)
	authHandler := _handlerAuth.NewAuth(authService)

	adminData := _dataAdmin.NewAdmin(db)
	adminService := _serviceAdmin.NewAdmin(adminData)
	adminHandler := _handlerAdmin.NewAdmin(adminService)

	userData := _dataUser.NewUser(db, uploadService)
	userService := _serviceUser.NewUser(userData)
	userHandler := _handlerUser.NewUser(userService)

	productData := _dataProduct.NewProduct(db)
	productService := _serviceProduct.NewProduct(productData)
	productHandler := _handlerProduct.NewProduct(productService, awsService)

	// login
	e.POST("/login", authHandler.Login)
	e.POST("/register", authHandler.Register)
	e.PUT("/update-password", authHandler.UpdatePassword, middlewares.JWTMiddleware())

	//admin
	e.GET("/users", adminHandler.GetAllUsers, middlewares.JWTMiddleware())
	e.GET("/users/search", adminHandler.SearchUsersByQuery, middlewares.JWTMiddleware())
	e.GET("/orders", adminHandler.GetAllOrders, middlewares.JWTMiddleware())
	e.GET("/orders/search", adminHandler.SearchOrderByQuery, middlewares.JWTMiddleware())

	// user
	e.GET("/user", userHandler.SelectUser, middlewares.JWTMiddleware())
	e.DELETE("/user", userHandler.Delete, middlewares.JWTMiddleware())
	e.PUT("/user", userHandler.UpdateUser, middlewares.JWTMiddleware())
	e.GET("/user/shop", userHandler.SelectShop, middlewares.JWTMiddleware())
	e.PUT("/user/shop", userHandler.UpdateShop, middlewares.JWTMiddleware())

	// product
	e.GET("/product", productHandler.SelectAllProduct)
	e.GET("/product/:product_id", productHandler.SelectProductById)
	e.POST("/product", productHandler.CreateProduct, middlewares.JWTMiddleware())
	e.GET("/products/search", productHandler.SearchProductByQuery)
	e.POST("/product/:product_id/image", productHandler.AddImageProduct, middlewares.JWTMiddleware())

}
