package payload

type Beasiswa struct {
	Id               int    `json:"id"`
	IdMitra          int    `json:"id_mitra"`
	JudulBeasiswa    string `json:"judul_beasiswa"`
	Deskripsi        string `json:"deskripsi"`
	TanggalPembukaan string `json:"tanggal_pembukaan"`
	TanggalPenutupan string `json:"tanggal_penutupan"`
	Benefits         string `json:"benefits"`
}

type BeasiswaResponse struct {
	Data []Beasiswa `json:"data"`
}
