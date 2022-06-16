package repository

import (
	"FinalProject/entity"
	"FinalProject/utility"
	"database/sql"
	"log"
	"strings"
	"fmt"
)

type beasiswaRepositoryImpl struct {
	db *sql.DB
}

func NewBeasiswaRepositoryImpl(db *sql.DB) *beasiswaRepositoryImpl {
	return &beasiswaRepositoryImpl{
		db: db,
	}
}

func (b *beasiswaRepositoryImpl) GetBeasiswaById(id string) ([]*entity.Beasiswa, error) {
	query := `
	SELECT
	id, id_mitra, benefits, judul_beasiswa, deskripsi, tanggal_pembukaan, tanggal_penutupan
	FROM
		fp_beasiswa
	WHERE
		id = ?
	`

	rows, err := b.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	beasiswa := make([]*entity.Beasiswa, 0)
	for rows.Next() {
		beasiswaItem := entity.Beasiswa{}

		err = rows.Scan(
			&beasiswaItem.Id,
			&beasiswaItem.IdMitra,
			&beasiswaItem.Benefits,
			&beasiswaItem.JudulBeasiswa,
			&beasiswaItem.Deskripsi,
			&beasiswaItem.TanggalPembukaan,
			&beasiswaItem.TanggalPenutupan,
		)
		beasiswa = append(beasiswa, &beasiswaItem)
	}
	log.Println(query)
	log.Println(beasiswa)

	if err != nil {
		return nil, err
	}

	if len(beasiswa) == 0 {
		return nil, utility.ErrNoDataFound
	}

	return beasiswa, nil
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

func (b *beasiswaRepositoryImpl) GetListBeasiswa(page int, limit int, nama string) ([]*entity.Beasiswa, error) {
	offset := limit * (page-1)

	query := `
	SELECT
		id, id_mitra, judul_beasiswa, benefits, deskripsi, tanggal_pembukaan, tanggal_penutupan
	FROM
		fp_beasiswa
	LIMIT ?
	OFFSET ?
	`

	if len(strings.Trim(nama, " ")) != 0 {
		query = fmt.Sprintf(`
		SELECT
			id, id_mitra, judul_beasiswa, benefits, deskripsi, tanggal_pembukaan, tanggal_penutupan
		FROM
			(
				SELECT
					id, id_mitra, judul_beasiswa, benefits, deskripsi, tanggal_pembukaan, tanggal_penutupan
				FROM
					fp_beasiswa WHERE judul_beasiswa LIKE "%s%s%s"
			) AS fp_beasiswa
		LIMIT ?
		OFFSET ?`, "%", nama, "%")
	}
	log.Println(query)
	log.Println(nama)

	rows, err := b.db.Query(query, limit, offset)
	if err != nil {
		return []*entity.Beasiswa{}, err
	}
	defer rows.Close()

	listBeasiswa := make([]*entity.Beasiswa, 0)
	for rows.Next() {
		row := &entity.Beasiswa{}
		if err := rows.Scan(
			&row.Id,
			&row.IdMitra,
			&row.JudulBeasiswa,
			&row.Benefits,
			&row.Deskripsi,
			&row.TanggalPembukaan,
			&row.TanggalPenutupan,
		); err != nil {
			return []*entity.Beasiswa{}, err
		}

		listBeasiswa = append(listBeasiswa, row)
	}

	return listBeasiswa, nil
}
