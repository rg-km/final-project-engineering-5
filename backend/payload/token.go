package payload

import (
	"github.com/golang-jwt/jwt/v4"
)
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Email string `json:"email"`
	Role string `json:"role"`
	IdUser int `json:"idUser"`
	Token 	string `json:"token"`
}

type Claims struct {
	Email string
	Role string
	IdUser int
	jwt.StandardClaims
}