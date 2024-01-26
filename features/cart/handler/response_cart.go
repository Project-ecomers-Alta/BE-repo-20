package handler

import (
	_cart "BE-REPO-20/features/cart"
)

type CartResponse struct {
	ID        uint            `json:"id"`
	ProductId uint            `json:"product_id"`
	UserId    uint            `json:"user_id"`
	Quantity  int             `json:"quantity"`
	Product   ProductResponse `json:"product"`
}
type ProductResponse struct {
	ID          uint     `json:"id"`
	UserId      uint     `json:"user_id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Quantity    uint     `json:"quantity"`
	Price       uint     `json:"price"`
	Category    string   `json:"category"`
	User        UserCore `json:"user"`
}
type UserCore struct {
	ID          uint   `json:"id"`
	UserName    string `json:"user_name"`
	Email       string `json:"email"`
	Domicile    string `json:"domicile"`
	PhoneNumber string `json:"phone_number"`
	Image       string `json:"image"`
	Tagline     string `json:"tag_line"`
	Province    string `json:"provinci"`
	City        string `json:"city"`
	Subdistrict string `json:"subdistrict"`
	Address     string `json:"address"`
	Category    string `json:"category"`
}

func CoreToResponse(c _cart.CartCore) CartResponse {
	return CartResponse{
		ID:        c.ID,
		ProductId: c.ProductID,
		UserId:    c.UserID,
		Quantity:  c.Quantity,
		Product: ProductResponse{
			ID:          c.ID,
			UserId:      c.UserID,
			Name:        c.Product.Name,
			Description: c.Product.Description,
			Quantity:    c.Product.Quantity,
			Price:       c.Product.Price,
			Category:    c.Product.Category,
			User: UserCore{
				ID:          c.Product.User.ID,
				UserName:    c.Product.User.UserName,
				Email:       c.Product.User.Email,
				Domicile:    c.Product.User.Domicile,
				PhoneNumber: c.Product.User.PhoneNumber,
				Image:       c.Product.User.Image,
				Tagline:     c.Product.User.TagLine,
				Province:    c.Product.User.Province,
				City:        c.Product.User.City,
				Subdistrict: c.Product.User.Subdistrict,
				Address:     c.Product.User.Address,
				Category:    c.Product.User.Category,
			},
		},
	}
}

func CoreToResponseList(c []_cart.CartCore) []CartResponse {
	var results []CartResponse
	for _, v := range c {
		results = append(results, CoreToResponse(v))
	}
	return results
}
