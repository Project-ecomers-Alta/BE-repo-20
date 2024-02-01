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

type AdminOrderCore struct {
	ID            uint `gorm:"primaryKey"`
	UserID        uint `gorm:"column:user_id"`
	Address       string
	PaymentMethod string
	Status        string
	Invoice       string
	Total         uint
	VirtualAcc    string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type AdminItemOrderCore struct {
	ID           uint `gorm:"primaryKey"`
	OrderID      uint `gorm:"column:order_id"`
	ProductName  string
	ProductPrice uint
	Quantity     uint
	SubTotal     uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Order        AdminOrderCore `gorm:"foreignKey:OrderID;tableName:item_orders"`
}

type AdminDataInterface interface {
	GetUserRoleById(userId int) (string, error)
	SelectAllUser() ([]AdminUserCore, error)
	SearchUserByQuery(query string) ([]AdminUserCore, error)
	SelectAllOrder() ([]AdminItemOrderCore, error)
	SearchOrderByQuery(query string) ([]AdminItemOrderCore, error)
}

type AdminServiceInterface interface {
	GetUserRoleById(userId int) (string, error)
	SelectAllUser() ([]AdminUserCore, error)
	SearchUserByQuery(query string) ([]AdminUserCore, error)
	SelectAllOrder() ([]AdminItemOrderCore, error)
	SearchOrderByQuery(query string) ([]AdminItemOrderCore, error)
}
