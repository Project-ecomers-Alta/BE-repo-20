package handler

import (
	_cart "BE-REPO-20/features/cart"
	"BE-REPO-20/features/product/data"
	_dataUser "BE-REPO-20/features/user/data"
)

type CartRequest struct {
	ProductId uint `json:"product_id"`
	UserId    uint `json:"user_id"`
	Quantity  int  `json:"quantity"`
}

func RequestToCore(input CartRequest) _cart.CartCore {
	return _cart.CartCore{
		ProductID: input.ProductId,
		UserID:    input.UserId,
		Quantity:  input.Quantity,
		Product:   data.Product{},
		User:      _dataUser.User{},
	}
}
