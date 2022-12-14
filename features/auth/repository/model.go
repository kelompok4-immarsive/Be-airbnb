package repository

import (
	"fajar/testing/features/auth"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint
	Name      string `gorm:"type:varchar(50)"`
	Password  string `gorm:"type:varchar(500)"`
	Email     string
	Phone     string
	Address   string
	Status    string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (data User) ToCore() auth.CoreUser {
	return auth.CoreUser{
		ID:        data.ID,
		Name:      data.Name,
		Password:  data.Password,
		Email:     data.Email,
		Phone:     data.Phone,
		Address:   data.Address,
		Role:      data.Role,
		Status:    data.Status,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}
