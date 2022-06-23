package repository

import "FinalProject/entity"

type BeasiswaRepository interface {
	GetBeasiswaById(id int) (*entity.Beasiswa, error)
	IsBeasiswaExistsById(id int) (bool, error)
	GetTotalBeasiswa(nama string) (int, error)
	GetListBeasiswa(page int, limit int, nama string) ([]*entity.Beasiswa, error)
	CreateBeasiswa(beasiswa *entity.Beasiswa) (*entity.Beasiswa, error)
	UpdateBeasiswa(beasiswa entity.Beasiswa, id int) (*entity.Beasiswa, error)
	DeleteBeasiswa(id int) error
}
