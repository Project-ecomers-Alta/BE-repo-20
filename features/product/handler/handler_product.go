package handler

import (
	"BE-REPO-20/app/middlewares"
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

func (handler *ProductHandler) SearchProductByQuery(c echo.Context) error {
	query := c.QueryParam("search")

	offset := 0
	limit := 10
	products, err := handler.productService.SearchProductByQuery(query, offset, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error reading data. "+err.Error(), nil))
	}

	productsResponse := CoreToResponseList(products)

	return c.JSON(http.StatusOK, responses.WebResponse("success reading products.", productsResponse))
}

func (handler *ProductHandler) CreateProduct(c echo.Context) error {
	idJWT := middlewares.ExtractTokenUserId(c)
	if idJWT == 0 {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("unauthorized or jwt expired", nil))
	}

	NewProduct := ProductRequest{}
	errBind := c.Bind(&NewProduct)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data."+errBind.Error(), nil))
	}

	NewProduct.UserId = uint(idJWT)
	productCore := RequestToCore(NewProduct)

	err := handler.productService.CreateProduct(idJWT, productCore)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error create product."+err.Error(), nil))
	}	

	return c.JSON(http.StatusOK, responses.WebResponse("success insert product", NewProduct))
}

func (handler *ProductHandler) SelectAllProduct(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	products, err := handler.productService.SelectAllProduct(page)
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
