package payload

type (
	BeasiswaSiswa struct {
		Id            int    `json:"id"`
		IdSiswa       int    `json:"idSiswa"`
		NamaSiswa     string `json:"namaSiswa"`
		IdBeasiswa    int    `json:"idBeasiswa"`
		NamaBeasiswa  string `json:"namaBeasiswa"`
		IdMitra       int    `json:"idMitra"`
		NamaMitra     string `json:"namaMitra"`
		Status        string `json:"status"`
		TanggalDaftar string `json:"tanggalDaftar"`
	}

	BeasiswaSiswaStatusUpdateRequest struct {
		Id         int    `json:"id" binding:"required"`
		IdSiswa    int    `json:"idSiswa" binding:"required"`
		IdBeasiswa int    `json:"idBeasiswa" binding:"required"`
		Status     string `json:"status" binding:"required"`
	}

	BeasiswaSiswaStatusUpdateResponse struct {
		Message       string        `json:"message"`
		BeasiswaSiswa BeasiswaSiswa `json:"beasiswaSiswa"`
	}

	BeasiswaSiswaApplyRequest struct {
		IdSiswa    int `json:"idSiswa" binding:"required"`
		IdBeasiswa int `json:"idBeasiswa" binding:"required"`
	}

	BeasiswaSiswaApplyResponse struct {
		Message       string        `json:"message"`
		BeasiswaSiswa BeasiswaSiswa `json:"beasiswaSiswa"`
	}

	ListBeasiswaSiswaByMitraIdRequest struct {
		IdMitra int    `json:"idMitra" binding:"required"`
		Page    int    `json:"page"`
		Limit   int    `json:"limit"`
		Nama    string `json:"nama"`
	}

	ListBeasiswaSiswaByMitraIdResponse struct {
		Data         []BeasiswaSiswa `json:"data"`
		PaginateInfo PaginateInfo    `json:"paginateInfo"`
	}

	ListBeasiswaSiswaBySiswaIdRequest struct {
		IdSiswa int `json:"idSiswa" binding:"required"`
	}

	ListBeasiswaSiswaBySiswaIdResponse struct {
		Data []BeasiswaSiswa `json:"data"`
	}
)
