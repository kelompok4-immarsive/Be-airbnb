package repository

import (
	"fajar/testing/features/gambar"
	"fajar/testing/features/user/repository"

	"gorm.io/gorm"
)

type Gambar struct {
	gorm.Model
	Image_url string
	RoomID    string
	Room      repository.Room
}

func CoretoModel(dataCore gambar.GambarCore) Gambar {
	classGorm := Gambar{

		Image_url: dataCore.Image_url,
		RoomID:    dataCore.RoomID,
	}
	return classGorm
}

func (dataModel Gambar) ModeltoCore() gambar.GambarCore {
	return gambar.GambarCore{
		Image_url: dataModel.Image_url,
		RoomID:    dataModel.RoomID,
	}

}

// func LoadUserModeltoCore(model repository.Room) Room.CoreUser {

// 	return Room.CoreUser{
// 		Name_room: model.Name_room,
// 		Price:     model.Price,
// 		Deskripsi:   model.Deskripsi,
// 	}

// }
