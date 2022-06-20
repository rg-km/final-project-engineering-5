package utility

import (
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "beasiswa.db")
	if err != nil {
		log.Panicln("Failed connect to db", err)
	}

	return db
}

func MigrationDB(db *sql.DB) error {
	log.Printf("Migrasi database sedang dijalankan...")
	_, err := db.Exec(`
	DROP TABLE IF EXISTS fp_beasiswa_siswa;
	DROP TABLE IF EXISTS fp_beasiswa;
	DROP TABLE IF EXISTS fp_mitra;
	DROP TABLE IF EXISTS fp_siswa;
	DROP TABLE IF EXISTS fp_user;

	CREATE TABLE IF NOT EXISTS fp_user (
		id integer PRIMARY KEY AUTOINCREMENT,
		email varchar(150) NOT NULL UNIQUE,
		password varchar(150) NOT NULL,
		kategori_user varchar(50) NOT NULL
	);

	CREATE TABLE IF NOT EXISTS fp_siswa (
		id integer PRIMARY KEY AUTOINCREMENT,
		id_user integer NULL,
		nama varchar (150),
		tanggal_lahir date NOT NULL,
		nomor_telepon varchar(20) NOT NULL UNIQUE,
		nama_instansi VARCHAR(150) NOT NULL,
		tingkat_pendidikan varchar(50) NOT NULL,
		nomor_rekening varchar(100) NOT NULL UNIQUE,
		nama_bank varchar(100) NOT NULL,
		alamat TEXT NOT NULL,
		FOREIGN KEY (id_user) REFERENCES fp_user (id) ON UPDATE CASCADE ON DELETE RESTRICT
	);

	CREATE TABLE IF NOT EXISTS fp_mitra (
		id integer PRIMARY KEY AUTOINCREMENT,
		id_user integer NULL,
		nama varchar(150) NOT NULL,
		about text NOT NULL,
		nomor_pic varchar(20) NOT NULL UNIQUE,
		nama_pic varchar(150) NOT NULL,
		situs varchar(150) NOT NULL,
		alamat TEXT NOT NULL,
		FOREIGN KEY (id_user) REFERENCES fp_user (id) ON UPDATE CASCADE ON DELETE RESTRICT
	);

	CREATE TABLE IF NOT EXISTS fp_beasiswa (
		id integer PRIMARY KEY AUTOINCREMENT,
		id_mitra integer NULL,
		judul_beasiswa text NOT NULL,
		benefits text NOT NULL,
		deskripsi text NOT NULL,
		tanggal_pembukaan date NOT NULL,
		tanggal_penutupan date NOT NULL,
		FOREIGN KEY (id_mitra) REFERENCES fp_mitra (id) ON UPDATE CASCADE ON DELETE RESTRICT
	);
	
	CREATE TABLE IF NOT EXISTS fp_beasiswa_siswa (
		id integer PRIMARY KEY AUTOINCREMENT,
		id_siswa integer NULL,
		id_beasiswa integer NULL,
		tanggal_daftar date NOT NULL,
		status varchar NOT NULL,
		FOREIGN KEY (id_siswa) REFERENCES fp_siswa (id) ON UPDATE CASCADE ON DELETE RESTRICT,
		FOREIGN KEY (id_beasiswa) REFERENCES fp_beasiswa (id) ON UPDATE CASCADE ON DELETE RESTRICT
	);

	INSERT INTO fp_user 
		(id, email, password, kategori_user)
	VALUES
		(1, "denny@email.com", "123456", "SISWA"),
		(2, "rezky@email.com", "123456", "MITRA"),
		(3, "sinulingga@email.com", "123456", "SISWA");
	
	INSERT INTO fp_siswa
		(id, id_user, nama, tanggal_lahir, nomor_telepon, nama_instansi, tingkat_pendidikan, nomor_rekening, nama_bank, alamat)
	VALUES
		(1, 1, "Denny", "2020-06-15", "0123456789012", "Universitas Cinta Damai", "S1", "0138913139739793", "Bank Peyimpanan", "Jl. Cinta Damai, Kel. Kedaiaman"),
		(2, 3, "Sinulingga", "2020-06-15", "0123456789032", "Universitas Cinta Damai", "SMA", "0138913139739733", "Bank Peyimpanan Juga", "Jl. Cinta Damai, Kel. Kedaiaman");
	
	INSERT INTO fp_mitra
		(id, id_user, nama, about, nomor_pic, nama_pic, situs, alamat)
	VALUES
		(1,2, "PT. Maju Bersama, Tbk", "Maju Bersama (part of Maju Group) telah 50 tahun menjadi perusahaan yang mengedepankan customer", "098765432112", "Stefani", "www.majubersama.com", "Jakarta");
	
	INSERT INTO fp_beasiswa
		(id, id_mitra, judul_beasiswa, benefits, deskripsi, tanggal_pembukaan, tanggal_penutupan)
	VALUES
		(1, 1, "Leadership Scholarship 2022", "1.Uang Saku\n2.Mentoring\n3.Networking", "Apakah kamu...", "2022-06-15", "2020-08-15"),
		(2, 1, "Leadership Scholarship 2021", "1.Uang Saku\n2.Mentoring\n3.Networking", "Apakah kamu...", "2022-06-15", "2020-08-15"),
		(3, 1, "Leadership Scholarship 2020", "1.Uang Saku\n2.Mentoring\n3.Networking", "Apakah kamu...", "2022-06-15", "2020-08-15");
	`)

	return err
}