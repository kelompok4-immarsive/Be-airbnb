package komentar

import "time"

type CoreKomentar struct {
	ID           uint
	Vote_star    uint
	UserID       uint
	RoomID       uint
	Isi_komentar string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type ServiceInterface interface {
	Create(input CoreKomentar) (err error)
	GetAll() (data []CoreKomentar, err error)
}
type RepositoryInterface interface {
	Create(input CoreKomentar) (row int, err error)
	GetAll() (data []CoreKomentar, err error)
}
