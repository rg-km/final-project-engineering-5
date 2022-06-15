package repository

import (
	"FinalProject/entity"
	"database/sql"
	"errors"
)

type siswaLogin struct {

}

type siswaRepositoryImpl struct {
	db *sql.DB
}

func NewSiswaRepositoryImpl(db *sql.DB) *siswaRepositoryImpl {
	return &siswaRepositoryImpl{
		db: db,
	}
}

func (s *siswaRepositoryImpl) Login(username string, password string) (*entity.Siswa,  error) {
	query := `
	SELECT
		id, email, password, kategori_user
	FROM
		fp_user
	WHERE
		email = ? AND password = ? AND kategori_user = "SISWA"
	`

	siswa := entity.Siswa{}

	row := s.db.QueryRow(query, username, password)
	if err := row.Scan(
		&siswa.Id,
		&siswa.Email,
		&siswa.Password,
		&siswa.KategoriUser,
	); err != nil {
		return nil, err
	}

	// siswa tidak ada
	if siswa == (entity.Siswa{}) {
		return nil, errors.New("Tidak ada data.")
	}

	return &siswa, nil
}