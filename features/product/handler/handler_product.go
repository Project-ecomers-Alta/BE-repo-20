package handler

import (
	"BE-REPO-20/app/middlewares"
	"BE-REPO-20/features/product"
	"BE-REPO-20/utils/responses"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

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

func (handler *ProductHandler) ListProductPenjualan(c echo.Context) error {
	idJWT := middlewares.ExtractTokenUserId(c)
	userID := uint(idJWT)
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	products, err := handler.productService.ListProductPenjualan(page, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data. "+err.Error(), nil))
	}

	productsResponse := CoreToResponseList(products)

	return c.JSON(http.StatusOK, responses.WebResponse("success read products.", productsResponse))
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
		return c.JSON(http.StatusNotFound, responses.WebResponse("error read data. "+err.Error(), nil))
	}
	productResponse := CoreToResponse(*product)

	return c.JSON(http.StatusOK, responses.WebResponse("success read product.", productResponse))
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
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data.", nil))
	}
	NewProduct.UserId = uint(idJWT)
	productCore := RequestToCore(NewProduct)
	err := handler.productService.CreateProduct(idJWT, productCore)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error create product."+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success insert product", NewProduct))
}

func (handler *ProductHandler) UpdateProduct(c echo.Context) error {
	idJWT := middlewares.ExtractTokenUserId(c)
	if idJWT == 0 {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("unauthorized or jwt expired", nil))
	}

	id := c.Param("product_id")
	idParam, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error. id should be number", nil))
	}

	var reqDataProduct = ProductRequest{}
	errBind := c.Bind(&reqDataProduct)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind, data not valid", nil))
	}
	reqDataProduct.UserId = uint(idJWT)

	productCore := RequestToCore(reqDataProduct)

	err := handler.productService.UpdateProductById(idJWT, idParam, productCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error update data. update failed "+err.Error(), nil))
	}
	return c.JSON(http.StatusOK, responses.WebResponse("success update data", reqDataProduct))
}

func (handler *ProductHandler) DeleteProduct(c echo.Context) error {
	idJWT := middlewares.ExtractTokenUserId(c)
	if idJWT == 0 {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("unauthorized or jwt expired", nil))
	}

	id := c.Param("product_id")
	idParam, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error. id should be number", nil))
	}

	err := handler.productService.DeleteProductById(idJWT, idParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error delete data. delete failed"+err.Error(), nil))
	}
	return c.JSON(http.StatusOK, responses.WebResponse("success delete data", nil))
}

func (handler *ProductHandler) CreateProductImage(c echo.Context) error {
	var fileSize int64
	var nameFile string
	idJWT := middlewares.ExtractTokenUserId(c)
	if idJWT == 0 {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("unauthorized or jwt expired", nil))
	}

	id := c.Param("product_id")
	idParam, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error. id should be number", nil))
	}

	var prodImgReq = ProductImageRequest{}
	errBind := c.Bind(&prodImgReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind, data not valid", nil))
	}
	prodImgReq.ProductID = uint(idParam)
	prodImgCore := RequestToCoreImage(prodImgReq)

	// upload img
	fileHeader, _ := c.FormFile("image_url")
	var file multipart.File
	if fileHeader != nil {
		openFileHeader, _ := fileHeader.Open()
		file = openFileHeader

		// get data type
		nameFile = fileHeader.Filename
		nameFileSplit := strings.Split(nameFile, ".")
		indexFile := len(nameFileSplit) - 1

		// data type validation
		if nameFileSplit[indexFile] != "jpeg" && nameFileSplit[indexFile] != "png" && nameFileSplit[indexFile] != "jpg" {
			return c.JSON(http.StatusBadRequest, responses.WebResponse("error invalid type format, format file not valid", nil))
		}

		// data size validation max 2mb
		fileSize = fileHeader.Size
		if fileSize >= 2000000 {
			return c.JSON(http.StatusBadRequest, responses.WebResponse("error size data, file size is too big", nil))
		}
	}

	err := handler.productService.CreateProductImage(file, prodImgCore, nameFile, idJWT)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error post image data. post image failed "+err.Error(), nil))
	}
	return c.JSON(http.StatusOK, responses.WebResponse("success post image data", nil))
}

func (handler *ProductHandler) DeleteProductImageId(c echo.Context) error {
	idJWT := middlewares.ExtractTokenUserId(c)
	if idJWT == 0 {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("unauthorized or jwt expired", nil))
	}

	id := c.Param("product_id")
	idParamProduct, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error. id should be number", nil))
	}

	idImg := c.Param("image_id")
	idParamImg, errConv := strconv.Atoi(idImg)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error. id should be number", nil))
	}

	err := handler.productService.DeleteProductImageById(idJWT, idParamProduct, idParamImg)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error delete image data. delete image failed "+err.Error(), nil))
	}
	return c.JSON(http.StatusOK, responses.WebResponse("success delete image data", nil))
}
