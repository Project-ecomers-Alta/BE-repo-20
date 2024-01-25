package data

import (
	"BE-REPO-20/app/middlewares"
	"BE-REPO-20/features/auth"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type authQuery struct {
	db *gorm.DB
}

// CheckPassword implements auth.AuthDataInterface.
func (repo *authQuery) CheckPassword(savedPassword, inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(savedPassword), []byte(inputPassword))
	return err == nil
}

func NewAuth(db *gorm.DB) auth.AuthDataInterface {
	return &authQuery{
		db: db,
	}
}

func (repo *authQuery) Register(input auth.AuthCore) (data *auth.AuthCore, token string, err error) {
	inputRegisterGorm := User{
		UserName: input.UserName,
		Email:    input.Email,
		Domicile: input.Domicile,
		Password: input.Password,
	}

	tx := repo.db.Create(&inputRegisterGorm)
	if tx.Error != nil {
		return nil, "", tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, "", errors.New("insert failed, row affected = 0")
	}

	var authGorm User
	tx = repo.db.Where("email = ?", input.Email).First(&authGorm)
	if tx.Error != nil {
		return nil, "", tx.Error
	}

	result := authGorm.ModelToCore()

	generatedToken, err := middlewares.CreateToken(int(result.ID))
	if err != nil {
		return nil, "", err
	}

	return &result, generatedToken, nil
}

func (repo *authQuery) Login(email string, password string) (data *auth.AuthCore, err error) {
	var authGorm User
	tx := repo.db.Where("email = ?", email).First(&authGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	result := authGorm.ModelToCore()

	return &result, nil
}

func (repo *authQuery) UpdatePassword(id uint, input auth.AuthCorePassword) error {
	authInput := User{
		Password: input.Password,
	}

	tx := repo.db.Model(&User{}).Where("id = ?", id).Updates(&authInput)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("edit failed, row affected = 0")
	}

	return nil
}
