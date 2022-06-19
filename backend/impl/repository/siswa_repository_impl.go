package repository

import (
	"FinalProject/entity"
	"FinalProject/utility"
	"database/sql"
	"fmt"
	"strings"
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
		return nil, utility.ErrNoDataFound
	}

	return &siswa, nil
}

func (s *siswaRepositoryImpl) GetTotalSiswa(nama string) (int, error) {
	count := 0
	
	query := `
	SELECT
		COUNT(id)
	FROM
		fp_siswa
	`
	if len(strings.Trim(nama, " ")) != 0 {
		query = fmt.Sprintf(`
		SELECT
			COUNT(id)
		FROM
			fp_siswa
		WHERE nama LIKE "%s%s%s"`, "%", nama, "%s")
	}

	row := s.db.QueryRow(query)
	if err := row.Scan(
		&count,
	); err != nil {
		return -1, err
	}

	return count, nil
}

func (s *siswaRepositoryImpl) GetListSiswa(page int, limit int, nama string) ([]*entity.SiswaDetail, error) {
	offset := limit * (page-1)
	
	query := `
	SELECT
		id, id_user, nama, tanggal_lahir, nomor_telepon, nama_instansi, tingkat_pendidikan, nomor_rekening, nama_bank, alamat
	FROM
		fp_siswa
	LIMIT ?
	OFFSET ?
	`
	if len(strings.Trim(nama, " ")) != 0 {
		query = fmt.Sprintf(`
		SELECT
			id, id_user, nama, tanggal_lahir, nomor_telepon, nama_instansi, tingkat_pendidikan, nomor_rekening, nama_bank, alamat
		FROM
			(
				SELECT
					id, id_user, nama, tanggal_lahir, nomor_telepon, nama_instansi, tingkat_pendidikan, nomor_rekening, nama_bank, alamat
				FROM
					fp_siswa WHERE nama LIKE "%s%s%s"
			) AS fp_siswa
		LIMIT ?
		OFFSET ?
		`,"%",nama,"%")
	}

	rows, err := s.db.Query(query, limit, offset)
	if err != nil {
		return []*entity.SiswaDetail{}, err
	}
	defer rows.Close()

	listSiswa := make([]*entity.SiswaDetail, 0)
	for rows.Next() {
		row := &entity.SiswaDetail{}
		if err := rows.Scan(
			&row.Id,
			&row.IdUser,
			&row.Nama,
			&row.TanggalLahir,
			&row.NomorTelepon,
			&row.NamaInstansi,
			&row.TingkatPendidikan,
			&row.NomorRekening,
			&row.NamaBank,
			&row.Alamat,
		); err != nil {
			return []*entity.SiswaDetail{}, err
		}

		listSiswa = append(listSiswa, row)
	}

	return listSiswa, nil
}
