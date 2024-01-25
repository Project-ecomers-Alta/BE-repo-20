package data

import (
	"BE-REPO-20/features/product"
	"BE-REPO-20/features/user"
	_userData "BE-REPO-20/features/user/data"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	UserID        uint
	Name          string `gorm:"default:null"`
	Price         uint   `gorm:"default:null"`
	Quantity      uint   `gorm:"default:null"`
	Description   string `gorm:"default:null"`
	Category      string `gorm:"default:null"`
	User          _userData.User
	ProductImages []ProductImage
}

type ProductImage struct {
	gorm.Model
	ProductID uint
	Url       string
	PublicID  string
}

func (u Product) ModelToCore() product.ProductCore {
	return product.ProductCore{
		ID:          u.ID,
		UserID:      u.UserID,
		Name:        u.Name,
		Price:       u.Price,
		Quantity:    u.Quantity,
		Description: u.Description,
		Category:    u.Category,
		User: user.UserCore{
			ID:          u.User.ID,
			UserName:    u.User.UserName,
			ShopName:    u.User.ShopName,
			Email:       u.User.Email,
			PhoneNumber: u.User.PhoneNumber,
			Domicile:    u.User.Domicile,
			Address:     u.User.Address,
			Image:       u.User.Image,
			Province:    u.User.Province,
			City:        u.User.City,
			Subdistrict: u.User.Subdistrict,
			Tagline:     u.User.TagLine,
			ShopImage:   u.User.ShopImage,
		},
		ProductImages: ProductImageGormToCore(u.ProductImages),
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

func (u ProductImage) ModelToCoreImage() product.ProductImageCore {
	return product.ProductImageCore{
		ID:        u.ID,
		ProductID: u.ProductID,
		Url:       u.Url,
		PublicID:  u.PublicID,
	}
}

func CoreToModelImage(p product.ProductImageCore) ProductImage {
	return ProductImage{
		ProductID: p.ProductID,
		Url:       p.Url,
		PublicID:  p.PublicID,
	}
}

func ProductImageGormToCore(p []ProductImage) []product.ProductImageCore {
	var results []product.ProductImageCore
	for _, v := range p {
		results = append(results, product.ProductImageCore{
			ID:        v.ID,
			ProductID: v.ProductID,
			Url:       v.Url,
			PublicID:  v.PublicID,
		})
	}
	return results
}
