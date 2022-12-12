package user

import "time"

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

type ServiceEntities interface {
	Getall() (data []CoreUser, err error)
	Register(input CoreUser) (err error)
	Update(id int, input CoreUser) error
	GetId(id int) (CoreUser, err error)
	DeleteId(id int) (CoreUser, error)
}
type RepositoryEntities interface {
	Getall() (data []CoreUser, err error)
	Register(input CoreUser) (rows int, err error)
	Update(id int, input CoreUser) error
	GetId(id int) (CoreUser, err error)
	DeleteId(id int) (CoreUser, error)
}
