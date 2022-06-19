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
