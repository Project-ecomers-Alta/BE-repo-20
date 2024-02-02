package service

import (
	"BE-REPO-20/features/product"
	"BE-REPO-20/mocks"
	"errors"
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
	repo.On("ListProductPenjualan", 0, 10, mock.Anything).Return(expectedProducts, nil).Once()

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
	repo.On("SelectAllProduct", 0, 10).Return(expectedProducts, nil).Once()

	result, err := srv.SelectAllProduct(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedProducts, result)
	repo.AssertExpectations(t)
}

func TestCreateProduct(t *testing.T) {
	t.Run("Valid Product", func(t *testing.T) {
		repo := new(mocks.ProductData)
		srv := NewProduct(repo)

		validInput := product.ProductCore{Name: "Product 1", Price: 1000, Quantity: 10, Category: "Category 1"}

		repo.On("CreateProduct", 1, validInput).Return(nil).Once()

		err := srv.CreateProduct(1, validInput)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Empty Product Name", func(t *testing.T) {
		repo := new(mocks.ProductData)
		srv := NewProduct(repo)

		invalidInput := product.ProductCore{Name: ""}

		err := srv.CreateProduct(1, invalidInput)
		expectedErr := errors.New("field name must be filled")
		assert.EqualError(t, err, expectedErr.Error())
	})

	t.Run("Empty Product Price", func(t *testing.T) {
		repo := new(mocks.ProductData)
		srv := NewProduct(repo)

		invalidInput := product.ProductCore{Name: "Product 1", Price: 0}

		err := srv.CreateProduct(1, invalidInput)
		expectedErr := errors.New("field price must be filled")
		assert.EqualError(t, err, expectedErr.Error())
	})

	t.Run("Empty Product Quantity", func(t *testing.T) {
		repo := new(mocks.ProductData)
		srv := NewProduct(repo)

		invalidInput := product.ProductCore{Name: "Product 1", Price: 1000, Quantity: 0}

		err := srv.CreateProduct(1, invalidInput)
		expectedErr := errors.New("field quantity must be filled")
		assert.EqualError(t, err, expectedErr.Error())
	})

	t.Run("Empty Product Category", func(t *testing.T) {
		repo := new(mocks.ProductData)
		srv := NewProduct(repo)

		invalidInput := product.ProductCore{Name: "Product 1", Price: 1000, Quantity: 10, Category: ""}

		err := srv.CreateProduct(1, invalidInput)
		expectedErr := errors.New("field category must be filled")
		assert.EqualError(t, err, expectedErr.Error())
	})
}

func TestSelectProductById(t *testing.T) {
	repo := new(mocks.ProductData)
	srv := NewProduct(repo)
	expectedProduct := &product.ProductCore{ID: 1, Name: "Product 1", Price: 1000, Quantity: 10, Category: "Category 1"}
	repo.On("SelectProductById", 1, mock.Anything).Return(expectedProduct, nil)

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
	repo.On("SearchProductByQuery", "query", 0, 10).Return(expectedProducts, nil).Once()

	result, err := srv.SearchProductByQuery("query", 0, 10)

	assert.NoError(t, err)
	assert.Equal(t, expectedProducts, result)
	repo.AssertExpectations(t)
}

func TestUpdateProductById(t *testing.T) {
	repo := new(mocks.ProductData)
	srv := NewProduct(repo)
	input := product.ProductCore{Name: "Product 1", Price: 1000, Quantity: 10, Category: "Category 1"}
	repo.On("UpdateProductById", 1, 1, mock.Anything).Return(nil).Once()

	err := srv.UpdateProductById(1, 1, input)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestDeleteProductById(t *testing.T) {
	repo := new(mocks.ProductData)
	srv := NewProduct(repo)
	repo.On("DeleteProductById", 1, 1).Return(nil).Once()

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
	repo.On("CreateProductImage", *mockFile, input, "filename", 1).Return(nil).Once()

	err := srv.CreateProductImage(*mockFile, input, "filename", 1)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestDeleteProductImageById(t *testing.T) {
	t.Run("Valid Image ID", func(t *testing.T) {
		repo := new(mocks.ProductData)
		srv := NewProduct(repo)

		validIdImage := 1
		repo.On("DeleteProductImageById", 1, 1, validIdImage).Return(nil).Once()
		err := srv.DeleteProductImageById(1, 1, validIdImage)

		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Invalid Image ID", func(t *testing.T) {
		repo := new(mocks.ProductData)
		srv := NewProduct(repo)

		invalidIdImage := 0
		err := srv.DeleteProductImageById(1, 1, invalidIdImage)
		expectedErr := errors.New("invalid id")
		assert.EqualError(t, err, expectedErr.Error())
	})
}
