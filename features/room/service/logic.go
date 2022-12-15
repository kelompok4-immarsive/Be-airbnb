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
func (service *RoomService) GetAll() (data []room.RoomCore, err error) {
	data, err = service.roomRepository.GetAll() // memanggil struct entities repository yang ada di entities yang berisi coding logic
	return
}

// UpdateRoom implements room.ServiceInterface
func (service *RoomService) UpdateRoom(id int, input room.RoomCore) error {
	err := service.roomRepository.UpdateRoom(id, input)
	if err != nil {
		return errors.New("id not found")
	}
	return nil
}

// DeleteMentee implements room.ServiceInterface
func (service *RoomService) DeleteRoom(id int) (room.RoomCore, error) {
	data, err := service.roomRepository.DeleteRoom(id)
	return data, err
}

// GetUserRoom implements room.ServiceInterface
func (service *RoomService) GetUserRoom(id uint) (room.RoomCore, error) {
	data, err := service.roomRepository.GetUserRoom(id)
	return data, err
}
