package product

import (
	"BE-REPO-20/features/user"
)

type ProductCore struct {
	ID          uint
	UserID      uint
	Name        string `validate:"required"`
	Price       uint   `validate:"required"`
	Quantity    uint   `validate:"required"`
	Description string
	Category    string
	User        user.UserCore
}

type ProductDataInterface interface {
	CreateProduct(userId int, input ProductCore) error
	SelectAllProduct() ([]ProductCore, error)
	SelectProductById(userId int, id int) (*ProductCore, error)
	UpdateProductById(userId int, id int, input ProductCore) error
	DeleteProductById(userId int, id int) error
	SearchProductByQuery(query string) ([]ProductCore, error)
}

type ProductServiceInterface interface {
	CreateProduct(userId int, input ProductCore) error
	SelectAllProduct() ([]ProductCore, error)
	SelectProductById(userId int, id int) (*ProductCore, error)
	UpdateProductById(userId int, id int, input ProductCore) error
	DeleteProductById(userId int, id int) error
	SearchProductByQuery(query string) ([]ProductCore, error)
}
