package handler

import "BE-REPO-20/features/user"

type UserResponse struct {
	UserName    string `json:"user_name"`
	Domicile    string `json:"domicile"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Image       string `json:"image"`
}

func CoreToResponse(data user.UserCore) UserResponse {
	return UserResponse{
		UserName:    data.UserName,
		Domicile:    data.Domicile,
		Email:       data.Email,
		PhoneNumber: data.PhoneNumber,
		Image:       data.Imgage,
	}
}
