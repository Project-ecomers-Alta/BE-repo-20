package data

import (
	"BE-REPO-20/features/cart/data"
	"BE-REPO-20/features/order"
	"errors"

	"gorm.io/gorm"
)

type orderQuery struct {
	db *gorm.DB
}

func NewOrder(db *gorm.DB) order.OrderDataInterface {
	return &orderQuery{
		db: db,
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
	var itemOrders []ItemOrder
	// get cart
	carts, errCart := repo.GetCart(uint(userId))
	if errCart != nil {
		return nil, errCart
	}
	var amount int

	for i := 0; i < len(carts); i++ {
		amount += carts[i].Quantity * int(carts[i].Product.Price)
	}

	orderGorm.Total = uint(amount)
	// repo.db.Transaction(
	repo.db.Transaction(func(tx *gorm.DB) error {
		// Create Data Order
		orderGorm = OrderCoreToModel(input)
		orderGorm.Total = uint(amount)
		if errOrder := tx.Create(&orderGorm).Error; errOrder != nil {
			return errOrder
		}

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
		return nil
	})

	// get order id
	orders, errOrder := repo.GetOrder(orderGorm.ID)
	if errOrder != nil {
		return nil, errOrder
	}

	return orders, nil
}
