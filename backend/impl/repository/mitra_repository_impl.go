package repository

import (
	"FinalProject/entity"
	"FinalProject/utility"
	"database/sql"
)

type mitraRepositoryImpl struct {
	db *sql.DB
}

func NewMitraRepositoryImpl(db *sql.DB) *mitraRepositoryImpl {
	return &mitraRepositoryImpl{
		db: db,
	}
}

func (m *mitraRepositoryImpl) Login(username string, password string) (*entity.Mitra, error) {
	query := `
	SELECT
		id, email, password, kategori_user
	FROM
		fp_user
	WHERE
		email = ? AND password = ? AND kategori_user = "MITRA"
	`

	mitra := entity.Mitra{}

	row := m.db.QueryRow(query, username, password)
	if err := row.Scan(
		&mitra.Id,
		&mitra.Email,
		&mitra.Password,
		&mitra.KategoriUser,
	); err != nil {
		return nil, err
	}

	if mitra == (entity.Mitra{}) {
		return nil, utility.ErrNoDataFound
	}

	return &mitra, nil
}