package order

import (
	"BE-REPO-20/features/user"
)

type OrderCore struct {
	Id         uint
	UserID     uint
	Address    string
	CreditCard uint
	Status     string
	Invoice    string
	Total      uint
	VirtualAcc uint
	User       user.UserCore
	ItemOrders []ItemOrderCore
}

type ItemOrderCore struct {
	ID           uint
	OrderID      uint
	ProductName  string
	ProductPrice uint
	Quantity     uint
	SubTotal     uint
}

type OrderDataInterface interface {
	PostOrder(userId uint, input OrderCore) (*OrderCore, error)
	GetOrder(userId uint) (*OrderCore, error)
}

type OrderServiceInterface interface {
	PostOrder(userId uint, input OrderCore) (*OrderCore, error)
}
