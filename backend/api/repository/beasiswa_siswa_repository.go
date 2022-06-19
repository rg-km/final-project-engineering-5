package repository

import "FinalProject/entity"

type BeasiswaSiswaRepostiroy interface {
	UpdateStatusBeasiswa(a entity.BeasiswaSiswaStatusUpdate, id int) (*entity.BeasiswaSiswa, error)
}