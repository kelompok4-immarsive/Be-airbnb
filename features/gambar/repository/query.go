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
	fmt.Println("gambar", RoomGorm)
	tx := repo.db.Create(&RoomGorm)

	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}

	return nil
}

// Getall implements user.RepositoryEntities
func (repo *GambarRepository) Getall() (data []gambar.GambarCore, err error) {
	var gambars []Gambar //mengambil data gorm model(model.go)
	tx := repo.db.Unscoped().Find(&gambars)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var DataCore = ListModelTOCore(gambars) //mengambil data dari gorm model(file repository(model.go))

	return DataCore, nil

}
