package service

import (
	_cart "BE-REPO-20/features/cart"
)

type cartService struct {
	cartData _cart.CartDataInterface
}

// GetCart implements auth.CartServiceInterface.

func NewCart(repo _cart.CartDataInterface) _cart.CartServiceInterface {
	return &cartService{
		cartData: repo,
	}
}

func (service *cartService) CreateCart(userId int, productId uint) error {
	err := service.cartData.CreateCart(userId, productId)
	return err
}

