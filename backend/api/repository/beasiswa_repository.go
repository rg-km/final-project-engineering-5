package repository

import "FinalProject/entity"

type BeasiswaRepository interface {
	GetBeasiswaById(id string) ([]*entity.Beasiswa, error)
	GetTotalBeasiswa(nama string) (int, error)
	GetListBeasiswa(page int, limit int, nama string) ([]*entity.Beasiswa, error)
	CreateBeasiswa(beasiswa *entity.Beasiswa) (*entity.Beasiswa, error)
}
