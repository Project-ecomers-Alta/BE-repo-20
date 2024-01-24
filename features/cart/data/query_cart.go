package data

import (
	_cart "BE-REPO-20/features/cart"

	"gorm.io/gorm"
)

type cartQuery struct {
	db *gorm.DB
}

func NewCart(db *gorm.DB) _cart.CartDataInterface {
	return &cartQuery{
		db: db,
	}
}

func (repo *cartQuery) CreateCart(userId int, productId uint) error {
	var existingCart _cart.CartCore
	result := repo.db.Where(&_cart.CartCore{UserID: uint(userId), ProductID: productId}).First(&existingCart)

	if result.Error == nil {
		existingCart.Quantity++
		result = repo.db.Save(&existingCart)
	} else {
		newCart := _cart.CartCore{
			UserID:    uint(userId),
			ProductID: productId,
			Quantity:  1,
		}
		result = repo.db.Create(&newCart)
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}
