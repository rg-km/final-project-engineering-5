package payload

type MitraDetail struct {
	Id       int    `json:"id"`
	IdUser   int    `json:"idUser"`
	Nama     string `json:"nama"`
	About    string `json:"about"`
	NomorPic string `json:"nomorPic"`
	NamaPic  string `json:"namaPic"`
	Situs    string `json:"situs"`
	Alamat   string `json:"alamat"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterMitraDetailRequest struct {
	Nama     string `json:"nama" binding:"required"`
	About    string `json:"about" binding:"required"`
	NomorPic string `json:"nomorPic" binding:"required"`
	NamaPic  string `json:"namaPic" binding:"required"`
	Situs    string `json:"situs" binding:"required"`
	Alamat   string `json:"alamat" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

