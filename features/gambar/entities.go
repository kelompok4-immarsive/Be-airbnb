package gambar

import (
	"time"
)

type GambarCore struct {
	ID        uint
	Image_url string `validate:"required"`
	RoomID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ServiceInterface interface {
	CreateGambar(input GambarCore) error
}

type RepositoryInterface interface {
	CreateGambar(input GambarCore) error
}
