package repository

import (
	"errors"
	"fajar/testing/features/komentar"
	"fmt"

	"gorm.io/gorm"
)

type komentarRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) komentar.RepositoryInterface {
	return &komentarRepository{
		db: db,
	}
}

func (repo *komentarRepository) Create(input komentar.CoreKomentar) (row int, err error) {
	komentarGorm := fromCore(input)
	tx := repo.db.Create(&komentarGorm) // proses insert data
	fmt.Print(tx.Error)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failed")
	}
	return int(tx.RowsAffected), nil
}

func (repo *komentarRepository) GetAll() (data []komentar.CoreKomentar, err error) {
	var komentars []Komentar

	tx := repo.db.Find(&komentars)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(komentars)
	return dataCore, nil
}
