package gambar

import (
	"fajar/testing/features/room"
	"time"
)

type GambarCore struct {
	ID        uint
	Image_url string `validate:"required"`
	RoomID    string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Room      room.RoomCore
}

type ServiceInterface interface {
	CreateGambar(input GambarCore) error
}

type RepositoryInterface interface {
	CreateGambar(input GambarCore) error
}
