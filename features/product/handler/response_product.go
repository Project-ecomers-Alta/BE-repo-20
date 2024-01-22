package handler

import (
	"BE-REPO-20/features/product"
	"BE-REPO-20/features/user"
)

type ProductResponse struct {
	UserId      uint
	Name        string
	Description string
	Quantity    uint
	Price       uint
	Category    string
	User        user.UserCore
}

func CoreToResponse(p product.ProductCore) ProductResponse {
	return ProductResponse{
		UserId:      p.UserID,
		Name:        p.Name,
		Description: p.Description,
		Category:    p.Category,
		Quantity:    p.Quantity,
		Price:       p.Price,
		User: user.UserCore{
			ID:       p.User.ID,
			UserName: p.User.UserName,
			// ShopName:    p.User.ShopName,
			Email:       p.User.Email,
			PhoneNumber: p.User.PhoneNumber,
			Domicile:    p.User.Domicile,
			Address:     p.User.Address,
			Image:       p.User.Image,
			Province:    p.User.Province,
			City:        p.User.City,
			Subdistrict: p.User.Subdistrict,
			Tagline:     p.User.Tagline,
			// ShopImage:   p.User.ShopImage,
			Category: p.User.Category,
		},
	}
}

func CoreToResponseList(p []product.ProductCore) []ProductResponse {
	var results []ProductResponse
	for _, v := range p {
		results = append(results, CoreToResponse(v))
	}
	return results
}
