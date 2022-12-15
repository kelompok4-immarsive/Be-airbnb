package repository

import (
	_user "fajar/testing/features/user"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string `gorm:"type:varchar(50)"`
	Password    string `gorm:"type:varchar(500)"`
	Email       string
	Phone       string
	Address     string
	Status      string
	Role        string
	Rooms       []Room
	Komentar    []Komentar
	Reservation []Reservation
}

type Room struct {
	gorm.Model
	Name_room string
	Price     float64
	Deskripsi string
	Status    string
	Longitude float64
	Latitude  float64
	UserID    uint
}

type Komentar struct {
	gorm.Model
	Vote_star    uint
	UserID       uint
	RoomID       uint
	Isi_komentar string
	// User         User
	// Room         Room
}
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
}

func FromCoreUser(CoreData _user.CoreUser) User {
	ModelUser := User{
		Name:     CoreData.Name,
		Password: CoreData.Password,
		Email:    CoreData.Email,
		Phone:    CoreData.Phone,
		Address:  CoreData.Address,
		Status:   CoreData.Status,
		Role:     CoreData.Role,
	}
	return ModelUser
}
func (dataModel *User) ModelsToCore() _user.CoreUser { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	return _user.CoreUser{
		ID:       dataModel.ID,
		Name:     dataModel.Name,
		Email:    dataModel.Email, //mapping data core ke data gorm model
		Password: dataModel.Password,
		Phone:    dataModel.Phone,
		Address:  dataModel.Address,
		Status:   dataModel.Status,
		Role:     dataModel.Role,
	}
}

func ListModelTOCore(dataModel []User) []_user.CoreUser { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	var dataCore []_user.CoreUser
	for _, value := range dataModel {
		dataCore = append(dataCore, value.ModelsToCore())
	}
	return dataCore //  untuk menampilkan data ke controller
}
