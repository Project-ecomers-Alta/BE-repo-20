package service

import (
	"BE-REPO-20/features/product"
	"errors"
)

type productService struct {
	prouctData product.ProductDataInterface
}

func NewProduct(repo product.ProductDataInterface) product.ProductServiceInterface {
	return &productService{
		prouctData: repo,
	}
}

// CreateProduct implements product.ProductServiceInterface.
func (service *productService) CreateProduct(userId int, input product.ProductCore) error {
	// if input.Name == "" {
	// 	return errors.New("buat nama project")
	// }
	err := service.prouctData.CreateProduct(userId, input)
	if err != nil {
		return err
	}
	return nil
}

// SelectAllProduct implements product.ProductServiceInterface.
func (service *productService) SelectAllProduct() ([]product.ProductCore, error) {
	data, err := service.prouctData.SelectAllProduct()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// SelectProductById implements product.ProductServiceInterface.
func (service *productService) SelectProductById(userId int, id int) (*product.ProductCore, error) {
	data, err := service.prouctData.SelectProductById(userId, id)
	if err != nil {
		return nil, err
	}
	// fmt.Println(data)
	return data, nil
}

// UpdateProductById implements product.ProductServiceInterface.
func (service *productService) UpdateProductById(userId int, id int, input product.ProductCore) error {
	panic("unimplemented")
}

// DeleteProductById implements product.ProductServiceInterface.
func (service *productService) DeleteProductById(userId int, id int) error {
	panic("unimplemented")
}

// SearchProductByQuery implements product.ProductServiceInterface.
func (service *productService) SearchProductByQuery(query string) ([]product.ProductCore, error) {
	data, err := service.prouctData.SearchProductByQuery(query)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// AddImageProduct implements product.ProductServiceInterface.
func (service *productService) AddImageProduct(productID int, image string) error {
	if productID <= 0 {
		return errors.New("invalid id")
	}

	err := service.prouctData.AddImageProduct(productID, image)
	return err
}
