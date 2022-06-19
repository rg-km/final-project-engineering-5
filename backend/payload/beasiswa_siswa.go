package payload

type (
	BeasiswaSiswa struct{
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

	BeasiswaSiswaStatusUpdateRequest struct {
		Id int `json:"id"`
		IdSiswa int `json:"idSiswa"`
		IdBeasiswa int `json:"idBeasiswa"`
		Status string `json:"status"`
	}

	BeasiswaSiswaStatusUpdateResponse struct {
		Message string `json:"message"`
		BeasiswaSiswa BeasiswaSiswa `json:"beasiswaSiswa"`
	}
)