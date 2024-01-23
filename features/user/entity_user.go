package user

import "mime/multipart"

type UserCore struct {
	ID          uint
	UserName    string `validate:"required"`
	ShopName    string
	Email       string `validate:"required,email"`
	PhoneNumber string
	Domicile    string `validate:"required"`
	Address     string
	Image       string
	Province    string
	City        string
	Subdistrict string
	Tagline     string
	ShopImage   string
	Category    string
}

type UserDataInterface interface {
	SelectUser(id int) (*UserCore, error)
	SelectShop(id int) (*UserCore, error)
	UpdateUser(id int, input UserCore, file multipart.File, nameFile string) error
	UpdateShop(id int, input UserCore) error
	Delete(id int) error
}

type UserServiceInterface interface {
	SelectUser(id int) (*UserCore, error)
	SelectShop(id int) (*UserCore, error)
	UpdateUser(id int, input UserCore, file multipart.File, nameFile string) error
	UpdateShop(id int, input UserCore) error
	Delete(id int) error
}
