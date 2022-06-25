package repository

import "FinalProject/entity"

type SiswaRepository interface {
	IsSiswaExistsByEmail(email string) (bool, error)
	IsSiswaExistsById(id int) (bool, error)
	GetSiswaById(id int) (*entity.SiswaDetail, error)
	Login(email string, password string) (*entity.Siswa, error)
	GetTotalSiswa(nama string) (int, error)
	GetListSiswa(page int, limit int, nama string) ([]*entity.SiswaDetail, error)
	RegisterSiswa(siswa *entity.SiswaDetail, user *entity.Siswa) (*entity.SiswaDetail, error)
}
