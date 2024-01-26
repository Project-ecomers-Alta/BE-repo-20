package controller

import "github.com/labstack/echo/v4"

type MidtransController interface {
	CreateEcho(e echo.Context)
}
