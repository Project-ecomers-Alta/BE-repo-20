package handler

import (
	"BE-REPO-20/features/product"
)

type ProductRequest struct {
	UserId      uint
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	Price       uint   `json:"price"`
}

func RequestToCore(input ProductRequest) product.ProductCore {
	return product.ProductCore{
		UserID:      input.UserId,
		Name:        input.Name,
		Price:       input.Price,
		Quantity:    input.Quantity,
		Description: input.Description,
	}
}
