package data

import (
	"BE-REPO-20/features/product"

	"gorm.io/gorm"
)

type productQuery struct {
	db *gorm.DB
}

func NewProduct(db *gorm.DB) product.ProductDataInterface {
	return &productQuery{
		db: db,
	}
}

// CreateProduct implements product.ProductDataInterface.
func (repo *productQuery) CreateProduct(userId int, input product.ProductCore) error {
	panic("unimplemented")
}

// SelectAllProduct implements product.ProductDataInterface.
func (repo *productQuery) SelectAllProduct() ([]product.ProductCore, error) {
	var productGorm []Product
	tx := repo.db.Preload("User").Find(&productGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	produtCore := ModelToCoreList(productGorm)

	return produtCore, nil
}

// SelectProductById implements product.ProductDataInterface.
func (repo *productQuery) SelectProductById(userId int, id int) (*product.ProductCore, error) {
	var productGorm Product
	tx := repo.db.Preload("User").First(&productGorm, id)

	if tx.Error != nil {
		return nil, tx.Error
	}

	result := productGorm.ModelToCore()
	return &result, nil
}

// SearchProductByQuery implements product.ProductDataInterface.
func (repo *productQuery) SearchProductByQuery(query string) ([]product.ProductCore, error) {
	var productGorm []Product
	tx := repo.db.Where("name LIKE ? OR category LIKE ?", "%"+query+"%", "%"+query+"%").Find(&productGorm)

	if tx.Error != nil {
		return nil, tx.Error
	}

	itemOrderCoreList := ModelToCoreList(productGorm)

	return itemOrderCoreList, nil
}

// UpdateProductById implements product.ProductDataInterface.
func (repo *productQuery) UpdateProductById(userId int, id int, input product.ProductCore) error {
	panic("unimplemented")
}

// DeleteProductById implements product.ProductDataInterface.
func (repo *productQuery) DeleteProductById(userId int, id int) error {
	panic("unimplemented")
}
