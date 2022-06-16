package repository

import (
	"FinalProject/entity"
	"FinalProject/utility"
	"database/sql"
	"log"
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
