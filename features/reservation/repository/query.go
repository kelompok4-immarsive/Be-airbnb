package repository

import (
	"fajar/testing/features/reservation"
	"fajar/testing/features/reservation/service"
	"fajar/testing/features/room/repository"

	"gorm.io/gorm"
)

type ReservationRepository struct {
	db *gorm.DB
}

func NewReservation(db *gorm.DB) reservation.RepositoryInterface {
	return &ReservationRepository{
		db: db,
	}
}

// CreateReservasi implements reservasi.RepositoryInterface
func (repo *ReservationRepository) CreateReservasi(input reservation.ReservationCore) error {
	model := repository.Room{}
	tx := repo.db.First(&model, input.RoomID)
	if tx.Error != nil {
		return tx.Error
	}
	input.Price = model.Price
	input.Total_Price = service.WaktuNginap(input.Start_date, input.End_date) * float64(model.Price)

	reservasi := CoretoModel(input) //dari gorm model ke user core yang ada di entities

	tx1 := repo.db.Create(&reservasi) // proses insert data

	if tx1.Error != nil {
		return tx.Error
	}
	return nil
}
