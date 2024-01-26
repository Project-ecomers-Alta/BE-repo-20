package handler

import (
	"BE-REPO-20/app/middlewares"
	"BE-REPO-20/features/order"
	"BE-REPO-20/utils/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	orderService order.OrderServiceInterface
}

func NewOrder(service order.OrderServiceInterface) *OrderHandler {
	return &OrderHandler{
		orderService: service,
	}
}

func (handler *OrderHandler) CreateOrder(c echo.Context) error {
	idJWT := middlewares.ExtractTokenUserId(c)
	if idJWT == 0 {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("unauthorized or jwt expired", nil))
	}

	newOrder := OrderRequest{}
	errBind := c.Bind(&newOrder)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data.", nil))
	}
	newOrder.UserId = uint(idJWT)
	// newOrder.UserId = 4
	orderCore := OrderRequestToCore(newOrder)

	results, err := handler.orderService.PostOrder(uint(idJWT), orderCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("Error order. "+err.Error(), nil))
	}
	orderResponse := CoreToResponse(*results)
	return c.JSON(http.StatusOK, responses.WebResponse("Success order.", orderResponse))
}
