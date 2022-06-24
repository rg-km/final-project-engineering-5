package payload


type (
	Beasiswa struct {
		Id               int    `json:"id"`
		IdMitra          int    `json:"idMitra"`
		JudulBeasiswa    string `json:"judulBeasiswa"`
		Deskripsi        string `json:"deskripsi"`
		TanggalPembukaan string `json:"tanggalPembukaan"`
		TanggalPenutupan string `json:"tanggalPenutupan"`
		Benefits         string `json:"benefits" binding:"required"`
	}
	
	CreateBeasiswaRequest struct {
		IdMitra          int    `json:"idMitra" binding:"required"`
		JudulBeasiswa    string `json:"judulBeasiswa" binding:"required"`
		Deskripsi        string `json:"deskripsi" binding:"required"`
		TanggalPembukaan string `json:"tanggalPembukaan" binding:"required"`
		TanggalPenutupan string `json:"tanggalPenutupan" binding:"required"`
		Benefits         string `json:"benefits" binding:"required"`
	}

	UpdateBeasiswaRequest struct {
		Id               int    `json:"id" binding:"required"`
		IdMitra          int    `json:"idMitra" binding:"required"`
		JudulBeasiswa    string `json:"judulBeasiswa" binding:"required"`
		Deskripsi        string `json:"deskripsi" binding:"required"`
		TanggalPembukaan string `json:"tanggalPembukaan" binding:"required"`
		TanggalPenutupan string `json:"tanggalPenutupan" binding:"required"`
		Benefits         string `json:"benefits" binding:"required"`
	}

	BeasiswaResponse struct {
		Message  string   `json:"message"`
		Beasiswa Beasiswa `json:"beasiswa"`
	}

	DeleteBeasiswaResponse struct {
		Message string `json:"message"`
	}

	ListBeasiswaRequest struct {
		Page  int    `json:"page"`
		Limit int    `json:"limit"`
		Nama  string `json:"nama"`
	}
	
	ListBeasiswaResponse struct {
		Data         []Beasiswa   `json:"data"`
		PaginateInfo PaginateInfo `json:"paginateInfo"`
	}
	
)
