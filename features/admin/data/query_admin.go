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
