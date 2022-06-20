package auth

import (
	"FinalProject/payload"
	"FinalProject/utility"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

const (
	COST = 10
	SECRET_KEY = "Final Project Beasiswa"
)

func CreatePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), COST)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func ComparePassword(hashedPassword string, password string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false, nil
	}
	return true, nil
}

func CreateJWTToken(email string, role string) (string, error) {
	if len(strings.Trim(email, " ")) == 0 || len(strings.Trim(role, " ")) == 0 {
		return "", utility.ErrBadRequest
	}

	claims := payload.Claims{
		Email: email,
		Role: role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * 60 * time.Minute).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
