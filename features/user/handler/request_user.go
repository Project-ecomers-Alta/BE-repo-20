package handler

import "BE-REPO-20/features/user"

type UserRequest struct {
	UserName    string `json:"user_name" form:"username"`
	Domicile    string `json:"domicile"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Image       string `json:"image_url" form:"image"`
}

// type UserRequest struct {
// 	FullName    string `json:"shop_name" form:"shop_name"`
// 	Tagline     string `json:"tagline" form:"tagline"`
// 	Province    string `json:"province" form:"province"`
// 	City        string `json:"city" form:"city"`
// 	Subdistrict string `json:"subdistrict" form:"subdistrict"`
// 	Address     string `json:"address" form:"address"`
// }

func RequestToCore(input UserRequest) user.UserCore {
	return user.UserCore{
		UserName:    input.UserName,
		Domicile:    input.Domicile,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		Image:       input.Image,
	}
}
