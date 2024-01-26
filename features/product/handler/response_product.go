package handler

import (
	"BE-REPO-20/features/product"
)

type ProductResponse struct {
	ID            uint                 `json:"id"`
	UserId        uint                 `json:"user_id"`
	Name          string               `json:"name"`
	Description   string               `json:"description"`
	Quantity      uint                 `json:"quantity"`
	Price         uint                 `json:"price"`
	Category      string               `json:"category"`
	User          UserCore             `json:"user"`
	DetailsImages []ProductImageDetail `json:"details_images"`
}

type ProductImageDetail struct {
	ID        uint   `json:"id"`
	ProductID uint   `json:"product_id"`
	Images    string `json:"images"`
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
	ShopImage   string `json:"shop_image"`
}

func CoreToResponse(p product.ProductCore) ProductResponse {
	return ProductResponse{
		ID:            p.ID,
		UserId:        p.UserID,
		Name:          p.Name,
		Description:   p.Description,
		Category:      p.Category,
		Quantity:      p.Quantity,
		Price:         p.Price,
		DetailsImages: ProductImageCoreToResponseDetails(p.ProductImages),
		User: UserCore{
			ID:          p.User.ID,
			UserName:    p.User.UserName,
			Email:       p.User.Email,
			PhoneNumber: p.User.PhoneNumber,
			Domicile:    p.User.Domicile,
			Address:     p.User.Address,
			Image:       p.User.Image,
			Province:    p.User.Province,
			City:        p.User.City,
			Subdistrict: p.User.Subdistrict,
			Tagline:     p.User.Tagline,
			ShopImage:   p.User.ShopImage,
		},
	}
}

func CoreToResponseUpdate(p product.ProductCore) ProductResponse {
	return ProductResponse{
		ID:          p.ID,
		UserId:      p.UserID,
		Name:        p.Name,
		Description: p.Description,
		Category:    p.Category,
		Quantity:    p.Quantity,
		Price:       p.Price,
	}
}

func CoreToResponseList(p []product.ProductCore) []ProductResponse {
	var results []ProductResponse
	for _, v := range p {
		results = append(results, CoreToResponse(v))
	}
	return results
}

func ProductImageCoreToResponseDetails(p []product.ProductImageCore) []ProductImageDetail {
	var results []ProductImageDetail
	for _, v := range p {
		results = append(results, ProductImageDetail{
			ID:        v.ID,
			ProductID: v.ProductID,
			Images:    v.Url,
		})
	}
	return results
}
