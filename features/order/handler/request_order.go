package handler

import (
	"BE-REPO-20/features/order"
	"strconv"
)

type OrderRequest struct {
	UserId          uint
	Address         string `json:"address"`
	PaymentMethod   string `json:"payment_method"`
	TransactionTime string `json:"transaction_time" form:"transaction_time"`
	Status          string `json:"status"`
	Invoice         string `json:"invoice"`
	// Total      uint   `json:"total"`
	VirtualAcc uint `json:"virtual_acc"`
}

type WebhoocksRequest struct {
	OrderID           string `json:"order_id"`
	TransactionStatus string `json:"transaction_status"`
	SignatureKey      string `json:"signature_key"`
}

type CancelOrderRequest struct {
	Status string `json:"status"`
}

func OrderRequestToCore(input OrderRequest) order.OrderCore {
	return order.OrderCore{
		UserID:          input.UserId,
		Address:         input.Address,
		PaymentMethod:   input.PaymentMethod,
		TransactionTime: input.TransactionTime,
		Status:          input.Status,
		Invoice:         input.Invoice,
		// Total:      input.Total,
		VirtualAcc: input.VirtualAcc,
		// User:       user.UserCore{},
		// ItemOrders: []order.ItemOrderCore{},
	}
}

func RequestToCoreOrder(input OrderRequest) order.OrderCore {
	return order.OrderCore{
		UserID:          input.UserId,
		Address:         input.Address,
		PaymentMethod:   input.PaymentMethod,
		TransactionTime: input.TransactionTime,
		Status:          input.Status,
		Invoice:         input.Invoice,
		// Total:         input.Total,
		VirtualAcc: input.VirtualAcc,
		// User:       user.UserCore{},
		// ItemOrders: []order.ItemOrderCore{},
	}
}

func WebhoocksRequestToCore(input WebhoocksRequest) order.OrderCore {
	orderId, _ := strconv.Atoi(input.OrderID)
	return order.OrderCore{
		Id:     uint(orderId),
		Status: input.TransactionStatus,
	}
}

func CancelRequestToCoreOrder(input CancelOrderRequest) order.OrderCore {
	return order.OrderCore{
		Status: input.Status,
	}
}
