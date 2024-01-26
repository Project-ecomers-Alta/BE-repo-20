package midtrans

import (
	"github.com/midtrans/midtrans-go"
)

type MidtransResponse struct {
	Token       string `json:"token"`
	RedirectUrl string `json:"redirect_url"`
}

type MidtransRequest struct {
	UserId   int    `json:"user_id" binding:"required"`
	Amount   int64  `json:"amount" binding:"required"`
	ItemID   string `json:"item_id" binding:"required"`
	ItemName string `json:"item_name" binding:"required"`
}

type MidtransInterface interface {
	CreatePayment()
}

type payment struct{}

func NewMidtransService() MidtransInterface {
	return &payment{}
}

// CreatePayment implements MidtransInterface.
func (p *payment) CreatePayment() {
	midtransServerKey := "SB-Mid-server-XGdfn0Un08oUioFXuEZuTkb-"
	// midtransEnvironment := midtrans.Sandbox

	//Initiate client for Midtrans CoreAPI
	// var c = coreapi.Client
	c.New(midtransServerKey, midtrans.Sandbox)

}
