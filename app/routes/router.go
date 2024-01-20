package routes

import (
	"BE-REPO-20/app/middlewares"
	_dataAuth "BE-REPO-20/features/auth/data"
	_handlerAuth "BE-REPO-20/features/auth/handler"
	_serviceAuth "BE-REPO-20/features/auth/service"

	_dataUser "BE-REPO-20/features/user/data"
	_handlerUser "BE-REPO-20/features/user/handler"
	_serviceUser "BE-REPO-20/features/user/service"
	"BE-REPO-20/utils/encrypts"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRuter(db *gorm.DB, e *echo.Echo) {
	hashService := encrypts.NewHashService()

	authData := _dataAuth.NewAuth(db)
	authService := _serviceAuth.NewAuth(authData, hashService)
	authHandler := _handlerAuth.NewAuth(authService)

	userData := _dataUser.New(db)
	userService := _serviceUser.NewUser(userData)
	userHandler := _handlerUser.New(userService)

	// login
	e.POST("/login", authHandler.Login)
	e.POST("/register", authHandler.Register)

	// user
	e.DELETE("/user", userHandler.Delete, middlewares.JWTMiddleware())
}
