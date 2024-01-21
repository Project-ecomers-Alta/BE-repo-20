package user

type UserCore struct {
	ID          uint
	UserName    string `validate:"required"`
	Email       string `validate:"required,email"`
	Domicile    string `validate:"required"`
	PhoneNumber string
	Imgage      string
	Tagline     string
	Province    string
	City        string
	Subdistrict string
	Address     string
	Category    string
}

type UserDataInterface interface {
	SelectUser(id int) (*UserCore, error)
	SelectShop(id int) (*UserCore, error)
	UpdateUser(id int, input UserCore) error
	UpdateShop(id int, input UserCore) error
	Delete(id int) error
}

type UserServiceInterface interface {
	SelectUser(id int) (*UserCore, error)
	SelectShop(id int) (*UserCore, error)
	UpdateUser(id int, input UserCore) error
	UpdateShop(id int, input UserCore) error
	Delete(id int) error
}
