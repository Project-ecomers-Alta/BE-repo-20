package handler

import (
	"BE-REPO-20/app/middlewares"
	_cart "BE-REPO-20/features/cart"
	"BE-REPO-20/utils/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Carthandler struct {
	cartService _cart.CartDataInterface
}

func NewCart(service _cart.CartServiceInterface) *Carthandler {
	return &Carthandler{
		cartService: service,
	}
}

func (handler *Carthandler) SelectAllCart(c echo.Context) error {
	idJWT := middlewares.ExtractTokenUserId(c)

	userID := uint(idJWT)

	carts, err := handler.cartService.SelectAllCart(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data. "+err.Error(), nil))
	}

	cartResponse := CoreToResponseList(carts)
	return c.JSON(http.StatusOK, responses.WebResponse("success read cart.", cartResponse))
}

func (handler *Carthandler) CreateCart(c echo.Context) error {
	idJWT := middlewares.ExtractTokenUserId(c)
	productID := getProductIDFromParams(c)
	err := handler.cartService.CreateCart(idJWT, productID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error create cart."+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success menambahkan cart",
	})
}

func getProductIDFromParams(c echo.Context) uint {
	productID, err := strconv.ParseUint(c.Param("product_id"), 10, 64)
	if err != nil {
		return 0
	}
	return uint(productID)
}
