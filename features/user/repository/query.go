package repository

import (
	"errors"
	"fajar/testing/features/user"

	"gorm.io/gorm"
)

type RepoUser struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) user.RepositoryEntities {
	return &RepoUser{
		db: db,
	}
}

// DeleteId implements user.RepositoryEntities
func (*RepoUser) DeleteId(id int) (user.CoreUser, error) {
	panic("unimplemented")
}

// GetId implements user.RepositoryEntities
func (*RepoUser) GetId(id int) (CoreUser error, err error) {
	panic("unimplemented")
}

// Getall implements user.RepositoryEntities
func (*RepoUser) Getall() (data []user.CoreUser, err error) {
	panic("unimplemented")
}

// Register implements user.RepositoryEntities
func (repo *RepoUser) Register(input user.CoreUser) (rows int, err error) {
	userModel := FromCoreUser(input)

	tx := repo.db.Create(&userModel)

	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("register gagal")
	}
	return int(tx.RowsAffected), nil

}

// Update implements user.RepositoryEntities
func (*RepoUser) Update(id int, input user.CoreUser) error {
	panic("unimplemented")
}
