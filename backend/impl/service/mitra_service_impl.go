package service

import (
	"FinalProject/api/repository"
	"FinalProject/auth"
	"FinalProject/entity"
	"FinalProject/payload"
	"FinalProject/utility"
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
	isThere, err := m.mitraRepository.IsMitraExistsByEmail(request.Email)
	if err != nil {
		return nil, err
	}

	if !isThere {
		return nil, utility.ErrNoDataFound
	}
	
	mitra, err := m.mitraRepository.Login(request.Email, request.Password)
	if err != nil {
		return nil, err
	}
	
	tokenString, err := auth.CreateJWTToken(mitra.Email, mitra.KategoriUser, mitra.Id)
	if err != nil {
		return nil, err
	}

	return &payload.LoginResponse{
		Role:  mitra.KategoriUser,
		Token: tokenString,
	}, nil
}

func (m *mitraServiceImpl) RegisterMitra(request payload.MitraDetail) (*payload.LoginResponse, error) {
	mitra, err := m.mitraRepository.RegisterMitra(&entity.MitraDetail{
		Id:       request.Id,
		IdUser:   request.IdUser,
		Nama:     request.Nama,
		About:    request.About,
		NomorPic: request.NomorPic,
		NamaPic:  request.NamaPic,
		Situs:    request.Situs,
		Alamat:   request.Alamat,
		Email:    request.Email,
	}, &entity.Mitra{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return nil, err
	}

	claims := payload.Claims{
		Email: mitra.Email,
		Role:  "MITRA",
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
		Role:  "MITRA",
		Token: tokenString,
	}, nil

}
