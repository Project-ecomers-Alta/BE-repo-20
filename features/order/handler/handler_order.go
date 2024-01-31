package handler

import (
	"BE-REPO-20/app/middlewares"
	"BE-REPO-20/features/order"
	"BE-REPO-20/utils/responses"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// _midtransController "BE-REPO-20/features/midtrans/controller"
	_midtransService "BE-REPO-20/features/midtrans/service"
	"BE-REPO-20/features/midtrans/web"

	"github.com/go-playground/validator"
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

	midtransReq := OrderToMidtrans(orderResponse)

	fmt.Println(midtransReq)

	// midtrans
	validate := validator.New()
	midtransService := _midtransService.NewMidtransServiceImpl(validate)
	midtransResponse := _midtransService.MidtransService.CreateEcho(midtransService, c, midtransReq)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   midtransResponse,
	}
	return c.JSON(http.StatusOK, responses.WebResponse("Success order.", webResponse))
}

func (handler *OrderHandler) GetOrders(c echo.Context) error {
	idJWT := middlewares.ExtractTokenUserId(c)
	if idJWT == 0 {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("unauthorized or jwt expired", nil))
	}

	results, err := handler.orderService.GetOrders(uint(idJWT))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("Error order. "+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("Success get order.", results))
}

func (handler *OrderHandler) GetOrders2(c echo.Context) error {
	idJWT := middlewares.ExtractTokenUserId(c)
	if idJWT == 0 {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("unauthorized or jwt expired", nil))
	}

	newOrder := OrderRequest{}
	errBind := c.Bind(&newOrder)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data order not valid", nil))
	}

	orderCore := RequestToCoreOrder(newOrder)
	payment, errInsert := handler.orderService.PostOrder(uint(idJWT), orderCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert order "+errInsert.Error(), nil))
	}

	results := CoreToResponse(*payment)

	return c.JSON(http.StatusOK, responses.WebResponse("Success get order.", results))
}

func (handler *OrderHandler) WebhoocksNotification(c echo.Context) error {
	json_map := make(map[string]interface{})
	errJson := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if errJson != nil {
		panic(errJson)
	}
	fmt.Println(json_map)

	var webhoocksReq = WebhoocksRequest{}
	errBind := c.Bind(&webhoocksReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}
	// fmt.Println(c.Request().Body)
	fmt.Println(webhoocksReq)

	orderCore := WebhoocksRequestToCore(webhoocksReq)
	err := handler.orderService.WebhoocksService(orderCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error Notif "+err.Error(), nil))
	}

	log.Println("transaction success")
	return c.JSON(http.StatusOK, responses.WebResponse("transaction success", nil))
}
