package handler

import (
	"BE-REPO-20/features/admin"
)

type AdminUserResponse struct {
	ID        uint   `json:"id"`
	FullName  string `json:"full_name"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	Domicile  string `json:"domicile"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func CoreToResponse(data admin.AdminUserCore) AdminUserResponse {
	return AdminUserResponse{
		ID:        data.ID,
		FullName:  data.FullName,
		UserName:  data.UserName,
		Email:     data.Email,
		Domicile:  data.Domicile,
		Role:      data.Role,
		CreatedAt: data.CreatedAt.Format("2006-01-02"),
		UpdatedAt: data.UpdatedAt.Format("2006-01-02"),
	}
}

func CoreToResponseList(data []admin.AdminUserCore) []AdminUserResponse {
	var results []AdminUserResponse
	for _, v := range data {
		results = append(results, CoreToResponse(v))
	}
	return results
}
