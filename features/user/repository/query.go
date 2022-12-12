package repository

import (
	"errors"
	"fajar/testing/features/user"
	"time"

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
func (repo *RepoUser) Getall() (data []user.CoreUser, err error) {
	var users []User //mengambil data gorm model(model.go)
	tx := repo.db.Unscoped().Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var DataCore = ListModelTOCore(users) //mengambil data dari gorm model(file repository(model.go))

	return DataCore, nil

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
func (repo *RepoUser) Update(id int, input user.CoreUser) error {
	userGorm := FromCoreUser(input)

	if userGorm.UpdatedAt != (time.Time{}) {
		userGorm.Status = "Not-Active"
	} else {
		userGorm.Status = "Active"
	}
	tx := repo.db.Model(&userGorm).Where("id = ?", id).Updates(&userGorm)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
