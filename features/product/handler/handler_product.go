package handler

import (
	"BE-REPO-20/app/middlewares"
	"BE-REPO-20/features/product"
	"BE-REPO-20/utils/aws"
	"BE-REPO-20/utils/responses"
	"fmt"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productService product.ProductServiceInterface
	awsSession     *aws.Session
}

func NewProduct(service product.ProductServiceInterface, awsSession *aws.Session) *ProductHandler {
	return &ProductHandler{
		productService: service,
		awsSession:     awsSession,
	}
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
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error create product.", nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success insert product", nil))
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

func (handler *ProductHandler) AddImageProduct(c echo.Context) error {
	// Extract product ID from the request
	id := c.Param("product_id")
	idParam, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error. product_id should be a number", nil))
	}

	// Extract and validate the uploaded image from the form data
	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error uploading image. "+err.Error(), nil))
	}

	// Create a temporary file to store the uploaded image
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error opening uploaded image. "+err.Error(), nil))
	}
	defer src.Close()

	// Upload the image to AWS S3 using the AWS session and get the URL
	imageURL, err := handler.awsSession.Upload(file, src)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error uploading image to AWS S3. "+err.Error(), nil))
	}

	// Update the product with the image information and URL
	err = handler.productService.AddImageProduct(idParam, imageURL)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error updating product with image. "+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success uploading image product.", nil))
}
