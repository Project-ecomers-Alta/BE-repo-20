package data

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName    string `gorm:"default:null"`
	FullName    string `gorm:"default:null"`
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

// struct user gorm model
type Order struct {
	gorm.Model
	UserID  uint   `gorm:"foreignKey" json:"user_id" form:"user_id"`
	Address string `json:"addres" form:"addres"`
	// CreditCard string `json:"credit_card" form:"credit_card"`
	PaymentMethod   string `json:"payment_method" form:"payment_method"`
	TransactionTime string `json:"transaction_time" form:"transaction_time"`
	Status          string `json:"status" form:"status"`
	Invoice         string `json:"invoice" form:"invoice"`
	Total           uint   `json:"total" form:"total"`
	VirtualAcc      string `json:"virtual_acc" form:"virtual_acc"`
	User            User   `gorm:"foreignKey:UserID"`
}

// struct project gorm model
type ItemOrder struct {
	gorm.Model
	OrderID      uint
	ProductName  string `json:"product_name" form:"product_name"`
	ProductPrice uint   `json:"product_price" form:"product_price"`
	Quantity     uint   `json:"quantity" form:"quantity"`
	SubTotal     uint   `json:"sub_total" form:"sub_total"`
	Order        Order
}
