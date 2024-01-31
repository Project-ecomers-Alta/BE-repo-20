package service

import (
	"BE-REPO-20/features/order"
	_order "BE-REPO-20/features/order"
	"errors"
	// _midtransService "BE-REPO-20/features/midtrans/service"
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

	//midtrans
	// midtransResponse := _midtransService.MidtransService.CreateEcho()
	return res, nil
}

// GetOrder implements order.OrderServiceInterface.
func (service *orderService) GetOrders(userId uint) ([]_order.OrderCore, error) {
	results, err := service.orderData.GetOrders(userId)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (service *orderService) WebhoocksService(webhoocksReq order.OrderCore) error {
	if webhoocksReq.Id == 0 {
		return errors.New("invalid order id")
	}

	err := service.orderData.WebhoocksData(webhoocksReq)
	if err != nil {
		return err
	}

	return nil
}
