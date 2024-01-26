package web

type MidtransRequest struct {
	UserId      int    `json:"user_id" binding:"required"`
	Amount      int64  `json:"amount" binding:"required"`
	OrderID     uint   `json:"order_id" binding:"required"`
	ItemName    string `json:"item_name" binding:"required"`
	FName       string `json:"first_name,omitempty"`
	LName       string `json:"last_name,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Address     string `json:"address,omitempty"`
	City        string `json:"city,omitempty"`
	Postcode    string `json:"postal_code,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	Email       string `json:"email,omitempty"`
}
