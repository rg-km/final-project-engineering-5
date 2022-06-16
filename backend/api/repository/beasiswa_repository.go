package repository

import "FinalProject/entity"

type BeasiswaRepository interface {
	GetBeasiswaById(id string) ([]*entity.Beasiswa, error)
}
