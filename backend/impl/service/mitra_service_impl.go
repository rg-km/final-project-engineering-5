package service

import (
	"FinalProject/api/repository"
	"FinalProject/payload"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type mitraServiceImpl struct {
	mitraRepository repository.MitraRepository
}

func NewMitraServiceImpl(mitraRepository repository.MitraRepository) *mitraServiceImpl {
	return &mitraServiceImpl{
		mitraRepository: mitraRepository,
	}
}

func (m *mitraServiceImpl) Login(request payload.LoginRequest) (*payload.LoginResponse, error) {
	mitra, err := m.mitraRepository.Login(request.Email, request.Password)
	if err != nil {
		return nil, err
	}

	claims := payload.Claims{
		Email: mitra.Email,
		Role: mitra.KategoriUser,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(3 * 60 * time.Minute).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	return &payload.LoginResponse{
		Role: mitra.KategoriUser,
		Token: tokenString,
	}, nil
}