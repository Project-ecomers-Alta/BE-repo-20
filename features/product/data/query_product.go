package data

import (
	"BE-REPO-20/features/product"
	_userData "BE-REPO-20/features/user/data"
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

func (repo *productQuery) GetUserId(userId int) (uint, error) {
	var user _userData.User
	if err := repo.db.Table("users").Where("id = ?", userId).First(&user).Error; err != nil {
		return 0, err
	}

	return user.ID, nil
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
	idGorm, errGorm := repo.GetUserId(userId)
	if errGorm != nil {
		return errGorm
	}
	if userId != int(idGorm) {
		return errors.New("id unauthorized")
	}
	productGorm := CoreToModel(input)
	tx := repo.db.Model(&Product{}).Where("id = ? AND user_id = ?", id, userId).Updates(productGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("error not found")
	}
	return nil
}

// DeleteProductById implements product.ProductDataInterface.
func (repo *productQuery) DeleteProductById(userId int, id int) error {
	idGorm, errGorm := repo.GetUserId(userId)
	if errGorm != nil {
		return errGorm
	}
	if userId != int(idGorm) {
		return errors.New("id unauthorized")
	}
	tx := repo.db.Where("id = ? AND user_id = ?", id, userId).Delete(&Product{})
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("error not found")
	}
	return nil
}
