package payload

type Siswa struct {
	Id                int    `json:"id"`
	Nama              string `json:"nama"`
	NamaInstansi      string `json:"nama_instansi"`
	TingkatPendidikan string `json:"tingkat_pendidikan"`
	Alamat            string `json:"alamat"`
	NoHp              string `json:"no_hp"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	TanggalLahir      string `json:"tanggal_lahir"`
	Rekening          int    `json:"rekening"`
}

type Mitra struct {
	Id       int    `json:"id"`
	Nama     string `json:"nama"`
	Profile  string `json:"profile"`
	Pic      string `json:"pic"`
	Situs    string `json:"situs"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Role 	string `json:"role"`
	Token 	string `json:"token"`
}
