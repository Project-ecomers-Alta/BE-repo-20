package service

import (
	"BE-REPO-20/app/middlewares"
	"BE-REPO-20/features/auth"
	"BE-REPO-20/utils/encrypts"
	"errors"

	"github.com/go-playground/validator"
)

type authService struct {
	authData    auth.AuthDataInterface
	hashService encrypts.HashInterface
	validate    *validator.Validate
}

func NewAuth(repo auth.AuthDataInterface, hash encrypts.HashInterface) auth.AuthServiceInterface {
	return &authService{
		authData:    repo,
		hashService: hash,
		validate:    validator.New(),
	}
}

func (service *authService) Register(input auth.AuthCore) (data *auth.AuthCore, token string, err error) {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return nil, "", errValidate
	}

	if input.Password != "" {
		hashedPass, errHash := service.hashService.HashPassword(input.Password)
		if errHash != nil {
			return nil, "", errors.New("rror hashing password")
		}
		input.Password = hashedPass
	}

	data, generatedToken, err := service.authData.Register(input)
	return data, generatedToken, err
}

func (service *authService) Login(email string, password string) (data *auth.AuthCore, token string, err error) {
	if email == "" || password == "" {
		return nil, "", errors.New("email dan password wajib diisi")
	}

	data, err = service.authData.Login(email, password)
	if err != nil {
		return nil, "", errors.New("Email atau password salah")
	}
	isValid := service.hashService.CheckPasswordHash(data.Password, password)
	if !isValid {
		return nil, "", errors.New("password tidak sesuai.")
	}

	token, errJwt := middlewares.CreateToken(int(data.ID))
	if errJwt != nil {
		return nil, "", errJwt
	}
	return data, token, err
}

func (service *authService) UptdatePassword(id uint, input auth.AuthCorePassword) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}

	if input.Password != "" {
		hashedPass, errHash := service.hashService.HashPassword(input.Password)
		if errHash != nil {
			return errors.New("Error hash password.")
		}
		input.Password = hashedPass
	}

	err := service.authData.UpdatePassword(id, input)
	return err
}
