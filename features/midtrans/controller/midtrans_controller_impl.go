package controller

import (
	"BE-REPO-20/features/midtrans/service"
	"BE-REPO-20/features/midtrans/web"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MidtransControllerHandler struct {
	MidtransService service.MidtransService
}

func NewMidtransControllerHandler(midtransService service.MidtransService) *MidtransControllerHandler {
	return &MidtransControllerHandler{
		MidtransService: midtransService,
	}
}

func (handler *MidtransControllerHandler) CreateEcho(c echo.Context) error {
	var request web.MidtransRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   web.ErrorResponse{Message: err.Error()},
		})
	}

	midtransResponse := handler.MidtransService.CreateEcho(c, request)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   midtransResponse,
	}

	return c.JSON(http.StatusOK, webResponse)
}
