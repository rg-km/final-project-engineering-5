package payload

type (
	Beasiswa struct {
		Id int
		IdMitra int
		NamaMitra string
		JudulBeasiwa string
		Benefist string
		Deskripsi string
		TanggalPembukaan string
		TanggalPenutupan string
	}

	ListBeasiswaRequest struct {
		Page int `json:"page"`
		Limit int `json:"limit"`
		Nama string `json:"nama"`
	}
	
	ListBeasiswaResponse struct {
		Data []Beasiswa `json:"data"`
		PaginateInfo PaginateInfo `json:"paginateInfo"`
	}
)