package handler

import (
	"BE-REPO-20/features/product"
)

type ProductRequest struct {
	UserId      uint   `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Quantity    uint   `json:"quantity"`
	Price       uint   `json:"price"`
}

type ProductImageRequest struct {
	ID        uint
	ProductID uint
	Url       string
	PublicID  string
}

func RequestToCore(input ProductRequest) product.ProductCore {
	return product.ProductCore{
		UserID:      input.UserId,
		Name:        input.Name,
		Price:       input.Price,
		Quantity:    input.Quantity,
		Description: input.Description,
		Category:    input.Category,
	}
}

func RequestToCoreImage(input ProductImageRequest) product.ProductImageCore {
	return product.ProductImageCore{
		ProductID: input.ProductID,
		Url:       input.Url,
		PublicID:  input.PublicID,
	}
}
