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

func (m *mitraRepositoryImpl) RegisterMitra(mitra *entity.MitraDetail, user *entity.Mitra) (*entity.MitraDetail, error) {
	queryUser := `
	INSERT INTO
		fp_user (email, password, kategori_user)
	VALUES
		(?, ?, "MITRA")
	`
	result, err := m.db.Exec(queryUser, user.Email, user.Password)
	if err != nil {
		return nil, err
	}
	userID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	mitra.IdUser = int(userID)

	query := `
	INSERT INTO fp_mitra (id_user, nama, about, nomor_pic, nama_pic, situs, alamat)
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	_, err = m.db.Exec(query, mitra.IdUser, mitra.Nama, mitra.About, mitra.NomorPic, mitra.NamaPic, mitra.Situs, mitra.Alamat)
	if err != nil {
		return nil, err
	}

	return mitra, nil
}
