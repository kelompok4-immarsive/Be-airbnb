package repository

import (
	"fajar/testing/features/reservation"
	"fajar/testing/features/user/repository"
	"time"

	"gorm.io/gorm"
)

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
	Room        repository.Room
	User        repository.User
}

func CoretoModel(dataCore reservation.ReservationCore) Reservation {
	classGorm := Reservation{

		Price:       dataCore.Price,
		Total_Price: dataCore.Total_Price,
		Start_date:  dataCore.Start_date,
		Status:      dataCore.Status,
		End_date:    dataCore.End_date,
		PayID:       dataCore.PayID,
		UserID:      dataCore.UserID,
		RoomID:      dataCore.RoomID,
	}
	return classGorm
}
