package repository

import "FinalProject/entity"

type BeasiswaSiswaRepostiroy interface {
	IsBeasiswaSiswaExistsById(id int) (bool, error)
	UpdateStatusBeasiswa(a entity.BeasiswaSiswaStatusUpdate, id int) (*entity.BeasiswaSiswa, error)
}
