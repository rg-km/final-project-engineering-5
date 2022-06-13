package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "beasiswa.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS siswa (
			id INTEGER PRIMARY KEY,
			nama VARCHAR(255),
			nama_instansi VARCHAR(255),
			tingkat_pendidikan VARCHAR(255),
			alamat VARCHAR(255),
			no_hp VARCHAR(255),
			email VARCHAR(255),
			password VARCHAR(255),
			tanggal_lahir DATE,
			rekening INTEGER
		);

		CREATE TABLE IF NOT EXISTS mitra (
			id INTEGER PRIMARY KEY,
			nama VARCHAR(255),
			profile TEXT,
			pic INTEGER,
			situs VARCHAR(255),
			email VARCHAR(255),
			password VARCHAR(255)
		);

		INSERT INTO mitra (nama, profile, pic, situs, email, password) VALUES ("Gudang Garam", "Pabrik Roko", "087777123", "gudanggaram.com", "garam@gmail.com", "123456");

		INSERT INTO siswa (nama, nama_instansi, tingkat_pendidikan, alamat, no_hp, email, password, tanggal_lahir, rekening) 
		VALUES ("Siswa 1", "Instansi 1", "S1", "Alamat 1", "081234567890", "hatta@gmail.com", "123456", "2020-01-01", "123456789");
	`)
	if err != nil {
		panic(err)
	}

}
