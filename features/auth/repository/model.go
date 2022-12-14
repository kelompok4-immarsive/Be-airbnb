package repository

import (
	"fajar/testing/features/auth"

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

func (data User) ToCore() auth.CoreUser {
	return auth.CoreUser{
		Name:     data.Name,
		Password: data.Password,
		Email:    data.Email,
		Phone:    data.Phone,
		Address:  data.Address,
		Role:     data.Role,
		Status:   data.Status,
	}
}
