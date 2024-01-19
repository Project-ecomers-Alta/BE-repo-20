package service

import (
	"BE-REPO-20/features/user"
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
	data, err := service.userData.SelectUser(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// UpdateShop implements user.UserServiceInterface.
func (service *userService) UpdateShop(id int, input user.UserCore) error {
	panic("unimplemented")
}

// UpdateUser implements user.UserServiceInterface.
func (service *userService) UpdateUser(id int, input user.UserCore) error {
	panic("unimplemented")
}

// Delete implements user.UserServiceInterface.
func (service *userService) Delete(id int) error {
	panic("unimplemented")
}
