package routes

import (
	"BE-REPO-20/app/middlewares"
	_dataAuth "BE-REPO-20/features/auth/data"
	_handlerAuth "BE-REPO-20/features/auth/handler"
	_serviceAuth "BE-REPO-20/features/auth/service"

	_dataAdmin "BE-REPO-20/features/admin/data"
	_handlerAdmin "BE-REPO-20/features/admin/handler"
	_serviceAdmin "BE-REPO-20/features/admin/service"

	_dataUser "BE-REPO-20/features/user/data"
	_handlerUser "BE-REPO-20/features/user/handler"
	_serviceUser "BE-REPO-20/features/user/service"

	_dataProduct "BE-REPO-20/features/product/data"
	_handlerProduct "BE-REPO-20/features/product/handler"
	_serviceProduct "BE-REPO-20/features/product/service"

	_dataCart "BE-REPO-20/features/cart/data"
	_handlerCart "BE-REPO-20/features/cart/handler"
	_serviceCart "BE-REPO-20/features/cart/service"

	_dataOrder "BE-REPO-20/features/order/data"
	_handlerOrder "BE-REPO-20/features/order/handler"
	_serviceOrder "BE-REPO-20/features/order/service"

	"BE-REPO-20/utils/encrypts"
	"BE-REPO-20/utils/midtrans"
	"BE-REPO-20/utils/uploads"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	hashService := encrypts.NewHashService()
	uploadService := uploads.NewCloudService()
	midtrans := midtrans.New()

	authData := _dataAuth.NewAuth(db)
	authService := _serviceAuth.NewAuth(authData, hashService)
	authHandler := _handlerAuth.NewAuth(authService)

	adminData := _dataAdmin.NewAdmin(db)
	adminService := _serviceAdmin.NewAdmin(adminData)
	adminHandler := _handlerAdmin.NewAdmin(adminService)

	userData := _dataUser.NewUser(db, uploadService)
	userService := _serviceUser.NewUser(userData)
	userHandler := _handlerUser.NewUser(userService)

	productData := _dataProduct.NewProduct(db, uploadService)
	productService := _serviceProduct.NewProduct(productData)
	productHandler := _handlerProduct.NewProduct(productService)

	cartData := _dataCart.NewCart(db)
	cartService := _serviceCart.NewCart(cartData)
	carthandler := _handlerCart.NewCart(cartService)

	orderData := _dataOrder.NewOrder(db, midtrans)
	orderService := _serviceOrder.NewOrder(orderData)
	orderHandler := _handlerOrder.NewOrder(orderService)

	// login
	e.POST("/login", authHandler.Login)
	e.POST("/register", authHandler.Register)
	e.PUT("/update-password", authHandler.UpdatePassword, middlewares.JWTMiddleware())

	//admin
	e.GET("/users", adminHandler.GetAllUsers, middlewares.JWTMiddleware())
	e.GET("/users", adminHandler.SearchUsersByQuery, middlewares.JWTMiddleware())
	e.GET("/orders", adminHandler.GetAllOrders, middlewares.JWTMiddleware())
	e.GET("/orders", adminHandler.SearchOrderByQuery, middlewares.JWTMiddleware())

	// user
	e.GET("/user", userHandler.SelectUser, middlewares.JWTMiddleware())
	e.DELETE("/user", userHandler.Delete, middlewares.JWTMiddleware())
	e.PUT("/user", userHandler.UpdateUser, middlewares.JWTMiddleware())
	e.GET("/user/shop", userHandler.SelectShop, middlewares.JWTMiddleware())
	e.PUT("/user/shop", userHandler.UpdateShop, middlewares.JWTMiddleware())

	// product
	e.GET("/product", productHandler.SelectAllProduct)
	e.GET("/product/:product_id", productHandler.SelectProductById)
	e.GET("/products", productHandler.SearchProductByQuery)
	e.POST("/product", productHandler.CreateProduct, middlewares.JWTMiddleware())
	e.PUT("/product/:product_id", productHandler.UpdateProduct, middlewares.JWTMiddleware())
	e.DELETE("/product/:product_id", productHandler.DeleteProduct, middlewares.JWTMiddleware())
	e.POST("/product/:product_id/image", productHandler.CreateProductImage, middlewares.JWTMiddleware())
	e.DELETE("/product/:product_id/image/:image_id", productHandler.DeleteProductImageId, middlewares.JWTMiddleware())
	e.GET("product-penjualan", productHandler.ListProductPenjualan, middlewares.JWTMiddleware())

	// cart
	e.POST("/cart", carthandler.CreateCart, middlewares.JWTMiddleware())
	e.GET("/cart", carthandler.SelectAllCart, middlewares.JWTMiddleware())
	e.DELETE("/cart", carthandler.DeleteCarts)

	// order
	e.POST("/order", orderHandler.CreateOrder, middlewares.JWTMiddleware())
	e.GET("/order", orderHandler.GetOrders, middlewares.JWTMiddleware())
	e.PUT("/orders/:order_id", orderHandler.CancelOrderById, middlewares.JWTMiddleware())
	e.POST("/payment/notification", orderHandler.WebhoocksNotification)
}
