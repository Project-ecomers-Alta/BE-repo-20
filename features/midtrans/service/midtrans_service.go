package service

import (
	"BE-REPO-20/features/midtrans/web"

	"github.com/labstack/echo/v4"
)

type MidtransService interface {
	CreateEcho(c echo.Context, request web.MidtransRequest) web.MidtransResponse
}
