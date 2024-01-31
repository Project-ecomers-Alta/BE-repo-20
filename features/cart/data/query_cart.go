package data

import (
	_cart "BE-REPO-20/features/cart"
	_dataMode "BE-REPO-20/features/product/data"
	_dataModel "BE-REPO-20/features/user/data"
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

func (repo *cartQuery) SelectAllCart(userId uint) ([]_cart.CartCore, error) {
	var cartData []Cart
	tx := repo.db.Preload("Product").Preload("Product.User").Where("user_id = ?", userId).Find(&cartData)
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
			Product: _dataMode.Product{
				UserID:      value.Product.UserID,
				Name:        value.Product.Name,
				Description: value.Product.Description,
				Quantity:    value.Product.Quantity,
				Price:       value.Product.Price,
				Category:    value.Product.Category,
			},
			User: _dataModel.User{
				UserName:    value.User.UserName,
				Email:       value.User.Email,
				Domicile:    value.User.Domicile,
				PhoneNumber: value.User.PhoneNumber,
				Image:       value.User.Image,
				TagLine:     value.User.TagLine,
				Province:    value.User.Province,
				City:        value.User.City,
				Subdistrict: value.User.Subdistrict,
				Address:     value.User.Address,
				Category:    value.User.Category,
			},
		}
		cartDataCore = append(cartDataCore, cartCore)
	}

	return cartDataCore, nil
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

func (repo *cartQuery) CreateCart(input _cart.CartCore) error {
	var existingCart Cart
	result := repo.db.Where(&Cart{UserID: int(input.UserID), ProductID: int(input.ProductID)}).First(&existingCart)

	if result.Error == nil {
		existingCart.Quantity += 1
		result = repo.db.Save(&existingCart)
	} else {
		cartInput := Cart{
			UserID:    int(input.UserID),
			ProductID: int(input.ProductID),
			Quantity:  1,
		}
		result = repo.db.Create(&cartInput)
	}

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}
	return nil
}
