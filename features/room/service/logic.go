package service

import (
	"errors"
	"fajar/testing/features/room"

	"github.com/go-playground/validator"
)

type RoomService struct {
	roomRepository room.RepositoryInterface
	validasi       *validator.Validate
}

func NewRoom(repo room.RepositoryInterface) room.ServiceInterface {
	return &RoomService{
		roomRepository: repo,
		validasi:       validator.New(),
	}
}

// CreateRoom implements room.ServiceInterface
func (service *RoomService) CreateRoom(input room.RoomCore) error {

	err := service.roomRepository.CreateRoom(input)

	if err != nil {
		return errors.New("failed to add Room")
	}
	return nil
}
