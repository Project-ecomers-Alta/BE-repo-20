package service

import (
	"BE-REPO-20/features/product"
	"BE-REPO-20/mocks"
	"errors"
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListProductPenjualan(t *testing.T) {
	repo := new(mocks.ProductData)
	srv := NewProduct(repo)
	expectedProducts := []product.ProductCore{
		{
			ID:       1,
			Name:     "Product 1",
			Price:    1000,
			Quantity: 10,
			Category: "Category 1"},
	}
	repo.On("ListProductPenjualan", 0, 10, uint(1)).Return(expectedProducts, nil)

	result, err := srv.ListProductPenjualan(1, 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedProducts, result)
	repo.AssertExpectations(t)
}

func TestSelectAllProduct(t *testing.T) {
	repo := new(mocks.ProductData)
	srv := NewProduct(repo)
	expectedProducts := []product.ProductCore{
		{
			ID:       1,
			Name:     "Product 1",
			Price:    1000,
			Quantity: 10,
			Category: "Category 1"},
	}
	repo.On("SelectAllProduct", 0, 10).Return(expectedProducts, nil)

	result, err := srv.SelectAllProduct(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedProducts, result)
	repo.AssertExpectations(t)
}

func TestCreateProduct(t *testing.T) {
	repo := new(mocks.ProductData)
	srv := NewProduct(repo)

	validInput := product.ProductCore{Name: "Product 1", Price: 1000, Quantity: 10, Category: "Category 1"}
	invalidInput := product.ProductCore{}

	repo.On("CreateProduct", 1, validInput).Return(nil)

	err := srv.CreateProduct(1, validInput)
	assert.NoError(t, err)
	repo.AssertExpectations(t)

	invalidInput.Name = "" // Missing name
	err = srv.CreateProduct(1, invalidInput)
	expectedErr := errors.New("field name must be filled")
	assert.EqualError(t, err, expectedErr.Error())

	invalidInput.Name = "Product 1" // Reset name, now missing price
	invalidInput.Price = 0
	err = srv.CreateProduct(1, invalidInput)
	expectedErr = errors.New("field price must be filled")
	assert.EqualError(t, err, expectedErr.Error())

	invalidInput.Price = 1000 // Reset price, now missing quantity
	invalidInput.Quantity = 0
	err = srv.CreateProduct(1, invalidInput)
	expectedErr = errors.New("field quantity must be filled")
	assert.EqualError(t, err, expectedErr.Error())

	invalidInput.Quantity = 10 // Reset quantity, now missing category
	invalidInput.Category = ""
	err = srv.CreateProduct(1, invalidInput)
	expectedErr = errors.New("field category must be filled")
	assert.EqualError(t, err, expectedErr.Error())
}

func TestSelectProductById(t *testing.T) {
	repo := new(mocks.ProductData)
	srv := NewProduct(repo)
	expectedProduct := &product.ProductCore{ID: 1, Name: "Product 1", Price: 1000, Quantity: 10, Category: "Category 1"}
	repo.On("SelectProductById", 1, 1).Return(expectedProduct, nil)

	result, err := srv.SelectProductById(1, 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedProduct, result)
	repo.AssertExpectations(t)
}

func TestSearchProductByQuery(t *testing.T) {
	repo := new(mocks.ProductData)
	srv := NewProduct(repo)
	expectedProducts := []product.ProductCore{
		{ID: 1, Name: "Product 1", Price: 1000, Quantity: 10, Category: "Category 1"},
	}
	repo.On("SearchProductByQuery", "query", 0, 10).Return(expectedProducts, nil)

	result, err := srv.SearchProductByQuery("query", 0, 10)

	assert.NoError(t, err)
	assert.Equal(t, expectedProducts, result)
	repo.AssertExpectations(t)
}

func TestUpdateProductById(t *testing.T) {
	repo := new(mocks.ProductData)
	srv := NewProduct(repo)
	input := product.ProductCore{Name: "Product 1", Price: 1000, Quantity: 10, Category: "Category 1"}
	repo.On("UpdateProductById", 1, 1, input).Return(nil)

	err := srv.UpdateProductById(1, 1, input)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestDeleteProductById(t *testing.T) {
	repo := new(mocks.ProductData)
	srv := NewProduct(repo)
	repo.On("DeleteProductById", 1, 1).Return(nil)

	err := srv.DeleteProductById(1, 1)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestCreateProductImage(t *testing.T) {
	repo := new(mocks.ProductData)
	srv := NewProduct(repo)
	mockFile := new(multipart.File)
	input := product.ProductImageCore{
		ID:        1,
		ProductID: 1,
		Url:       "wwww.cloudinary.com",
		PublicID:  "adqwfdqfavewa",
	}
	repo.On("CreateProductImage", *mockFile, input, "filename", 1).Return(nil)

	err := srv.CreateProductImage(*mockFile, input, "filename", 1)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestDeleteProductImageById(t *testing.T) {
	repo := new(mocks.ProductData)
	srv := NewProduct(repo)

	validIdImage := 1
	repo.On("DeleteProductImageById", 1, 1, validIdImage).Return(nil)
	err := srv.DeleteProductImageById(1, 1, validIdImage)

	assert.NoError(t, err)
	repo.AssertExpectations(t)

	// Call the function under test with invalid input for idImage <= 0
	invalidIdImage := 0
	err = srv.DeleteProductImageById(1, 1, invalidIdImage)
	expectedErr := errors.New("invalid id")
	assert.EqualError(t, err, expectedErr.Error())
}
