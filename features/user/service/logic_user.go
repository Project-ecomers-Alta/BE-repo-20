package service

import (
	"BE-REPO-20/features/user"
	"errors"
)

type userService struct {
	userData user.UserDataInterface
}

func NewUser(repo user.UserDataInterface) user.UserServiceInterface {
	return &userService{
		userData: repo,
	}
}

// SelectShop implements user.UserServiceInterface.
func (service *userService) SelectShop(id int) (*user.UserCore, error) {
	panic("unimplemented")
}

// SelectUser implements user.UserServiceInterface.
func (service *userService) SelectUser(id int) (*user.UserCore, error) {
	panic("unimplemented")
}

// UpdateShop implements user.UserServiceInterface.
func (service *userService) UpdateShop(id int, input user.UserCore) error {
	panic("unimplemented")
}

// UpdateUser implements user.UserServiceInterface.
func (service *userService) UpdateUser(id int, input user.UserCore) error {
	if id <= 0 {
		return errors.New("invalid id")
	}
	// fmt.Println(input)
	err := service.userData.UpdateUser(id, input)
	return err
}

// Delete implements user.UserServiceInterface.
func (service *userService) Delete(id int) error {
	panic("unimplemented")
}
