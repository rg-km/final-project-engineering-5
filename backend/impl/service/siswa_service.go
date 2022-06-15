package service

import (
	"FinalProject/api/repository"
	"FinalProject/payload"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte("Final Project Beasiswa")

type Claims struct {
	Email string 	`json:"email"`
	Role string 	`json:"role"`
	jwt.StandardClaims
}

type siswaServiceImpl struct {
	siswaRepository repository.SiswaRepository
}

func NewSiswaServiceImpl(siswaRepository repository.SiswaRepository) *siswaServiceImpl {
	return &siswaServiceImpl{
		siswaRepository: siswaRepository,
	}
}

func (s *siswaServiceImpl) Login(request payload.LoginRequest) (*payload.LoginResponse, error) {
	siswa, err := s.siswaRepository.Login(request.Email, request.Password)
	if err != nil {
		return nil, err
	}
	
	claims := Claims{
		Email: siswa.Email,
		Role: siswa.KategoriUser,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(3 * 60 * time.Minute).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSting, err := token.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	return &payload.LoginResponse{
		Role: siswa.KategoriUser,
		Token: tokenSting,
	}, nil
}