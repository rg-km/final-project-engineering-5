package entity

type (
	BeasiswaSiswaStatusUpdate struct {
		Id int
		IdSiswa int
		IdBeasiswa int
		Status string
	}

	BeasiswaSiswa struct {
		Id int
		IdSiswa int
		NamaSiswa string
		IdBeasiswa int
		NamaBeasiswa string
		IdMitra int
		NamaMitra string
		Status string
		TanggalDaftar string
	}
)
