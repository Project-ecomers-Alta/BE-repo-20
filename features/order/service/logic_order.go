package service

import (
	"BE-REPO-20/features/order"
	_order "BE-REPO-20/features/order"
	"errors"
	"math/rand"
	"strconv"
	"time"
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
	if userId <= 0 {
		return nil, errors.New("invalid id")
	}

	t := time.Now()
	year := strconv.Itoa(t.Year())
	month := int(t.Month())
	day := strconv.Itoa(t.Day())
	randomNumb := strconv.Itoa(rand.Intn(100000))

	input.Invoice = year + strconv.Itoa(month) + day + "/" + randomNumb
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

// CancelOrder implements order.OrderServiceInterface.
func (os *orderService) CancelOrder(userIdLogin int, orderId string, orderCore order.OrderCore) error {
	if orderCore.Status == "" {
		orderCore.Status = "cancelled"
	}

	err := os.orderData.CancelOrder(userIdLogin, orderId, orderCore)
	return err
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
