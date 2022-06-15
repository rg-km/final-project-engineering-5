package repository

import "FinalProject/entity"

type SiswaRepository interface {
	Login(username string, password string) (*entity.Siswa, error)
}