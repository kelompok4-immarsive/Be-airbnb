package repository

import (
	"errors"
	"fajar/testing/features/room"
	_userRepo "fajar/testing/features/user/repository"
	"fmt"

	"gorm.io/gorm"
)

type RoomRepository struct {
	db *gorm.DB
}

func NewRoom(db *gorm.DB) room.RepositoryInterface {
	return &RoomRepository{
		db: db,
	}
}

// CreateRoom implements room.RepositoryInterface
func (repo *RoomRepository) CreateRoom(input room.RoomCore) error {

	var user _userRepo.User
	fmt.Println(" input", input)
	RoomGorm := CoretoModel(input)
	fmt.Println("RoomGorm", RoomGorm)
	tx := repo.db.Create(&RoomGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}

	tx2 := repo.db.Model(&user).Where("id = ?", RoomGorm.UserID).Update("Role", "Hosting")
	if tx2.Error != nil {
		return tx2.Error
	}
	return nil
}
