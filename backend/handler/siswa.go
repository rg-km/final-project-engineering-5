package handler

import (
	"FinalProject/repository"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Handler struct {
	Repository repository.Repository
}

func NewHandler(repo repository.Repository) Handler {
	return Handler{
		Repository: repo,
	}
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

var secretKey = []byte("Final Project Beasiswa")

func (s *Handler) Login(c *gin.Context) {
	var siswaLogin repository.Login
	// email := c.PostForm("email")
	// password := c.PostForm("password")
	err := json.NewDecoder(c.Request.Body).Decode(&siswaLogin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	siswa, err := s.Repository.Login(siswaLogin.Email, siswaLogin.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}
	claims := Claims{
		Email: *siswa,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(60 * time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"role":  siswa,
		"token": tokenString,
	})

}
