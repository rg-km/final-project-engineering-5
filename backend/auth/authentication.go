package auth

import (
	"FinalProject/payload"
	"FinalProject/utility"
	"net/http"
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

func CreateJWTToken(email string, role string, idUser int) (string, error) {
	if len(strings.Trim(email, " ")) == 0 || len(strings.Trim(role, " ")) == 0 {
		return "", utility.ErrBadRequest
	}

	claims := payload.Claims{
		Email: email,
		Role: role,
		IdUser: idUser,
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

func ExtractJwtFromHeader(r *http.Request) string {
	tokenString := r.Header.Get("Authorization")
	if len(strings.Split(tokenString, " ")) == 2 {
		return strings.Split(tokenString, " ")[1]
	}
	return ""
}

func GetClaimsFromJwt(tokenString string) (*payload.Claims, error) {
	var claims *payload.Claims
	token, err := jwt.ParseWithClaims(tokenString, &payload.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return nil, err
	}

	if token != nil {
		if claims, ok := token.Claims.(*payload.Claims); ok && token.Valid {
			return claims, nil
		}
	}

	if !token.Valid {
		return nil, utility.ErrUnauthorized
	}

	return claims, nil
}