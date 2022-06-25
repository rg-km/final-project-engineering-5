package service

import (
	"FinalProject/api/repository"
	"FinalProject/auth"
	"FinalProject/entity"
	"FinalProject/payload"
	"FinalProject/utility"
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
		Email: mitra.Email,
		Role: mitra.KategoriUser,
		IdUser: mitra.Id,
		Token: tokenString,
	}, nil
}

func (m *mitraServiceImpl) RegisterMitra(request payload.RegisterMitraDetailRequest) (*payload.LoginResponse, error) {
	mitra, err := m.mitraRepository.RegisterMitra(&entity.MitraDetail{
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

	tokenString, err := auth.CreateJWTToken(mitra.Email, "MITRA", mitra.Id)
	if err != nil {
		return nil, err
	}

	return &payload.LoginResponse{
		Email: mitra.Email,
		Role: "MITRA",
		IdUser: mitra.Id,
		Token: tokenString,
	}, nil

}
