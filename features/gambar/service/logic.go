package service

import (
	"errors"
	"fajar/testing/features/gambar"

	"github.com/go-playground/validator"
)

type RoomService struct {
	GambarRepository gambar.RepositoryInterface
	validasi         *validator.Validate
}

func NewGambar(repo gambar.RepositoryInterface) gambar.ServiceInterface {
	return &RoomService{
		GambarRepository: repo,
		validasi:         validator.New(),
	}
}

// CreateGambar implements gambar.ServiceInterface
func (service *RoomService) CreateGambar(input gambar.GambarCore) error {
	if validateERR := service.validasi.Struct(input); validateERR != nil {
		return validateERR
	}

	errCreate := service.GambarRepository.CreateGambar(input)
	if errCreate != nil {
		return errors.New("GAGAL MENAMBAH DATA , QUERY ERROR")
	}

	return nil
}

// CreateRoom implements room.ServiceInterface
func (service *RoomService) Getall() (data []gambar.GambarCore, err error) {
	data, err = service.GambarRepository.Getall() // memanggil struct entities repository yang ada di entities yang berisi coding logic
	return
}
