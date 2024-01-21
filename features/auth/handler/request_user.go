package handler

import "BE-REPO-20/features/auth"

type RegisterRequest struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Domicile string `json:"domicile"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdatePasswordRequest struct {
	Password string `json:"password"`
}

func RequestToUpdatePassword(input UpdatePasswordRequest) auth.AuthCorePassword {
	return auth.AuthCorePassword{
		Password: input.Password,
	}
}

func RequestToCore(input RegisterRequest) auth.AuthCore {
	role := "user"
	if input.Role != "" {
		role = input.Role
	}
	return auth.AuthCore{
		UserName: input.UserName,
		Email:    input.Email,
		Domicile: input.Domicile,
		Password: input.Password,
		Role:     role,
	}
}
