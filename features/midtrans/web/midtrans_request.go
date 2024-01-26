package web

type MidtransRequest struct {
	UserId   int    `json:"user_id" binding:"required"`
	Amount   int64  `json:"amount" binding:"required"`
	OrderID  string `json:"order_id" binding:"required"`
	ItemName string `json:"item_name" binding:"required"`
}
