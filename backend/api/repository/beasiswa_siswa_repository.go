package repository

import "FinalProject/entity"

type BeasiswaSiswaRepostiroy interface {
	IsBeasiswaSiswaExistsById(id int) (bool, error)
	GetTotalBeasiswaSiswa(nama string) (int, error)
	UpdateStatusBeasiswa(a entity.BeasiswaSiswaStatusUpdate, id int) (*entity.BeasiswaSiswa, error)
	ApplyBeasiswa(idSiswa int, idBeasiswa int) (*entity.BeasiswaSiswa, error)
	GetBeasiswaSiswaById(id int) (*entity.BeasiswaSiswa, error)
	GetListBeasiswaSiswaByIdSiswa(id int) ([]*entity.BeasiswaSiswa, error)
	GetListBeasiswaSiswaByIdBeasiswa(id int) ([]*entity.BeasiswaSiswa, error)
	GetListBeasiswaSiswaByIdMitra(idMitra int, page int, limit int, nama string) ([]*entity.BeasiswaSiswa, error)
}
