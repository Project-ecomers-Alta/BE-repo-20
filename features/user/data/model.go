package data

import (
	"BE-REPO-20/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName    string `gorm:"default:null"`
	FullName    string `gorm:"default:null"`
	Email       string `gorm:"default:null"`
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

func (u User) ModelToCore() user.UserCore {
	return user.UserCore{
		ID:          u.ID,
		UserName:    u.UserName,
		Email:       u.Email,
		Domicile:    u.Domicile,
		PhoneNumber: u.PhoneNumber,
		Image:       u.Image,
	}
}

func CoreToModel(input user.UserCore) User {
	return User{
		UserName:    input.UserName,
		Email:       input.Email,
		Domicile:    input.Domicile,
		PhoneNumber: input.PhoneNumber,
		Image:       input.Image,
	}
}
