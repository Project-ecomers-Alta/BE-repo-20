package service

import (
	_cart "BE-REPO-20/features/cart"
)

type cartService struct {
	cartData _cart.CartDataInterface
}

func NewCart(repo _cart.CartDataInterface) _cart.CartServiceInterface {
	return &cartService{
		cartData: repo,
	}
}

func (service *cartService) DeleteCarts(ids []uint) error {
	err := service.cartData.DeleteCarts(ids)
	return err
}

func (service *cartService) CreateCart(userId int, productId uint) error {
	err := service.cartData.CreateCart(userId, productId)
	return err
}

func (service *cartService) SelectAllCart(userId uint) ([]_cart.CartCore, error) {
	result, err := service.cartData.SelectAllCart(userId)
	return result, err
}
