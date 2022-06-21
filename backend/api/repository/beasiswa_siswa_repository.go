package repository

import "FinalProject/entity"

type BeasiswaSiswaRepostiroy interface {
	IsBeasiswaSiswaExistsById(id int) (bool, error)
	UpdateStatusBeasiswa(a entity.BeasiswaSiswaStatusUpdate, id int) (*entity.BeasiswaSiswa, error)
	ApplyBeasiswa(beasiswaSiswa *entity.BeasiswaSiswa) (*entity.BeasiswaSiswa, error)
}
