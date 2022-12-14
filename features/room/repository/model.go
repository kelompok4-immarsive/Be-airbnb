package repository

import (
	"fajar/testing/features/room"
	"fajar/testing/features/user"
	"fajar/testing/features/user/repository"

	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Name_room string
	Price     string
	Deskripsi string
	Status    string
	Longitude float64
	Latitude  float64
	UserID    uint
	User      repository.User
}

func CoretoModel(dataCore room.RoomCore) Room {
	classGorm := Room{

		Name_room: dataCore.Name_room,
		Price:     dataCore.Price,
		Deskripsi: dataCore.Deskripsi,
		Status:    dataCore.Status,
		Longitude: dataCore.Longitude,
		Latitude:  dataCore.Latitude,
		UserID:    dataCore.UserID,
	}
	return classGorm
}

func (dataModel Room) ModeltoCore() room.RoomCore {
	return room.RoomCore{
		Name_room: dataModel.Name_room,
		Price:     dataModel.Price,
		Deskripsi: dataModel.Deskripsi,
		Status:    dataModel.Status,
		Longitude: dataModel.Longitude,
		Latitude:  dataModel.Latitude,
		UserID:    dataModel.UserID,
		User:      LoadUserModeltoCore(dataModel.User),
	}

}

func LoadUserModeltoCore(model repository.User) user.CoreUser {

	return user.CoreUser{
		Name:    model.Name,
		Email:   model.Email,
		Address: model.Address,
	}

}
