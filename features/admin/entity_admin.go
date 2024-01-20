package admin

import "time"

type AdminUserCore struct {
	ID        uint
	FullName  string
	UserName  string
	Email     string
	Domicile  string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AdminDataInterface interface {
	GetUserRoleById(userId int) (string, error)
	SelectAllUser() ([]AdminUserCore, error)
	SearchUserByQuery(query string) ([]AdminUserCore, error)
}

type AdminServiceInterface interface {
	GetUserRoleById(userId int) (string, error)
	SelectAllUser() ([]AdminUserCore, error)
	SearchUserByQuery(query string) ([]AdminUserCore, error)
}
