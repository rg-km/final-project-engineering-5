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
		Id         int    `json:"id"`
		IdSiswa    int    `json:"idSiswa"`
		IdBeasiswa int    `json:"idBeasiswa"`
		Status     string `json:"status"`
	}

	BeasiswaSiswaStatusUpdateResponse struct {
		Message       string        `json:"message"`
		BeasiswaSiswa BeasiswaSiswa `json:"beasiswaSiswa"`
	}

	BeasiswaSiswaApplyRequest struct {
		IdSiswa    int `json:"idSiswa"`
		IdBeasiswa int `json:"idBeasiswa"`
	}

	BeasiswaSiswaApplyResponse struct {
		Message       string        `json:"message"`
		BeasiswaSiswa BeasiswaSiswa `json:"beasiswaSiswa"`
	}
)
