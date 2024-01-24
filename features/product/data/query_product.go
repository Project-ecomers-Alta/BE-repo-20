package data

import (
	"BE-REPO-20/features/product"
	"errors"

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

// SearchProductByQuery implements product.ProductDataInterface.
func (repo *productQuery) SearchProductByQuery(query string, offset, limit int) ([]product.ProductCore, error) {
	var productGorm []Product
	tx := repo.db.Where("name LIKE ? OR category LIKE ?", "%"+query+"%", "%"+query+"%").Offset(offset).Limit(limit).Find(&productGorm)

	if tx.Error != nil {
		return nil, tx.Error
	}

	itemOrderCoreList := ModelToCoreList(productGorm)

	return itemOrderCoreList, nil
}

// SelectAllProduct implements product.ProductDataInterface.
func (repo *productQuery) SelectAllProduct(offset, limit int) ([]product.ProductCore, error) {
	var productGorm []Product
	tx := repo.db.Offset(offset).Limit(limit).Preload("User").Find(&productGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	produtCore := ModelToCoreList(productGorm)

	return produtCore, nil
}

// CreateProduct implements product.ProductDataInterface.
func (repo *productQuery) CreateProduct(userId int, input product.ProductCore) error {
	productGorm := CoreToModel(input)
	tx := repo.db.Create(&productGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert product failed, row affected = 0")
	}
	return nil
}

// DeleteProductById implements product.ProductDataInterface.
func (repo *productQuery) DeleteProductById(userId int, id int) error {
	panic("unimplemented")
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

// UpdateProductById implements product.ProductDataInterface.
func (repo *productQuery) UpdateProductById(userId int, id int, input product.ProductCore) error {
	panic("unimplemented")
}
