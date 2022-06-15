package repository

import "FinalProject/entity"

type SiswaRepository interface {
	Login(email string, password string) (*entity.Siswa, error)
}