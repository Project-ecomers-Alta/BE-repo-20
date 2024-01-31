package auth

import (
	_dataMode "BE-REPO-20/features/product/data"
	_dataModel "BE-REPO-20/features/user/data"
)

type CartCore struct {
	ID        uint
	ProductID uint
	UserID    uint
	Quantity  int
	Product   _dataMode.Product
	User      _dataModel.User
}

type CartDataInterface interface {
	CreateCart(input CartCore) error
	SelectAllCart(userId uint) ([]CartCore, error)
	DeleteCarts(ids []uint) error
}

type CartServiceInterface interface {
	CreateCart(input CartCore) error
	SelectAllCart(userId uint) ([]CartCore, error)
	DeleteCarts(ids []uint) error
}
