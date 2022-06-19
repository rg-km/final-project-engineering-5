package repository

import "FinalProject/entity"

type BeasiswaSiswaRepostiroy interface {
	UpdateStatusBeasiswa(id int, siswaId int, mitraId int) (*entity.BeasiswaSiswa, error)
}