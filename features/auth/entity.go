package auth

import (
	"time"
)

type CoreUser struct {
	ID        uint
	Name      string `validate:"required"`
	Password  string `validate:"required"`
	Email     string `validate:"required,email"`
	Phone     string `validate:"required"`
	Address   string `validate:"required"`
	Role      string `validate:"required"`
	Status    string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ServiceInterface interface {
	Login(email, password string) (token string, err error)
}

type RepositoryInterface interface {
	FindUser(email, password string) (token string, err error)
}
