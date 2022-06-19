package repository

import (
	"FinalProject/entity"
	"FinalProject/utility"
	"database/sql"
	"fmt"
	"strings"
	"sync"
)

type beasiswaRepositoryImpl struct {
	db *sql.DB
	mu *sync.Mutex
}

func NewBeasiswaRepositoryImpl(db *sql.DB, mu *sync.Mutex) *beasiswaRepositoryImpl {
	return &beasiswaRepositoryImpl{
		db: db,
		mu: mu,
	}
}

func (b *beasiswaRepositoryImpl) GetBeasiswaById(id int) (*entity.Beasiswa, error) {
	query := `
	SELECT
	id, id_mitra, benefits, judul_beasiswa, deskripsi, tanggal_pembukaan, tanggal_penutupan
	FROM
		fp_beasiswa
	WHERE
		id = ?
	`
	
	row := b.db.QueryRow(query, id)
	
	beasiswa := &entity.Beasiswa{}
	if err := row.Scan(
		&beasiswa.Id,
		&beasiswa.IdMitra,
		&beasiswa.Benefits,
		&beasiswa.JudulBeasiswa,
		&beasiswa.Deskripsi,
		&beasiswa.TanggalPembukaan,
		&beasiswa.TanggalPenutupan,
	); err != nil {
		return nil, err
	}
	
	return beasiswa, nil
}

func (b *beasiswaRepositoryImpl) IsBeasiswaExistsById(id int) (bool, error) {
	count := 0

	query := `
	SELECT
		COUNT(id)
	FROM
		fp_beasiswa
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
	offset := limit * (page - 1)

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

func (b *beasiswaRepositoryImpl) CreateBeasiswa(beasiswa *entity.Beasiswa) (*entity.Beasiswa, error) {
	query := `
	INSERT INTO fp_beasiswa (
		id_mitra, judul_beasiswa, benefits, deskripsi, tanggal_pembukaan, tanggal_penutupan
	) 
	VALUES 
	(?, ?, ?, ?, ?, ?)
	
	`
	result, err := b.db.Exec(query,
		beasiswa.IdMitra,
		beasiswa.JudulBeasiswa,
		beasiswa.Benefits,
		beasiswa.Deskripsi,
		beasiswa.TanggalPembukaan,
		beasiswa.TanggalPenutupan,
	)
	if err != nil {
		return nil, err
	}
	
	lastId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	beasiswa.Id = int(lastId)
	return beasiswa, nil
}

func (b *beasiswaRepositoryImpl) UpdateBeasiswa(beasiswa entity.Beasiswa, id int) (*entity.Beasiswa, error) {
	query := `
	UPDATE
		fp_beasiswa
	SET
		id = ?,
		id_mitra = ?,
		judul_beasiswa = ?,
		benefits = ?,
		deskripsi = ?,
		tanggal_pembukaan = ?,
		tanggal_penutupan = ?
	WHERE
		id = ?
	`

	_, err := b.db.Exec(
		query,
		beasiswa.Id,
		beasiswa.IdMitra,
		beasiswa.JudulBeasiswa,
		beasiswa.Benefits,
		beasiswa.Deskripsi,
		beasiswa.TanggalPembukaan,
		beasiswa.TanggalPenutupan,
		id,
		);
	if err != nil {
		return nil, err
	}

	updatedBeasiswa, err := b.GetBeasiswaById(id)
	if err != nil {
		return nil, err
	}

	return updatedBeasiswa, nil
}
