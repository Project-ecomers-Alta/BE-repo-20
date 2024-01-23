package handler

import (
	"BE-REPO-20/features/product"
	"BE-REPO-20/utils/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productService product.ProductServiceInterface
}

func NewProduct(service product.ProductServiceInterface) *ProductHandler {
	return &ProductHandler{
		productService: service,
	}
}

func (handler *ProductHandler) SelectAllProduct(c echo.Context) error {
	products, err := handler.productService.SelectAllProduct()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data. "+err.Error(), nil))
	}

	productsResponse := CoreToResponseList(products)

	return c.JSON(http.StatusOK, responses.WebResponse("success read products.", productsResponse))
}
