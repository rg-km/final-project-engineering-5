package repository

import "FinalProject/entity"

type SiswaRepository interface {
	Login(email string, password string) (*entity.Siswa, error)
	GetTotalSiswa() (int, error)
	GetListSiswa(page int, limit int, nama string) ([]*entity.SiswaDetail, error)
}