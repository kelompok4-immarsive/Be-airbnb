package komentar

import "time"

type CoreKomentar struct {
	ID           uint
	Vote_star    uint
	User         CoreUser
	Room         CoreRoom
	Isi_komentar string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type CoreUser struct {
	ID uint
}
type CoreRoom struct {
	ID uint
}
type ServiceInterface interface {
	Create(input CoreKomentar) (err error)
	GetAll() (data []CoreKomentar, err error)
}
type RepositoryInterface interface {
	Create(input CoreKomentar) (row int, err error)
	GetAll() (data []CoreKomentar, err error)
}
