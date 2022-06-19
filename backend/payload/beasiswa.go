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
	Message string `json:"message"`
	Beasiswa Beasiswa `json:"beasiswa"`
}

type ListBeasiswaRequest struct {
	Page int `json:"page"`
	Limit int `json:"limit"`
	Nama string `json:"nama"`
}

type ListBeasiswaResponse struct {
	Data []Beasiswa `json:"data"`
	PaginateInfo PaginateInfo `json:"paginateInfo"`
}
