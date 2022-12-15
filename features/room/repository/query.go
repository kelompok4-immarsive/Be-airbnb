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

// DeleteMentee implements room.RepositoryInterface

// GetAll implements room.RepositoryInterface

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
func (repo *RoomRepository) GetAll() (data []room.RoomCore, err error) {
	var rooms []Room //mengambil data gorm model(model.go)
	tx := repo.db.Unscoped().Find(&rooms)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var DataCore = ListModelTOCore(rooms) //mengambil data dari gorm model(file repository(model.go))

	return DataCore, nil

}

// UpdateRoom implements room.RepositoryInterface
func (repo *RoomRepository) UpdateRoom(id int, input room.RoomCore) error {
	room := CoretoModel(input)
	tx := repo.db.Model(room).Where("id = ?", id).Updates(room)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}
func (repo *RoomRepository) DeleteRoom(id int) (room.RoomCore, error) {
	rooms := Room{}
	tx := repo.db.Delete(&rooms, id)
	if tx.Error != nil {
		return room.RoomCore{}, tx.Error
	}

	tx1 := repo.db.Unscoped().Where("id=?", id).Find(&rooms)
	if tx1.Error != nil {
		return room.RoomCore{}, tx1.Error
	}
	if tx.RowsAffected == 0 {
		return room.RoomCore{}, errors.New("id not found")

	}
	result := rooms.ModeltoCore()
	return result, nil
}

// GetUserRoom implements room.RepositoryInterface
func (repo *RoomRepository) GetUserRoom(id uint) (room.RoomCore, error) {
	Rooms := Room{}

	tx := repo.db.Preload("User").First(&Rooms, id)
	fmt.Println("rooms", tx.Error)
	if tx.Error != nil {
		return room.RoomCore{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return room.RoomCore{}, errors.New("id not found")

	}
	result := Rooms.ModeltoCore()
	return result, nil
}
