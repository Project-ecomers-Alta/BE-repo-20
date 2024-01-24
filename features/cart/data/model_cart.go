package data

import (
	_dataModel "BE-REPO-20/features/product/data"
	_dataUser "BE-REPO-20/features/user/data"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Quantity  int
	ProductID int
	UserID    int
	Product   _dataModel.Product
	User      _dataUser.User
}
