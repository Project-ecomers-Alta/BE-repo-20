package data

import (
	_cart "BE-REPO-20/features/cart"
	"errors"

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

func (repo *cartQuery) DeleteCarts(ids []uint) error {
	tx := repo.db.Where("id IN ?", ids).Delete(&Cart{})
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected != int64(len(ids)) {
		return errors.New("delete failed, not all rows affected")
	}

	return nil
}

func (repo *cartQuery) SelectAllCart(userId uint) ([]_cart.CartCore, error) {
	var cartData []Cart
	// tx := repo.db.Where("user_id = ?", userId).Find(&cartData)
	tx := repo.db.Preload("Product").Preload("User").Find(&cartData, "user_id = ?", userId)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var cartDataCore []_cart.CartCore
	for _, value := range cartData {
		var cartCore = _cart.CartCore{
			ID:        value.ID,
			ProductID: uint(value.ProductID),
			UserID:    uint(value.UserID),
			Quantity:  value.Quantity,
			// Product: []product.ProductCore,
		}
		cartDataCore = append(cartDataCore, cartCore)
	}

	return cartDataCore, nil
}

func (repo *cartQuery) CreateCart(userId int, productId uint) error {
	var existingCart Cart
	result := repo.db.Where(&Cart{UserID: userId, ProductID: int(productId)}).First(&existingCart)

	if result.Error == nil {
		existingCart.Quantity++
		result = repo.db.Save(&existingCart)
	} else {
		newCart := Cart{
			UserID:    userId,
			ProductID: int(productId),
			Quantity:  1,
		}
		result = repo.db.Create(&newCart)
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}
