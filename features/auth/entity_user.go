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

type AuthCorePassword struct {
	ID       uint
	Password string
}

type AuthDataInterface interface {
	Register(input AuthCore) (data *AuthCore, token string, err error)
	Login(email, password string) (data *AuthCore, err error)
	UpdatePassword(id uint, input AuthCorePassword) error
	CheckPassword(savedPassword, inputPassword string) bool
}

type AuthServiceInterface interface {
	Register(input AuthCore) (data *AuthCore, token string, err error)
	Login(email, password string) (data *AuthCore, token string, err error)
	UptdatePassword(id uint, input AuthCorePassword) error
}
