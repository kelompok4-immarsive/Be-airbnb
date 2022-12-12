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
