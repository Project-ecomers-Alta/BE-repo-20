package data

import (
	"BE-REPO-20/features/product"
	_userData "BE-REPO-20/features/user/data"
	"BE-REPO-20/utils/uploads"
	"errors"
	"mime/multipart"

	"gorm.io/gorm"
)

type productQuery struct {
	db            *gorm.DB
	uploadService uploads.CloudinaryInterface
}

func NewProduct(db *gorm.DB, cloud uploads.CloudinaryInterface) product.ProductDataInterface {
	return &productQuery{
		db:            db,
		uploadService: cloud,
	}
}

// ListProductPenjualan implements product.ProductDataInterface.
func (repo *productQuery) ListProductPenjualan(offset int, limit int, userId uint) ([]product.ProductCore, error) {
	var productData []Product
	tx := repo.db.Where("user_id = ?", userId).Offset(offset).Limit(limit).Preload("User").Preload("ProductImages").Find(&productData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	produtCore := ModelToCoreList(productData)

	return produtCore, nil
}

// SelectAllProduct implements product.ProductDataInterface.
func (repo *productQuery) SelectAllProduct(offset, limit int) ([]product.ProductCore, error) {
	var productGorm []Product
	tx := repo.db.Offset(offset).Limit(limit).Preload("User").Preload("ProductImages").Find(&productGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}
	produtCore := ModelToCoreList(productGorm)

	return produtCore, nil
}

func (repo *productQuery) GetTotalImagesOfProduct(productId int) (uint, error) {
	var productImg []ProductImage
	var result uint
	if err := repo.db.Table("product_images").Where("product_id = ?", productId).Find(&productImg).Error; err != nil {
		return 0, err
	}
	result = uint(len(productImg))
	if result < 1 {
		return 0, nil
	}

	return result, nil
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

// SelectProductById implements product.ProductDataInterface.
func (repo *productQuery) SelectProductById(userId int, id int) (*product.ProductCore, error) {
	var productGorm Product
	tx := repo.db.Preload("User").Preload("ProductImages").First(&productGorm, id)

	if tx.Error != nil {
		return nil, tx.Error
	}

	result := productGorm.ModelToCore()
	return &result, nil
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

// UpdateProductById implements product.ProductDataInterface.
func (repo *productQuery) UpdateProductById(userId int, id int, input product.ProductCore) error {
	dataProd, _ := repo.SelectProductById(userId, id)

	if dataProd.UserID != uint(userId) {
		return errors.New("user unauthorized id mismatched")
	}

	productGorm := CoreToModel(input)
	tx := repo.db.Model(&Product{}).Where("id = ? AND user_id = ?", id, userId).Updates(productGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("error not found row aff")
	}
	return nil
}

// DeleteProductById implements product.ProductDataInterface.
func (repo *productQuery) DeleteProductById(userId int, id int) error {
	dataProd, _ := repo.SelectProductById(userId, id)

	if dataProd.UserID != uint(userId) {
		return errors.New("user unauthorized id mismatched")
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

func (repo *productQuery) CreateProductImage(file multipart.File, input product.ProductImageCore, nameFile string, id int) error {
	dataProd, _ := repo.SelectProductById(id, int(input.ProductID))
	if dataProd.UserID != uint(id) {
		return errors.New("user unauthorized id mismatched")
	}

	numberImg, err := repo.GetTotalImagesOfProduct(int(input.ProductID))
	if err != nil {
		return errors.New("user unauthorized id mismatched")
	}

	if numberImg > 2 {
		return errors.New("the number of images meets the limit")
	}

	// foldering name at cloudinary
	var folderName string = "img/items"

	// upload file to cloudinary and input to db
	if file != nil {
		imgUrl, errUpload := repo.uploadService.Upload(file, nameFile, folderName)
		if errUpload != nil {
			return errors.New("error upload img")
		}
		input.PublicID = imgUrl.PublicID
		input.Url = imgUrl.SecureURL
	}

	imageProd := CoreToModelImage(input)

	tx := repo.db.Create(&imageProd)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert product image failed, row affected = 0")
	}
	return nil
}

func (repo *productQuery) DeleteProductImageById(userId, productId, idImage int) error {
	dataProd, _ := repo.SelectProductById(userId, productId)
	if dataProd.UserID != uint(userId) {
		return errors.New("user unauthorized id mismatched")
	}

	if idImage < 1 {
		return errors.New("error image product not found")
	}

	// get public id
	var prodImage ProductImage
	txProd := repo.db.Table("product_images").Where("id = ?", idImage).First(&prodImage)
	if txProd.Error != nil {
		return txProd.Error
	}

	if txProd.RowsAffected == 0 {
		return errors.New("insert product image failed, row affected = 0")
	}
	publicId := prodImage.PublicID
	repo.uploadService.Destroy(publicId)

	tx := repo.db.Where("id = ?", idImage).Delete(&ProductImage{})
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("error not found")
	}
	return nil
}
