package handler

import (
	"BE-REPO-20/features/order"
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
	OrderID           uint   `json:"order_id"`
	TransactionStatus string `json:"transaction_status"`
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
	return order.OrderCore{
		Id:     input.OrderID,
		Status: input.TransactionStatus,
	}
}
