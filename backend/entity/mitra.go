package entity

type (
	Mitra struct {
		Id           int
		Email        string
		Password     string
		KategoriUser string
		Mitra        MitraDetail
	}

	MitraDetail struct {
		Id       int
		IdUser   int
		Nama     string
		About    string
		NomorPic string
		NamaPic  string
		Situs    string
		Alamat   string
		Email    string
	}
)
