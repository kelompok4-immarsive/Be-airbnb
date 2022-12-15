package service

import (
	"errors"
	"fajar/testing/features/reservation"
	"time"

	"github.com/go-playground/validator"
)

type ReservationService struct {
	reservasiRepository reservation.RepositoryInterface
	validasi            *validator.Validate
}

func NewReservation(repo reservation.RepositoryInterface) reservation.ServiceInterface {
	return &ReservationService{
		reservasiRepository: repo,
		validasi:            validator.New(),
	}
}

// CreateReservasi implements reservasi.ServiceInterface
func (service *ReservationService) CreateReservasi(input reservation.ReservationCore) error {
	err := service.reservasiRepository.CreateReservasi(input)

	if err != nil {
		return errors.New("failed to add reservasi")
	}
	return nil
}
func WaktuNginap(start, end time.Time) float64 {
	difference := start.Sub(end)
	days := float64(difference.Hours() / 24)

	return days

}
