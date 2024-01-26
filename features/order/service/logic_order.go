package service

import (
	_order "BE-REPO-20/features/order"
)

type orderService struct {
	orderData _order.OrderDataInterface
}

func NewOrder(repo _order.OrderDataInterface) _order.OrderServiceInterface {
	return &orderService{
		orderData: repo,
	}
}

// PostOrder implements order.OrderServiceInterface.
func (service *orderService) PostOrder(userId uint, input _order.OrderCore) (*_order.OrderCore, error) {

	res, err := service.orderData.PostOrder(userId, input)
	if err != nil {
		return nil, err
	}

	return res, nil
}
