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

	"BE-REPO-20/utils/encrypts"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	hashService := encrypts.NewHashService()

	authData := _dataAuth.NewAuth(db)
	authService := _serviceAuth.NewAuth(authData, hashService)
	authHandler := _handlerAuth.NewAuth(authService)


	adminData := _dataAdmin.NewAdmin(db)
	adminService := _serviceAdmin.NewAdmin(adminData)
	adminHandler := _handlerAdmin.NewAdmin(adminService)

	// login
	e.POST("/login", authHandler.Login)
	e.POST("/register", authHandler.Register)

	//admin
	e.GET("/users", adminHandler.GetAllUsers, middlewares.JWTMiddleware())
	e.GET("/users/search", adminHandler.SearchUsersByQuery, middlewares.JWTMiddleware())
	e.GET("/orders", adminHandler.GetAllOrders, middlewares.JWTMiddleware())


	userData := _dataUser.New(db)
	userService := _serviceUser.NewUser(userData)
	userHandler := _handlerUser.New(userService)

	// login
	e.POST("/login", authHandler.Login)
	e.POST("/register", authHandler.Register)\
  e.PUT("/update-password", authHandler.UpdatePassword, middlewares.JWTMiddleware())

	// user
	e.GET("/user", userHandler.SelectUser, middlewares.JWTMiddleware())
	e.DELETE("/user", userHandler.Delete, middlewares.JWTMiddleware())

}
