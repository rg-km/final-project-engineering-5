package repository

import (
	"FinalProject/entity"
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type beasiswaRepositoryImpl struct {
	db *sql.DB
}

func NewBeasiswaRepositoryImpl(db *sql.DB) *beasiswaRepositoryImpl {
	return &beasiswaRepositoryImpl{
		db: db,
	}
}

func (b *beasiswaRepositoryImpl) GetTotalBeasiswa(nama string) (int, error) {
	count := 0
	
	query := `
	SELECT
		COUNT(id)
	FROM
		fp_beasiswa
	`

	if len(strings.Trim(nama, " ")) != 0 {
		query = fmt.Sprintf(`
		SELECT
			COUNT(id)
		FROM
			fp_beasiswa
		WHERE judul_beasiswa LIKE "%s%s%s"`, "%", nama, "%")
	}

	row := b.db.QueryRow(query)
	if err := row.Scan(
		&count,
	); err != nil {
		return -1, err
	}

	return count, nil
}

func (b *beasiswaRepositoryImpl) GetListBeasiswa(page int, limit int, nama string) ([]*entity.Beasiwa, error) {
	offset := limit * (page-1)

	query := `
	SELECT
		fp_b.id, fp_b.id_mitra, fp_m.nama, fp_b.judul_beasiswa, fp_b.benefits, fp_b.deskripsi, fp_b.tanggal_pembukaan, fp_b.tanggal_penutupan
	FROM
		fp_beasiswa fp_b
	INNER JOIN
		fp_mitra fp_m
	ON
		fp_b.id_mitra = fp_m.id
	LIMIT ?
	OFFSET ?
	`

	if len(strings.Trim(nama, " ")) != 0 {
		query = fmt.Sprintf(`
		SELECT
			fp_b.id, fp_b.id_mitra, fp_m.nama, fp_b.judul_beasiswa, fp_b.benefits, fp_b.deskripsi, fp_b.tanggal_pembukaan, fp_b.tanggal_penutupan
		FROM
			(
				SELECT
					id, id_mitra, judul_beasiswa, benefits, deskripsi, tanggal_pembukaan, tanggal_penutupan
				FROM
					fp_beasiswa WHERE judul_beasiswa LIKE "%s%s%s"
			) AS fp_b
		INNER JOIN
			fp_mitra fp_m
		ON
			fp_b.id_mitra = fp_m.id
		LIMIT ?
		OFFSET ?`, "%", nama, "%")
	}
	log.Println(query)
	log.Println(nama)

	rows, err := b.db.Query(query, limit, offset)
	if err != nil {
		return []*entity.Beasiwa{}, err
	}
	defer rows.Close()

	listBeasiswa := make([]*entity.Beasiwa, 0)
	for rows.Next() {
		row := &entity.Beasiwa{}
		if err := rows.Scan(
			&row.Id,
			&row.IdMitra,
			&row.NamaMitra,
			&row.JudulBeasiwa,
			&row.Benefist,
			&row.Deskripsi,
			&row.TanggalPembukaan,
			&row.TanggalPenutupan,
		); err != nil {
			return []*entity.Beasiwa{}, err
		}

		listBeasiswa = append(listBeasiswa, row)
	}

	return listBeasiswa, nil
}