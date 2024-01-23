package handler

import (
	"BE-REPO-20/features/product"
	"BE-REPO-20/utils/responses"
	"fmt"
	"net/http"
	"strconv"

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

func (handler *ProductHandler) SelectProductById(c echo.Context) error {
	id := c.Param("product_id")
	idParam, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error. id should be number", nil))
	}
	product, err := handler.productService.SelectProductById(1, idParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data. "+err.Error(), nil))
	}
	productResponse := CoreToResponse(*product)
	fmt.Println(productResponse)

	return c.JSON(http.StatusOK, responses.WebResponse("success read product.", productResponse))
}

func (handler *ProductHandler) SearchProductByQuery(c echo.Context) error {
	query := c.QueryParam("q")

	products, err := handler.productService.SearchProductByQuery(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error reading data. "+err.Error(), nil))
	}

	productsResponse := CoreToResponseList(products)

	return c.JSON(http.StatusOK, responses.WebResponse("success reading products.", productsResponse))
}
