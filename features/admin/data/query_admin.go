package data

import (
	"BE-REPO-20/features/admin"

	"gorm.io/gorm"
)

type adminQuery struct {
	db *gorm.DB
}

func NewAdmin(db *gorm.DB) admin.AdminDataInterface {
	return &adminQuery{
		db: db,
	}
}

// GetUserRoleById implements admin.AdminDataInterface.
func (repo *adminQuery) GetUserRoleById(userId int) (string, error) {
	var user admin.AdminUserCore
	if err := repo.db.Table("users").Where("id = ?", userId).First(&user).Error; err != nil {
		return "", err
	}

	return user.Role, nil
}

// SelectAllUser implements admin.AdminDataInterface.
func (repo *adminQuery) SelectAllUser() ([]admin.AdminUserCore, error) {
	var adminDataGorm []admin.AdminUserCore
	tx := repo.db.Table("users").Find(&adminDataGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// proses mapping dari struct gorm model ke struct core
	var adminsDataCore []admin.AdminUserCore
	for _, value := range adminDataGorm {
		var adminCore = admin.AdminUserCore{
			ID:        value.ID,
			FullName:  value.FullName,
			UserName:  value.UserName,
			Email:     value.Email,
			Domicile:  value.Domicile,
			Role:      value.Role,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		}
		adminsDataCore = append(adminsDataCore, adminCore)
	}

	return adminsDataCore, nil
}

// SearchUserByQuery implements admin.AdminDataInterface.
func (repo *adminQuery) SearchUserByQuery(query string) ([]admin.AdminUserCore, error) {
	var adminDataGorm []admin.AdminUserCore

	tx := repo.db.Table("users").Where("full_name LIKE ?", "%"+query+"%").
		Or("user_name LIKE ?", "%"+query+"%").
		Or("email LIKE ?", "%"+query+"%").
		Find(&adminDataGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Proses mapping dari struct gorm model ke struct core
	var adminsDataCore []admin.AdminUserCore
	for _, value := range adminDataGorm {
		var adminCore = admin.AdminUserCore{
			ID:        value.ID,
			FullName:  value.FullName,
			UserName:  value.UserName,
			Email:     value.Email,
			Domicile:  value.Domicile,
			Role:      value.Role,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		}
		adminsDataCore = append(adminsDataCore, adminCore)
	}

	return adminsDataCore, nil
}

// SelectAllOrder implements admin.AdminDataInterface.
func (repo *adminQuery) SelectAllOrder() ([]admin.AdminItemOrderCore, error) {
	var orderDataGorm []ItemOrder
	tx := repo.db.Preload("Order").Find(&orderDataGorm)

	if tx.Error != nil {
		return nil, tx.Error
	}

	// Process mapping from gorm model struct to core struct
	var ordersDataCore []admin.AdminItemOrderCore
	for _, value := range orderDataGorm {
		// Mapping data dari gorm model ke core struct
		var orderCore = admin.AdminItemOrderCore{
			ID:           value.ID,
			OrderID:      value.OrderID,
			ProductName:  value.ProductName,
			ProductPrice: value.ProductPrice,
			Quantity:     value.Quantity,
			SubTotal:     value.SubTotal,
			CreatedAt:    value.CreatedAt,
			UpdatedAt:    value.UpdatedAt,
			Order: admin.AdminOrderCore{
				ID:            value.Order.ID,
				UserID:        value.Order.UserID,
				Address:       value.Order.Address,
				PaymentMethod: value.Order.PaymentMethod,
				Status:        value.Order.Status,
				Invoice:       value.Order.Invoice,
				Total:         value.Order.Total,
				VirtualAcc:    value.Order.VirtualAcc,
				CreatedAt:     value.Order.CreatedAt,
				UpdatedAt:     value.Order.UpdatedAt,
			},
		}
		ordersDataCore = append(ordersDataCore, orderCore)
	}

	return ordersDataCore, nil
}

// SearchOrderByQuery implements admin.AdminDataInterface.
func (repo *adminQuery) SearchOrderByQuery(query string) ([]admin.AdminItemOrderCore, error) {
	var orderDataGorm []ItemOrder
	var tx *gorm.DB

	if query != "" {
		tx = repo.db.Preload("Order").
			Where("product_name LIKE ? OR order_id LIKE ?", "%"+query+"%", "%"+query+"%").
			Find(&orderDataGorm)
	} else {
		tx = repo.db.Preload("Order").Find(&orderDataGorm)
	}

	if tx.Error != nil {
		return nil, tx.Error
	}

	// Process mapping from gorm model struct to core struct
	var ordersDataCore []admin.AdminItemOrderCore
	for _, value := range orderDataGorm {
		// Mapping data dari gorm model ke core struct
		var orderCore = admin.AdminItemOrderCore{
			ID:           value.ID,
			OrderID:      value.OrderID,
			ProductName:  value.ProductName,
			ProductPrice: value.ProductPrice,
			Quantity:     value.Quantity,
			SubTotal:     value.SubTotal,
			CreatedAt:    value.CreatedAt,
			UpdatedAt:    value.UpdatedAt,
			Order: admin.AdminOrderCore{
				ID:            value.Order.ID,
				UserID:        value.Order.UserID,
				Address:       value.Order.Address,
				PaymentMethod: value.Order.PaymentMethod,
				Status:        value.Order.Status,
				Invoice:       value.Order.Invoice,
				Total:         value.Order.Total,
				VirtualAcc:    value.Order.VirtualAcc,
				CreatedAt:     value.Order.CreatedAt,
				UpdatedAt:     value.Order.UpdatedAt,
			},
		}
		ordersDataCore = append(ordersDataCore, orderCore)
	}

	return ordersDataCore, nil
}
