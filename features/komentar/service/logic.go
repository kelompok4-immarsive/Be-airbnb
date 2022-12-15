package service

import (
	"errors"
	"fajar/testing/features/komentar"

	"github.com/go-playground/validator"
)

type komentarService struct {
	komentarRepository komentar.RepositoryInterface
	validate           *validator.Validate
}

func New(repo komentar.RepositoryInterface) komentar.ServiceInterface {
	return &komentarService{
		komentarRepository: repo,
		validate:           validator.New(),
	}
}

func (service *komentarService) Create(input komentar.CoreKomentar) (err error) {
	// if errValidate := service.validate.Struct(input); errValidate != nil {
	// return errValidate
	// }
	_, errCreate := service.komentarRepository.Create(input)
	if errCreate != nil {
		return errors.New("failed to insert data, error query")
	}
	return nil
}

func (service *komentarService) GetAll() (data []komentar.CoreKomentar, err error) {
	data, err = service.komentarRepository.GetAll()
	return

}
