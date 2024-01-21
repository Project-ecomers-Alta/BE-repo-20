package data

import (
	"BE-REPO-20/features/user"
	"errors"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserDataInterface {
	return &userQuery{
		db: db,
	}
}

// SelectShop implements user.UserDataInterface.
func (repo *userQuery) SelectShop(id int) (*user.UserCore, error) {
	panic("unimplemented")
}

// SelectUser implements user.UserDataInterface.
func (repo *userQuery) SelectUser(id int) (*user.UserCore, error) {
	panic("unimplemented")
}

// UpdateShop implements user.UserDataInterface.
func (repo *userQuery) UpdateShop(id int, input user.UserCore) error {
	panic("unimplemented")
}

// UpdateUser implements user.UserDataInterface.
func (repo *userQuery) UpdateUser(id int, input user.UserCore) error {
	userGorm := CoreToModel(input)
	tx := repo.db.Model(&User{}).Where("id = ?", id).Updates(userGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("error not found")
	}
	return nil
}

// Delete implements user.UserDataInterface.
func (repo *userQuery) Delete(id int) error {
	panic("unimplemented")
}
