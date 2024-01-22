package service

import (
	"BE-REPO-20/features/product"
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
	return err
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
