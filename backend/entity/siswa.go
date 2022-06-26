package entity

type (
	Siswa struct {
		Id           int
		Email        string
		Password     string
		KategoriUser string
		Siswa        SiswaDetail
	}

	SiswaDetail struct {
		Id                int
		IdUser            int
		Email             string
		Nama              string
		TanggalLahir      string
		NomorTelepon      string
		NamaInstansi      string
		TingkatPendidikan string
		NomorRekening     string
		NamaBank          string
		Alamat            string
	}
)
