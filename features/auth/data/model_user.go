package data

import (
	"BE-REPO-20/features/auth"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName    string `gorm:"default:null"`
	Shopname    string `gorm:"default:null"`
	Email       string `gorm:"default:null;unique"`
	PhoneNumber string `gorm:"default:null"`
	Domicile    string `gorm:"default:null"`
	Password    string `gorm:"default:null"`
	Address     string `gorm:"default:null"`
	Role        string `gorm:"default:user"`
	Image       string `gorm:"default:null"`
	Province    string `gorm:"default:null"`
	City        string `gorm:"default:null"`
	Subdistrict string `gorm:"default:null"`
	TagLine     string `gorm:"default:null"`
	Category    string `gorm:"default:null"`
}

func (u User) ModelToCore() auth.AuthCore {
	return auth.AuthCore{
		ID:       u.ID,
		UserName: u.UserName,
		Email:    u.Email,
		Domicile: u.Domicile,
		Role:     u.Role,
		Password: u.Password,
	}
}
