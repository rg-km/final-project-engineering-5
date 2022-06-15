package payload

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Role 	string `json:"role"`
	Token 	string `json:"token"`
}