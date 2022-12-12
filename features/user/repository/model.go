package repository

import (
	_user "fajar/testing/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(50)"`
	Password string `gorm:"type:varchar(500)"`
	Email    string
	Phone    string
	Address  string
	Status   string
	Role     string
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
