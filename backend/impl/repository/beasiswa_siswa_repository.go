package repository

import (
	"FinalProject/entity"
	"database/sql"
)

type beasiswaSiswaRepositoryImpl struct {
	db *sql.DB
}

func NewBeasiswaSiswaRepositoryImpl(db *sql.DB) *beasiswaRepositoryImpl {
	return &beasiswaRepositoryImpl{
		db: db,
	}
}

func (b *beasiswaRepositoryImpl) UpdateStatusBeasiswa(beasiswaSiswaStatusUpdate entity.BeasiswaSiswaStatusUpdate) (*entity.BeasiswaSiswa, error) {
	query := `
	UPDATE
		fp_beasisw_siswa
	SET
		status = ?
	WHERE
		id = ? AND id_siswa = ? AND id_beasiswa = ?
	`

	_, err := b.db.Exec(
		beasiswaSiswaStatusUpdate.Status,
		beasiswaSiswaStatusUpdate.Id,
		beasiswaSiswaStatusUpdate.IdSiswa,
		beasiswaSiswaStatusUpdate.IdBeasiswa)
	if err != nil {
		return nil, err
	}

	query = `
	SELECT
		fp_bs.id,
		fp_bs.id_siswa,
		fp_s.nama,
		fp_bs.id_beasiswa,
		fp_b.judul_beasiswa,
		fp_b.id_mitra,
		fp_m.nama,
		fp_bs.status,
		fp_bs.tangga_daftar
	FROM
		fp_beasiswa_siswa fp_bs
	INNER JOIN
		fp_beasiswa fp_b
	ON
		fp_bs.id_beasiswa = fp_b.id
	INNER JOIN
		fp_mitra fp_m
	ON
		fp_b.id_mitra = fp_m.id
	WHERE
		fp_bs.id = ? AND fp_bs.id_siswa = ? AND fp_bs.id_beasiswa = ?
	`
	row := b.db.QueryRow(
		query, 
		beasiswaSiswaStatusUpdate.Id, 
		beasiswaSiswaStatusUpdate.IdSiswa,
		beasiswaSiswaStatusUpdate.IdBeasiswa)

	beasiswaSiswa := &entity.BeasiswaSiswa{}
	if err := row.Scan(
		&beasiswaSiswa.Id,
		&beasiswaSiswa.IdSiswa,
		&beasiswaSiswa.NamaSiswa,
		&beasiswaSiswa.IdBeasiswa,
		&beasiswaSiswa.NamaBeasiswa,
		&beasiswaSiswa.IdMitra,
		&beasiswaSiswa.NamaMitra,
		&beasiswaSiswa.Status,
		&beasiswaSiswa.TanggalDaftar,
	); err != nil {
		return nil, err
	}

	return beasiswaSiswa, nil
}

