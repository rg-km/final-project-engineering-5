package repository

import (
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (s *Repository) Login(username string, password string) (*string, error) {
	querySiswa := "SELECT * FROM siswa WHERE email = ? AND password = ?"
	rowSiswa := s.db.QueryRow(querySiswa, username, password)
	siswa := Siswa{}
	queryMitra := "SELECT * FROM mitra WHERE email = ? AND password = ?"
	rowMitra := s.db.QueryRow(queryMitra, username, password)
	mitra := Mitra{}

	errSiswa := rowSiswa.Scan(&siswa.Id, &siswa.Nama, &siswa.NamaInstansi, &siswa.TingkatPendidikan, &siswa.Alamat, &siswa.NoHp, &siswa.Email, &siswa.Password, &siswa.TanggalLahir, &siswa.Rekening)
	errMitra := rowMitra.Scan(&mitra.Id, &mitra.Nama, &mitra.Profile, &mitra.Pic, &mitra.Situs, &mitra.Email, &mitra.Password)
	var role string
	if errSiswa != nil && errMitra != nil {
		return nil, errSiswa
	}
	if errSiswa == nil {
		role = "siswa"
		return &role, nil
	}
	if errMitra == nil {
		role = "mitra"
		return &role, nil
	}

	return nil, nil

	// siswa := Siswa{}
	// err := row.Scan(&siswa.Id, &siswa.Nama, &siswa.NamaInstansi, &siswa.TingkatPendidikan, &siswa.Alamat, &siswa.NoHp, &siswa.Email, &siswa.Password, &siswa.TanggalLahir, &siswa.Rekening)
	// if err != nil {
	// 	return nil, err
	// }

}
