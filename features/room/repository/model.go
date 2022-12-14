package repository

import (
	"fajar/testing/features/room"
	_room "fajar/testing/features/room"
	"fajar/testing/features/user"
	"fajar/testing/features/user/repository"
	"time"

	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Name_room   string
	Price       float64
	Deskripsi   string
	Status      string
	Longitude   float64
	Latitude    float64
	UserID      uint
	Gambar      []Gambar
	User        repository.User
	Reservation []Reservation
}
type Gambar struct {
	gorm.Model
	Image_url string
	RoomID    string
}
type Reservation struct {
	gorm.Model
	Price       float64
	Total_Price float64
	Start_date  time.Time
	End_date    time.Time
	PayID       uint
	Status      string
	RoomID      uint
	UserID      uint
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

func (dataModel *Room) ModelstoCore() room.RoomCore {
	return room.RoomCore{
		Name_room: dataModel.Name_room,
		Price:     dataModel.Price,
		Deskripsi: dataModel.Deskripsi,
		Status:    dataModel.Status,
		Longitude: dataModel.Longitude,
		Latitude:  dataModel.Latitude,
		UserID:    dataModel.UserID,
	}

}
func LoadUserModeltoCore(model repository.User) user.CoreUser {

	return user.CoreUser{
		Name:    model.Name,
		Email:   model.Email,
		Address: model.Address,
	}

}
func ListModelTOCore(dataModel []Room) []_room.RoomCore { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	var dataCore []_room.RoomCore
	for _, value := range dataModel {
		dataCore = append(dataCore, value.ModelstoCore())
	}
	return dataCore //  untuk menampilkan data ke controller
}
