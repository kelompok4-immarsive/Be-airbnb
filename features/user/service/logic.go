package service

import (
	"errors"
	"fajar/testing/features/user"

	"github.com/go-playground/validator/v10"
)

type ServiceUser struct {
	RepoUser user.RepositoryEntities
	validasi *validator.Validate
}

func NewUser(repo user.RepositoryEntities) user.ServiceEntities {
	return &ServiceUser{
		RepoUser: repo,
		validasi: validator.New(),
	}
}

// DeleteId implements user.ServiceEntities
func (service *ServiceUser) DeleteId(id int) (user.CoreUser, error) {
	data, err := service.RepoUser.DeleteId(id) // memanggil struct entities repository yang ada di entities yang berisi coding logic
	return data, err
}

// GetId implements user.ServiceEntities
func (service *ServiceUser) GetById(id int) (data user.CoreUser, err error) {
	data, err = service.RepoUser.GetById(id) // memanggil struct entities repository yang ada di entities yang berisi coding logic
	return
}

// Getall implements user.ServiceEntities
func (service *ServiceUser) Getall() (data []user.CoreUser, err error) {
	data, err = service.RepoUser.Getall() // memanggil struct entities repository yang ada di entities yang berisi coding logic
	return
}

// Register implements user.ServiceEntities
func (service *ServiceUser) Register(input user.CoreUser) (err error) {
	if validasieror := service.validasi.Struct(input); validasieror != nil {
		return validasieror
	}
	_, RegisErr := service.RepoUser.Register(input)
	if RegisErr != nil {
		return errors.New("gagal melakukan register")
	}
	return nil
}

// Update implements user.ServiceEntities
func (service *ServiceUser) Update(id int, input user.CoreUser) error {
	input.Password = user.Bcript(input.Password)
	errUpdate := service.RepoUser.Update(id, input)
	if errUpdate != nil {
		return errors.New("GAGAL mengupdate data , QUERY ERROR")
	}

	return nil
}
