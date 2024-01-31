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

func (handler *Carthandler) DeleteCarts(c echo.Context) error {
	var request struct {
		Ids []uint `json:"ids"`
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error parsing request."+err.Error(), nil))
	}

	err := handler.cartService.DeleteCarts(request.Ids)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error delete data."+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success",
	})
}

func (handler *Carthandler) CreateCart(c echo.Context) error {
	idJWT := middlewares.ExtractTokenUserId(c)

	newCart := CartRequest{}
	errBind := c.Bind(&newCart)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid."+errBind.Error(), nil))
	}

	cartCore := _cart.CartCore{
		ProductID: newCart.ProductId,
		UserID:    uint(idJWT),
		Quantity:  newCart.Quantity,
	}

	err := handler.cartService.CreateCart(cartCore)
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
