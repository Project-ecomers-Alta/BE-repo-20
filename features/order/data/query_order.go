package data

import (
	"BE-REPO-20/features/cart/data"
	"BE-REPO-20/features/order"
	"BE-REPO-20/utils/midtrans"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type orderQuery struct {
	db              *gorm.DB
	paymentMidtrans midtrans.MidtransInterface
}

func NewOrder(db *gorm.DB, mid midtrans.MidtransInterface) order.OrderDataInterface {
	return &orderQuery{
		db:              db,
		paymentMidtrans: mid,
	}
}

func (repo *orderQuery) GetCart(userId uint) ([]data.Cart, error) {
	var cartGorm []data.Cart
	tx := repo.db.Preload("Product").Preload("User").Find(&cartGorm, "user_id = ?", userId)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, errors.New("insert order failed, row affected = 0")
	}
	return cartGorm, nil
}

func (repo *orderQuery) GetOrder(userId uint) (*order.OrderCore, error) {
	var orderGorm Order
	tx := repo.db.Preload("ItemOrders").Preload("User").First(&orderGorm, "id = ?", userId)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, errors.New("find order failed, row affected = 0")
	}

	orderCore := ModelToCore(orderGorm)
	return &orderCore, nil
}

func (repo *orderQuery) GetOrders(userId uint) ([]order.OrderCore, error) {
	var orderGorm []Order
	tx := repo.db.Preload("ItemOrders").Preload("User").Find(&orderGorm, "user_id = ?", userId)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, errors.New("find order failed, row affected = 0")
	}
	var orderCores []order.OrderCore
	for _, v := range orderGorm {
		orderCores = append(orderCores, ModelToCore(v))
	}

	return orderCores, nil
}

// PostOrder implements order.OrderDataInterface.
func (repo *orderQuery) PostOrder(userId uint, input order.OrderCore) (*order.OrderCore, error) {
	var orderGorm Order
	var lastOrder Order
	var itemOrders []ItemOrder
	var order order.OrderCore
	// get cart
	carts, errCart := repo.GetCart(uint(userId))
	if errCart != nil {
		return nil, errCart
	}
	var amount int

	for i := 0; i < len(carts); i++ {
		amount += carts[i].Quantity * int(carts[i].Product.Price)
	}

	result := repo.db.Last(&lastOrder)
	if result.RowsAffected == 0 {
		lastOrder.ID = 0
	} else {
		tx := repo.db.Last(&lastOrder)
		if tx.Error != nil {
			return nil, tx.Error
		}
	}
	fmt.Println(lastOrder.ID)
	input.Total = uint(amount)
	input.UserID = userId
	input.Id = lastOrder.ID + 1
	fmt.Println(input.Id)

	payment, errPay := repo.paymentMidtrans.Order(input)
	if errPay != nil {
		return nil, errPay
	}

	// repo.db.Transaction(
	repo.db.Transaction(func(tx *gorm.DB) error {
		// Create Data Order
		orderGorm = OrderCoreToModel(input)
		orderGorm.PaymentMethod = payment.PaymentMethod
		orderGorm.Status = payment.Status
		orderGorm.VirtualAcc = payment.VirtualAcc
		orderGorm.TransactionTime = payment.TransactionTime
		orderGorm.Total = uint(amount)
		if errOrder := tx.Create(&orderGorm).Error; errOrder != nil {
			return errOrder
		}

		// 	payment, errPay := repo.paymentMidtrans.Order(orderGorm)
		// if errPay != nil {
		// 	return nil, errPay
		// }

		// // Create Data ItemOrder
		itemOrders = CartToItemOrderList(carts)
		for i := range itemOrders {
			itemOrders[i].OrderID = orderGorm.ID
		}
		if errItemOrder := tx.Create(&itemOrders).Error; errItemOrder != nil {
			return errItemOrder
		}

		// Delete data from Cart
		if errDelCart := tx.Where("user_id = ?", userId).Delete(&data.Cart{}).Error; errDelCart != nil {
			return errDelCart
		}
		// return nil

		return nil
	})
	// get order id
	orders, _ := repo.GetOrder(orderGorm.ID)
	order = *orders

	return &order, nil
	// return payment, nil
}

func (repo *orderQuery) WebhoocksData(webhoocksReq order.OrderCore) error {
	dataGorm := WebhoocksCoreToModel(webhoocksReq)
	tx := repo.db.Model(&Order{}).Where("id = ?", webhoocksReq.Id).Updates(dataGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found ")
	}
	return nil
}

// SelectOrderAdmin implements order.OrderDataInterface.
func (repo *orderQuery) CancelOrder(userId int, orderId string, orderCore order.OrderCore) error {
	if orderCore.Status == "cancelled" {
		repo.paymentMidtrans.CancelOrder(orderId)
	}

	dataGorm := Order{
		Status: orderCore.Status,
	}
	tx := repo.db.Model(&Order{}).Where("id = ? AND user_id = ?", orderId, userId).Updates(dataGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found ")
	}
	return nil
}
