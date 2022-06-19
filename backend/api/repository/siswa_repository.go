package repository

import "FinalProject/entity"

type SiswaRepository interface {
	Login(email string, password string) (*entity.Siswa, error)
	GetTotalSiswa(nama string) (int, error)
	GetListSiswa(page int, limit int, nama string) ([]*entity.SiswaDetail, error)
	RegisterSiswa(siswa *entity.SiswaDetail, user *entity.Siswa) (*entity.SiswaDetail, error)
}
