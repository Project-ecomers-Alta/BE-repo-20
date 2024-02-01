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

type AdminOrderResponse struct {
	ID            uint   `json:"id"`
	UserID        uint   `json:"user_id"`
	Address       string `json:"address"`
	PaymentMethod string `json:"payment_method"`
	Status        string `json:"status"`
	Invoice       string `json:"invoice"`
	Total         uint   `json:"total"`
	VirtualAcc    string `json:"virtual_acc"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

type AdminItemOrderResponse struct {
	ID           uint               `json:"id"`
	OrderID      uint               `gorm:"column:order_id" json:"order_id"`
	ProductName  string             `json:"product_name"`
	ProductPrice uint               `json:"product_price"`
	Quantity     uint               `json:"quantity"`
	SubTotal     uint               `json:"sub_total"`
	CreatedAt    string             `json:"created_at"`
	UpdatedAt    string             `json:"updated_at"`
	Order        AdminOrderResponse `gorm:"foreignKey:OrderID" json:"order"`
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

func CoreToOrderResponse(data admin.AdminOrderCore) AdminOrderResponse {
	return AdminOrderResponse{
		ID:            data.ID,
		UserID:        data.UserID,
		Address:       data.Address,
		PaymentMethod: data.PaymentMethod,
		Status:        data.Status,
		Invoice:       data.Invoice,
		Total:         data.Total,
		VirtualAcc:    data.VirtualAcc,
		CreatedAt:     data.CreatedAt.Format("2006-01-02"),
		UpdatedAt:     data.UpdatedAt.Format("2006-01-02"),
	}
}

func CoreToOrderResponseList(data []admin.AdminOrderCore) []AdminOrderResponse {
	var results []AdminOrderResponse
	for _, v := range data {
		results = append(results, CoreToOrderResponse(v))
	}
	return results
}

func CoreToItemOrderResponse(data admin.AdminItemOrderCore) AdminItemOrderResponse {
	return AdminItemOrderResponse{
		ID:           data.ID,
		OrderID:      data.OrderID,
		ProductName:  data.ProductName,
		ProductPrice: data.ProductPrice,
		Quantity:     data.Quantity,
		SubTotal:     data.SubTotal,
		CreatedAt:    data.CreatedAt.Format("2006-01-02"),
		UpdatedAt:    data.UpdatedAt.Format("2006-01-02"),
		Order:        CoreToOrderResponse(data.Order),
	}
}

func CoreToItemOrderResponseList(data []admin.AdminItemOrderCore) []AdminItemOrderResponse {
	var results []AdminItemOrderResponse
	for _, v := range data {
		results = append(results, CoreToItemOrderResponse(v))
	}
	return results
}
