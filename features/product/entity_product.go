package product

import (
	"BE-REPO-20/features/user"
	"mime/multipart"
)

type ProductCore struct {
	ID            uint
	UserID        uint
	Name          string `validate:"required"`
	Price         uint   `validate:"required"`
	Quantity      uint   `validate:"required"`
	Description   string
	Category      string
	User          user.UserCore
	ProductImages []ProductImageCore
}

type ProductImageCore struct {
	ID        uint
	ProductID uint
	Url       string
	PublicID  string
}

type ProductDataInterface interface {
	CreateProduct(userId int, input ProductCore) error
	SelectAllProduct(offset, limit int) ([]ProductCore, error)
	SelectProductById(userId int, id int) (*ProductCore, error)
	UpdateProductById(userId int, id int, input ProductCore) error
	DeleteProductById(userId int, id int) error
	SearchProductByQuery(query string, offset, limit int) ([]ProductCore, error)
	CreateProductImage(file multipart.File, input ProductImageCore, nameFile string, id int) error
	DeleteProductImageById(userId, productId, idImage int) error
	ListProductPenjualan(offset, limit int, userId uint) ([]ProductCore, error)
}

type ProductServiceInterface interface {
	CreateProduct(userId int, input ProductCore) error
	SelectAllProduct(page int) ([]ProductCore, error)
	SelectProductById(userId int, id int) (*ProductCore, error)
	UpdateProductById(userId int, id int, input ProductCore) error
	DeleteProductById(userId int, id int) error
	SearchProductByQuery(query string, offset, limit int) ([]ProductCore, error)
	CreateProductImage(file multipart.File, input ProductImageCore, nameFile string, id int) error
	DeleteProductImageById(userId, productId, idImage int) error
	ListProductPenjualan(page int, userId uint) ([]ProductCore, error)
}
