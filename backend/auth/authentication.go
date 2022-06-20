package auth

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	COST = 10
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
