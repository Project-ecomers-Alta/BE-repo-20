package data

import (
	"BE-REPO-20/features/product"
	"BE-REPO-20/features/user"
	_userData "BE-REPO-20/features/user/data"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	UserID      uint
	Name        string `gorm:"default:null"`
	Price       uint   `gorm:"default:null"`
	Quantity    uint   `gorm:"default:null"`
	Description string `gorm:"default:null"`
	Category    string `gorm:"default:null"`
	User        _userData.User
}

func (u Product) ModelToCore() product.ProductCore {
	return product.ProductCore{
		ID:          u.ID,
		UserID:      u.UserID,
		Name:        u.Name,
		Price:       u.Price,
		Quantity:    u.Quantity,
		Category:    u.Category,
		Description: u.Description,
		User: user.UserCore{
			ID:          u.User.ID,
			UserName:    u.User.UserName,
			Email:       u.User.Email,
			PhoneNumber: u.User.PhoneNumber,
			Domicile:    u.User.Domicile,
			Address:     u.User.Address,
			Image:       u.User.Image,
			Province:    u.User.Province,
			City:        u.User.City,
			Subdistrict: u.User.Subdistrict,
			Tagline:     u.User.TagLine,
			Category:    u.User.Category,
		},
	}
}

func ModelToCoreList(data []Product) []product.ProductCore {
	var results []product.ProductCore
	for _, v := range data {
		results = append(results, v.ModelToCore())
	}
	return results
}

func CoreToModel(p product.ProductCore) Product {
	return Product{
		UserID:      p.UserID,
		Name:        p.Name,
		Price:       p.Price,
		Quantity:    p.Quantity,
		Category:    p.Category,
		Description: p.Description,
	}
}
