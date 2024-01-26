package handler

import (
	"BE-REPO-20/features/order"
)

type OrderRequest struct {
	UserId     uint
	Address    string `json:"address"`
	CreditCard uint   `json:"credit_card"`
	Status     string `json:"status"`
	Invoice    string `json:"invoice"`
	// Total      uint   `json:"total"`
	VirtualAcc uint `json:"virtual_acc"`
}

func OrderRequestToCore(input OrderRequest) order.OrderCore {
	return order.OrderCore{
		UserID:     input.UserId,
		Address:    input.Address,
		CreditCard: input.CreditCard,
		Status:     input.Status,
		Invoice:    input.Invoice,
		// Total:      input.Total,
		VirtualAcc: input.VirtualAcc,
		// User:       user.UserCore{},
		// ItemOrders: []order.ItemOrderCore{},
	}
}
