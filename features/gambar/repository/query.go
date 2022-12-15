package repository

import (
	"errors"
	"fajar/testing/features/gambar"
	"fmt"

	"gorm.io/gorm"
)

type GambarRepository struct {
	db *gorm.DB
}

func NewGambar(db *gorm.DB) gambar.RepositoryInterface {
	return &GambarRepository{
		db: db,
	}
}

// CreateGambar implements gambar.RepositoryInterface
func (repo *GambarRepository) CreateGambar(input gambar.GambarCore) error {
	RoomGorm := CoretoModel(input)

	tx := repo.db.Create(&RoomGorm)
	fmt.Println("gambar", tx.Error)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}

	return nil
}
