package service

import (
	"BE-REPO-20/features/product"
	"errors"
	"mime/multipart"
)

type productService struct {
	prouctData product.ProductDataInterface
}

func NewProduct(repo product.ProductDataInterface) product.ProductServiceInterface {
	return &productService{
		prouctData: repo,
	}
}

// ListProductPenjualan implements product.ProductServiceInterface.
func (service *productService) ListProductPenjualan(page int, userId uint) ([]product.ProductCore, error) {
	offset := (page - 1) * 10
	result, err := service.prouctData.ListProductPenjualan(offset, 10, userId)
	return result, err
}

// SelectAllProduct implements product.ProductServiceInterface.
func (service *productService) SelectAllProduct(page int) ([]product.ProductCore, error) {
	offset := (page - 1) * 10
	data, err := service.prouctData.SelectAllProduct(offset, 10)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// CreateProduct implements product.ProductServiceInterface.
func (service *productService) CreateProduct(userId int, input product.ProductCore) error {
	if input.Name == "" {
		return errors.New("field name must be filled")
	}
	if input.Price == 0 {
		return errors.New("field price must be filled")
	}
	if input.Quantity == 0 {
		return errors.New("field quantity must be filled")
	}
	if input.Category == "" {
		return errors.New("field category must be filled")
	}

	err := service.prouctData.CreateProduct(userId, input)
	if err != nil {
		return err
	}
	return nil
}

// SelectProductById implements product.ProductServiceInterface.
func (service *productService) SelectProductById(userId int, id int) (*product.ProductCore, error) {
	data, err := service.prouctData.SelectProductById(userId, id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// SearchProductByQuery implements product.ProductServiceInterface.
func (service *productService) SearchProductByQuery(query string, offest, limit int) ([]product.ProductCore, error) {
	data, err := service.prouctData.SearchProductByQuery(query, offest, limit)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// UpdateProductById implements product.ProductServiceInterface.
func (service *productService) UpdateProductById(userId int, id int, input product.ProductCore) error {
	if id <= 0 {
		return errors.New("invalid id")
	}

	err := service.prouctData.UpdateProductById(userId, id, input)
	return err
}

// DeleteProductById implements product.ProductServiceInterface.
func (service *productService) DeleteProductById(userId int, id int) error {
	if id <= 0 {
		return errors.New("invalid id")
	}
	err := service.prouctData.DeleteProductById(userId, id)
	return err
}

func (service *productService) CreateProductImage(file multipart.File, input product.ProductImageCore, nameFile string, id int) error {
	if id <= 0 {
		return errors.New("invalid id")
	}

	err := service.prouctData.CreateProductImage(file, input, nameFile, id)
	return err
}

func (service *productService) DeleteProductImageById(userId, produductId, idImage int) error {
	if idImage <= 0 {
		return errors.New("invalid id")
	}

	err := service.prouctData.DeleteProductImageById(userId, produductId, idImage)
	return err
}
