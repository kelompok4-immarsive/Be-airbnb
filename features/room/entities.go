package room

import (
	"fajar/testing/features/user"
	"time"
)

type RoomCore struct {
	ID        uint
	Name_room string  `validate:"required"`
	Price     float64 `validate:"required"`
	Deskripsi string  `validate:"required"`
	Status    string  `validate:"required"`
	Longitude float64 `validate:"required"`
	Latitude  float64 `validate:"required"`
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
	User      user.CoreUser
}

type ServiceInterface interface {
	CreateRoom(input RoomCore) error
	GetAll() (data []RoomCore, err error)
	UpdateRoom(id int, input RoomCore) error
	DeleteRoom(id int) (RoomCore, error)
	GetUserRoom(id uint) (RoomCore, error)
}

type RepositoryInterface interface {
	CreateRoom(input RoomCore) error
	GetAll() (data []RoomCore, err error)
	UpdateRoom(id int, input RoomCore) error
	DeleteRoom(id int) (RoomCore, error)
	GetUserRoom(id uint) (RoomCore, error)
}
