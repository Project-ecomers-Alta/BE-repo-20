package handler

type AuthRespon struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Domicile string `json:"domicile"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}
