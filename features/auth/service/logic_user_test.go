package service

import (
	"BE-REPO-20/features/auth"
	"BE-REPO-20/mocks"
	hasMock "BE-REPO-20/utils/encrypts/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	repo := new(mocks.AuthCore)
	hash := new(hasMock.HashService)
	inputData := auth.AuthCore{
		UserName: "alta",
		Email:    "alta@mail.id",
		Domicile: "Jakarta",
		Password: "qwerty",
	}

	t.Run("Success Register", func(t *testing.T) {
		hash.On("HashPassword", inputData.Password).Return("hashed_password", nil).Once()
		repo.On("Register", mock.Anything).Return(
			func(input auth.AuthCore) (*auth.AuthCore, string, error) {
				return &auth.AuthCore{}, "token", nil
			},
		).Once()

		srv := NewAuth(repo, hash)

		res, token, err := srv.Register(inputData)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, "token", token)

		repo.AssertExpectations(t)
	})

	t.Run("Error Register", func(t *testing.T) {
		hash.On("HashPassword", inputData.Password).Return("hashed_password", nil).Once()
		repo.On("Register", mock.Anything).Return(
			func(input auth.AuthCore) (*auth.AuthCore, string, error) {
				return &auth.AuthCore{}, "token", nil
			},
		).Once()

		srv := NewAuth(repo, hash)

		res, token, err := srv.Register(inputData)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, "token", token)

		repo.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	repo := new(mocks.AuthCore)
	hash := new(hasMock.HashService)
	email := "alta@mail.com"
	password := "qwerty123"

	t.Run("Successful Login", func(t *testing.T) {
		expectedUser := &auth.AuthCore{
			Email:    email,
			Password: "hashed_password"}

		repo.On("Login", email, password).Return(expectedUser, nil).Once()
		hash.On("CheckPasswordHash", "hashed_password", password).Return(true).Once()

		srv := NewAuth(repo, hash)

		resultUser, token, err := srv.Login(email, password)

		assert.NoError(t, err)
		assert.NotNil(t, resultUser)
		assert.Equal(t, expectedUser, resultUser)
		assert.NotEmpty(t, token)

		repo.AssertExpectations(t)
	})

	t.Run("Login Failure - Email atau password salah", func(t *testing.T) {
		repo.On("Login", email, password).Return(nil, errors.New("Email atau password salah")).Once()

		srv := NewAuth(repo, hash)

		resultUser, token, err := srv.Login(email, password)

		assert.Error(t, err)
		assert.Nil(t, resultUser)
		assert.Empty(t, token)
		assert.Equal(t, "Email atau password salah", err.Error())

		repo.AssertExpectations(t)
	})

	t.Run("Login Failure - Password Mismatch", func(t *testing.T) {
		expectedUser := &auth.AuthCore{
			Email:    email,
			Password: "hashed_password"}

		repo.On("Login", email, password).Return(expectedUser, nil).Once()
		hash.On("CheckPasswordHash", "hashed_password", password).Return(false).Once()

		srv := NewAuth(repo, hash)

		resultUser, token, err := srv.Login(email, password)

		assert.Error(t, err)
		assert.Nil(t, resultUser)
		assert.Empty(t, token)
		assert.Equal(t, "password tidak sesuai.", err.Error())

		repo.AssertExpectations(t)
	})
}

func TestUpdatePassword(t *testing.T) {
	repo := new(mocks.AuthCore)
	hash := new(hasMock.HashService)
	id := uint(1)

	t.Run("Success Update Password", func(t *testing.T) {
		inputData := auth.AuthCorePassword{
			Password: "new_password",
		}

		hash.On("HashPassword", inputData.Password).Return("hashed_password", nil).Once()
		repo.On("UpdatePassword", id, mock.Anything).Return(nil).Once()

		srv := NewAuth(repo, hash)

		err := srv.UptdatePassword(id, inputData)

		assert.NoError(t, err)

		repo.AssertExpectations(t)
	})

	t.Run("Error Validate", func(t *testing.T) {
		inputData := auth.AuthCorePassword{
			Password: "new_password",
		}

		expectedError := errors.New("validation error")

		hash.On("HashPassword", inputData.Password).Return("hashed_password", nil).Once()
		repo.On("UpdatePassword", id, mock.Anything).Return(expectedError).Once()

		srv := NewAuth(repo, hash)

		err := srv.UptdatePassword(id, inputData)

		assert.EqualError(t, err, expectedError.Error())

		repo.AssertExpectations(t)
	})

	t.Run("Error Hash Password", func(t *testing.T) {
		inputData := auth.AuthCorePassword{
			Password: "new_password",
		}

		expectedError := errors.New("hashing error")

		hash.On("HashPassword", inputData.Password).Return("", expectedError).Once()

		srv := NewAuth(repo, hash)

		err := srv.UptdatePassword(id, inputData)

		assert.EqualError(t, err, "Error hash password.")

		hash.AssertExpectations(t)
		repo.AssertExpectations(t)
	})

	t.Run("Error Update Password", func(t *testing.T) {
		inputData := auth.AuthCorePassword{
			Password: "new_password",
		}

		expectedError := errors.New("update password error")

		hash.On("HashPassword", inputData.Password).Return("hashed_password", nil).Once()
		repo.On("UpdatePassword", id, mock.Anything).Return(expectedError).Once()

		srv := NewAuth(repo, hash)

		err := srv.UptdatePassword(id, inputData)

		assert.EqualError(t, err, expectedError.Error())

		repo.AssertExpectations(t)
	})
}
