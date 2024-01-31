package data

import (
	"BE-REPO-20/features/cart/data"
	"BE-REPO-20/features/order"
	"BE-REPO-20/features/user"
	_userData "BE-REPO-20/features/user/data"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID          uint
	Address         string
	PaymentMethod   string
	TransactionTime string
	Status          string
	Invoice         string
	Total           uint
	VirtualAcc      uint
	User            _userData.User
	ItemOrders      []ItemOrder
}

type ItemOrder struct {
	gorm.Model
	OrderID      uint
	ProductName  string
	ProductPrice uint
	Quantity     uint
	SubTotal     uint
}

func CartToItemOrder(u data.Cart) ItemOrder {
	return ItemOrder{
		ProductName:  u.Product.Name,
		ProductPrice: u.Product.Price,
		Quantity:     uint(u.Quantity),
		SubTotal:     uint(u.Quantity) * u.Product.Price,
	}
}

func CartToItemOrderList(cart []data.Cart) []ItemOrder {
	var results []ItemOrder
	for _, v := range cart {
		results = append(results, CartToItemOrder(v))
	}
	return results
}

func OrderCoreToModel(o order.OrderCore) Order {
	return Order{
		UserID:          o.UserID,
		Address:         o.Address,
		PaymentMethod:   o.PaymentMethod,
		TransactionTime: o.TransactionTime,
		Status:          o.Status,
		Invoice:         o.Invoice,
		// Total:      o.Total,
		VirtualAcc: o.VirtualAcc,
		// User:       _userData.User{},
		// ItemOrders: ItemOrder{},
	}
}

func ModelToCore(o Order) order.OrderCore {
	return order.OrderCore{
		Id:              o.ID,
		UserID:          o.UserID,
		Address:         o.Address,
		PaymentMethod:   o.PaymentMethod,
		TransactionTime: o.TransactionTime,
		Status:          o.Status,
		Invoice:         o.Invoice,
		Total:           o.Total,
		VirtualAcc:      o.VirtualAcc,
		User: user.UserCore{
			ID:          o.User.ID,
			UserName:    o.User.UserName,
			ShopName:    o.User.ShopName,
			Email:       o.User.Email,
			PhoneNumber: o.User.PhoneNumber,
			Domicile:    o.User.Domicile,
			Address:     o.User.Address,
			Image:       o.User.Image,
			Province:    o.User.Province,
			City:        o.User.City,
			Subdistrict: o.User.Subdistrict,
			Tagline:     o.User.TagLine,
			ShopImage:   o.User.ShopImage,
		},
		ItemOrders: ItemOrderGormToCore(o.ItemOrders),
	}
}

func ItemOrderGormToCore(itemOrder []ItemOrder) []order.ItemOrderCore {
	var results []order.ItemOrderCore
	for _, v := range itemOrder {
		results = append(results, order.ItemOrderCore{
			ID:           v.ID,
			OrderID:      v.OrderID,
			ProductName:  v.ProductName,
			ProductPrice: v.ProductPrice,
			Quantity:     v.Quantity,
			SubTotal:     v.SubTotal,
		})
	}
	return results
}

func TotalAmount(o []data.Cart) int {
	var amount int
	for _, v := range o {
		amount += int(v.Quantity * int(v.Product.Price))
	}
	return amount
}

func WebhoocksCoreToModel(reqNotif order.OrderCore) Order {
	return Order{
		Status: reqNotif.Status,
	}
}
