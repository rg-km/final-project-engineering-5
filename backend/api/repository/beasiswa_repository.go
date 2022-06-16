package repository

import "FinalProject/entity"

type BeasiswaRepository interface {
	GetTotalBeasiswa(nama string) (int, error)
	GetListBeasiswa(page int, limit int, nama string) ([]*entity.Beasiwa, error)
}