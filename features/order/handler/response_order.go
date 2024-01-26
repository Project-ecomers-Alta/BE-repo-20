package handler

import (
	"BE-REPO-20/features/midtrans/web"
	"BE-REPO-20/features/order"
	"BE-REPO-20/features/user"
)

type OrderResponse struct {
	Id         uint                `json:"id"`
	UserID     uint                `json:"user_id"`
	Address    string              `json:"address"`
	CreditCard uint                `json:"credit_card"`
	Status     string              `json:"status"`
	Invoice    string              `json:"invoice"`
	Total      uint                `json:"total"`
	VirtualAcc uint                `json:"virtual_acc"`
	User       user.UserCore       `json:"user"`
	ItemOrders []ItemOrderResponse `json:"item_order"`
}

type ItemOrderResponse struct {
	ID           uint   `json:"id"`
	OrderID      uint   `json:"order_id"`
	ProductName  string `json:"product_name"`
	ProductPrice uint   `json:"product_price"`
	Quantity     uint   `json:"quantity"`
	SubTotal     uint   `json:"subtotal"`
}

func CoreToResponse(o order.OrderCore) OrderResponse {
	return OrderResponse{
		Id:         o.Id,
		UserID:     o.UserID,
		Address:    o.Address,
		CreditCard: o.CreditCard,
		Status:     o.Status,
		Invoice:    o.Invoice,
		Total:      o.Total,
		VirtualAcc: o.VirtualAcc,
		User: user.UserCore{
			ID:          o.User.ID,
			UserName:    o.User.UserName,
			ShopName:    o.User.ShopName,
			Email:       o.User.Email,
			PhoneNumber: o.User.PhoneNumber,
			Domicile:    o.User.Domicile,
			Address:     o.User.Address,
			Image:       o.User.Image,
			Province:    o.User.Province,
			City:        o.User.City,
			Subdistrict: o.User.Subdistrict,
			Tagline:     o.User.Tagline,
			ShopImage:   o.User.ShopImage,
		},
		ItemOrders: ItemOrderResponseToList(o.ItemOrders),
	}
}

func ItemOrderResponseToList(o []order.ItemOrderCore) []ItemOrderResponse {
	var results []ItemOrderResponse
	for _, v := range o {
		results = append(results, ItemOrderResponse{
			ID:           v.ID,
			OrderID:      v.OrderID,
			ProductName:  v.ProductName,
			ProductPrice: v.ProductPrice,
			Quantity:     v.Quantity,
			SubTotal:     v.SubTotal,
		})
	}
	return results
}

func OrderToMidtrans(o OrderResponse) web.MidtransRequest {
	return web.MidtransRequest{
		UserId:      int(o.User.ID),
		Amount:      int64(TotalAmount(o)),
		OrderID:     o.Id,
		ItemName:    "Kaos",
		FName:       o.User.UserName,
		LName:       o.User.ShopName,
		Phone:       o.User.PhoneNumber,
		Address:     o.Address,
		City:        o.User.City,
		Postcode:    "13660",
		CountryCode: "IDN",
		Email:       o.User.Email,
	}
}

func TotalAmount(o OrderResponse) int {
	var amount int
	for _, v := range o.ItemOrders {
		amount += int(v.SubTotal)
	}
	return amount
}
