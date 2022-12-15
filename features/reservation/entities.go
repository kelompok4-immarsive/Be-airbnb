package reservation

import (
	"time"
)

type ReservationCore struct {
	ID          uint
	Price       float64
	Total_Price float64
	Start_date  time.Time
	End_date    time.Time
	PayID       uint
	Status      string
	RoomID      uint
	UserID      uint
}

type ServiceInterface interface {
	CreateReservasi(input ReservationCore) error
	// GetAll() (data []RoomCore, err error)
	// UpdateRoom(id int, input RoomCore) error
	// DeleteRoom(id int) (RoomCore, error)
	// GetUserRoom(id uint) (RoomCore, error)
}

type RepositoryInterface interface {
	CreateReservasi(input ReservationCore) error
	// GetAll() (data []RoomCore, err error)
	// UpdateRoom(id int, input RoomCore) error
	// DeleteRoom(id int) (RoomCore, error)
	// GetUserRoom(id uint) (RoomCore, error)
}
