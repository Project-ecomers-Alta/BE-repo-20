package auth

type AuthCore struct {
	ID       uint
	UserName string `validate:"required"`
	Email    string `validate:"required,email"`
	Domicile string `validate:"required"`
	Password string `validate:"required"`
	Role     string `gorm:"default:user"`
	Token    string
}

type AuthDataInterface interface {
	Register(input AuthCore) (data *AuthCore, token string, err error)
	Login(email, password string) (data *AuthCore, err error)
}

type AuthServiceInterface interface {
	Register(input AuthCore) (data *AuthCore, token string, err error)
	Login(email, password string) (data *AuthCore, token string, err error)
}
