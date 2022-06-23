package repository

import (
	"FinalProject/entity"
	"FinalProject/utility"
	"database/sql"
	"time"
)

type beasiswaSiswaRepositoryImpl struct {
	db *sql.DB
}

func NewBeasiswaSiswaRepositoryImpl(db *sql.DB) *beasiswaRepositoryImpl {
	return &beasiswaRepositoryImpl{
		db: db,
	}
}

func (b *beasiswaRepositoryImpl) IsBeasiswaSiswaExistsById(id int) (bool, error) {
	count := 0

	query := `
	SELECT
		COUNT(id)
	FROM
		fp_beasiswa_siswa
	WHERE
		id = ?
	`

	row := b.db.QueryRow(query, id)
	if err := row.Scan(
		&count,
	); err != nil {
		return false, err
	}

	if count != 1 {
		return false, utility.ErrNoDataFound
	}

	return true, nil
}

func (b *beasiswaRepositoryImpl) UpdateStatusBeasiswa(
	beasiswaSiswaStatusUpdate entity.BeasiswaSiswaStatusUpdate, id int) (*entity.BeasiswaSiswa, error) {
	query := `
	UPDATE
		fp_beasiswa_siswa
	SET
		status = ?
	WHERE
		id = ? AND id_siswa = ? AND id_beasiswa = ?
	`
	_, err := b.db.Exec(
		query,
		beasiswaSiswaStatusUpdate.Status,
		id,
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
		fp_bs.tanggal_daftar
	FROM
		fp_beasiswa_siswa fp_bs
	LEFT JOIN
		fp_beasiswa fp_b
	ON
		fp_bs.id_beasiswa = fp_b.id
	INNER JOIN
		fp_mitra fp_m
	ON
		fp_b.id_mitra = fp_m.id
	INNER JOIN
		fp_siswa fp_s
	ON
		fp_bs.id_siswa = fp_s.id
	WHERE
		fp_bs.id = ? AND fp_bs.id_siswa = ? AND fp_bs.id_beasiswa = ?
	`
	row := b.db.QueryRow(
		query,
		id,
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

func (b *beasiswaRepositoryImpl) GetBeasiswaSiswaById(id int) (*entity.BeasiswaSiswa, error) {
	query := `
	SELECT
		fp_bs.id,
		fp_bs.id_siswa,
		fp_s.nama,
		fp_bs.id_beasiswa,
		fp_b.judul_beasiswa,
		fp_b.id_mitra,
		fp_m.nama,
		fp_bs.status,
		fp_bs.tanggal_daftar
	FROM
		fp_beasiswa_siswa fp_bs
	LEFT JOIN
		fp_beasiswa fp_b
	ON
		fp_bs.id_beasiswa = fp_b.id
	INNER JOIN
		fp_mitra fp_m
	ON
		fp_b.id_mitra = fp_m.id
	INNER JOIN
		fp_siswa fp_s
	ON
		fp_bs.id_siswa = fp_s.id
	WHERE
		fp_bs.id = ?
	`
	row := b.db.QueryRow(query, id)

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

func (b *beasiswaRepositoryImpl) GetListBeasiswaSiswaByIdSiswa(id int) ([]*entity.BeasiswaSiswa, error) {
	query := `
	SELECT
		fp_bs.id,
		fp_bs.id_siswa,
		fp_s.nama,
		fp_bs.id_beasiswa,
		fp_b.judul_beasiswa,
		fp_b.id_mitra,
		fp_m.nama,
		fp_bs.status,
		fp_bs.tanggal_daftar
	FROM
		fp_beasiswa_siswa fp_bs
	LEFT JOIN
		fp_beasiswa fp_b
	ON
		fp_bs.id_beasiswa = fp_b.id
	INNER JOIN
		fp_mitra fp_m
	ON
		fp_b.id_mitra = fp_m.id
	INNER JOIN
		fp_siswa fp_s
	ON
		fp_bs.id_siswa = fp_s.id
	WHERE
		fp_bs.id_siswa = ?
	`
	rows, err := b.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	beasiswaSiswaList := []*entity.BeasiswaSiswa{}
	for rows.Next() {
		beasiswaSiswa := &entity.BeasiswaSiswa{}
		if err := rows.Scan(
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
		beasiswaSiswaList = append(beasiswaSiswaList, beasiswaSiswa)
	}

	return beasiswaSiswaList, nil
}

func (b *beasiswaRepositoryImpl) GetListBeasiswaSiswaByIdBeasiswa(id int) ([]*entity.BeasiswaSiswa, error) {
	query := `
	SELECT
		fp_bs.id,
		fp_bs.id_siswa,
		fp_s.nama,
		fp_bs.id_beasiswa,
		fp_b.judul_beasiswa,
		fp_b.id_mitra,
		fp_m.nama,
		fp_bs.status,
		fp_bs.tanggal_daftar
	FROM
		fp_beasiswa_siswa fp_bs
	LEFT JOIN
		fp_beasiswa fp_b
	ON
		fp_bs.id_beasiswa = fp_b.id
	INNER JOIN
		fp_mitra fp_m
	ON
		fp_b.id_mitra = fp_m.id
	INNER JOIN
		fp_siswa fp_s
	ON
		fp_bs.id_siswa = fp_s.id
	WHERE
		fp_bs.id_beasiswa = ?
	`
	rows, err := b.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	beasiswaSiswaList := []*entity.BeasiswaSiswa{}
	for rows.Next() {
		beasiswaSiswa := &entity.BeasiswaSiswa{}
		if err := rows.Scan(
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
		beasiswaSiswaList = append(beasiswaSiswaList, beasiswaSiswa)
	}

	return beasiswaSiswaList, nil
}

func (b *beasiswaRepositoryImpl) GetListBeasiswaSiswaByIdMitra(id int) ([]*entity.BeasiswaSiswa, error) {
	query := `
	SELECT
		fp_bs.id,
		fp_bs.id_siswa,
		fp_s.nama,
		fp_bs.id_beasiswa,
		fp_b.judul_beasiswa,
		fp_b.id_mitra,
		fp_m.nama,
		fp_bs.status,
		fp_bs.tanggal_daftar
	FROM
		fp_beasiswa_siswa fp_bs
	LEFT JOIN
		fp_beasiswa fp_b
	ON
		fp_bs.id_beasiswa = fp_b.id
	INNER JOIN
		fp_mitra fp_m
	ON
		fp_b.id_mitra = fp_m.id
	INNER JOIN
		fp_siswa fp_s
	ON
		fp_bs.id_siswa = fp_s.id
	WHERE
		fp_b.id_mitra = ?
	`
	rows, err := b.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	beasiswaSiswaList := []*entity.BeasiswaSiswa{}
	for rows.Next() {
		beasiswaSiswa := &entity.BeasiswaSiswa{}
		if err := rows.Scan(
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
		beasiswaSiswaList = append(beasiswaSiswaList, beasiswaSiswa)
	}

	return beasiswaSiswaList, nil
}

func (b *beasiswaRepositoryImpl) ApplyBeasiswa(idSiswa int, idBeasiswa int) (*entity.BeasiswaSiswa, error) {
	query := `
	INSERT INTO
		fp_beasiswa_siswa (id_siswa, id_beasiswa, status, tanggal_daftar)
	VALUES
		(?, ?, ?, ?)
	`
	stmt, err := b.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	tglDaftar := time.Now().Format("2022-06-23")
	res, err := stmt.Exec(
		idSiswa,
		idBeasiswa,
		"DIPROSES",
		tglDaftar,
	)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	beasiswaSiswa, err := b.GetBeasiswaSiswaById(int(id))
	if err != nil {
		return nil, err
	}

	return beasiswaSiswa, nil

}
