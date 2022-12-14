package repository

import (
	_komentar "fajar/testing/features/komentar"

	"gorm.io/gorm"
)

type Komentar struct {
	gorm.Model
	Vote_star    uint
	UserID       uint
	RoomID       uint
	Isi_komentar string
	User         User
	Room         Room
}

type User struct {
	gorm.Model
	Name     string
	Password string
	Email    string
	Phone    string
	Address  string
	Status   string
	Role     string
	Komentar []Komentar
}

type Room struct {
	gorm.Model
	Name_room string
	Price     string
	Deskripsi string
	Status    string
	Longitude float64
	Latitude  float64
	Komentar  []Komentar
}

func fromCore(dataCore _komentar.CoreKomentar) Komentar {
	userGorm := Komentar{
		Vote_star:    dataCore.Vote_star,
		UserID:       dataCore.UserID,
		RoomID:       dataCore.RoomID,
		Isi_komentar: dataCore.Isi_komentar,
	}
	return userGorm
}

func (dataModel *Komentar) toCore() _komentar.CoreKomentar {
	return _komentar.CoreKomentar{
		ID:           dataModel.ID,
		Vote_star:    dataModel.Vote_star,
		UserID:       dataModel.UserID,
		RoomID:       dataModel.RoomID,
		Isi_komentar: dataModel.Isi_komentar,
	}

}

func toCoreList(dataModel []Komentar) []_komentar.CoreKomentar {
	var dataCore []_komentar.CoreKomentar
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
