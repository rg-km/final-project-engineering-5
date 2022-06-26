package service

import (
	"FinalProject/api/repository"
	"FinalProject/auth"
	"FinalProject/entity"
	"FinalProject/payload"
	"FinalProject/utility"
	"log"
)

var secretKey = []byte("Final Project Beasiswa")

type siswaServiceImpl struct {
	siswaRepository         repository.SiswaRepository
	beasiswaSiswaRepository repository.BeasiswaSiswaRepostiroy
}

func NewSiswaServiceImpl(
	siswaRepository repository.SiswaRepository,
	beasiswaSiswaRepository repository.BeasiswaSiswaRepostiroy) *siswaServiceImpl {
	return &siswaServiceImpl{
		siswaRepository:         siswaRepository,
		beasiswaSiswaRepository: beasiswaSiswaRepository,
	}
}

func (s *siswaServiceImpl) GetSiswaById(id int) (*payload.SiswaDetailResponse, error) {
	isThere, err := s.siswaRepository.IsSiswaExistsById(id)
	if err != nil {
		log.Println("masalah:", err)
		return nil, err
	}

	if !isThere {
		return nil, utility.ErrNoDataFound
	}

	siswa, err := s.siswaRepository.GetSiswaById(id)
	if err != nil {
		log.Println("masalah4:", err)
		return nil, err
	}

	rowsBeasiswaSiswa, err := s.beasiswaSiswaRepository.GetListBeasiswaSiswaByIdSiswa(siswa.Id)
	if err != nil {
		log.Println("masalah3:", err)
		return nil, err
	}

	listBeasiswaSiswa := make([]payload.BeasiswaSiswa, 0)
	for _, beasiswaSiswa := range rowsBeasiswaSiswa {
		listBeasiswaSiswa = append(listBeasiswaSiswa, payload.BeasiswaSiswa{
			Id:            beasiswaSiswa.Id,
			IdSiswa:       beasiswaSiswa.IdSiswa,
			NamaSiswa:     beasiswaSiswa.NamaSiswa,
			IdBeasiswa:    beasiswaSiswa.IdBeasiswa,
			NamaBeasiswa:  beasiswaSiswa.NamaBeasiswa,
			IdMitra:       beasiswaSiswa.IdSiswa,
			NamaMitra:     beasiswaSiswa.NamaSiswa,
			Status:        beasiswaSiswa.Status,
			TanggalDaftar: beasiswaSiswa.TanggalDaftar,
		})
	}

	return &payload.SiswaDetailResponse{
		Siswa: payload.Siswa{
			Id:                siswa.Id,
			IdUser:            siswa.IdUser,
			Nama:              siswa.Nama,
			NamaInstansi:      siswa.NamaInstansi,
			TingkatPendidikan: siswa.TingkatPendidikan,
			Alamat:            siswa.Alamat,
			NomorTelepon:      siswa.NomorTelepon,
			Email:             siswa.Email,
			TanggalLahir:      siswa.TanggalLahir,
			NomorRekening:     siswa.NomorRekening,
			NamaBank:          siswa.NamaBank,
		},
		Data: listBeasiswaSiswa,
	}, nil
}

func (s *siswaServiceImpl) Login(request payload.LoginRequest) (*payload.LoginResponse, error) {
	isThere, err := s.siswaRepository.IsSiswaExistsByEmail(request.Email)
	if err != nil {
		return nil, err
	}

	if !isThere {
		return nil, utility.ErrNoDataFound
	}

	siswa, err := s.siswaRepository.Login(request.Email, request.Password)
	if err != nil {
		return nil, err
	}

	tokenString, err := auth.CreateJWTToken(siswa.Email, siswa.KategoriUser, siswa.Id)
	if err != nil {
		return nil, err
	}

	return &payload.LoginResponse{
		Email:   siswa.Email,
		Role:    siswa.KategoriUser,
		IdUser:  siswa.Id,
		IdSiswa: siswa.Siswa.Id,
		Token:   tokenString,
	}, nil
}

func (s *siswaServiceImpl) GetListSiswa(request payload.ListSiswaRequest) (*payload.ListSiswaResponse, error) {
	totalSiswa, err := s.siswaRepository.GetTotalSiswa(request.Nama)
	if err != nil {
		return nil, err
	}

	nextPage, prevPage, totalPages := utility.GetPaginateURL("api/siswa", &request.Page, &request.Limit, totalSiswa)

	listSiswa, err := s.siswaRepository.GetListSiswa(request.Page, request.Limit, request.Nama)
	if err != nil {
		return nil, err
	}

	lenListSiswa := len(listSiswa)
	if lenListSiswa == 0 {
		return nil, utility.ErrNoDataFound
	}

	results := make([]payload.Siswa, 0)
	for i := 0; i < lenListSiswa; i++ {
		siswa := listSiswa[i]
		results = append(results, payload.Siswa{
			Id:                siswa.Id,
			IdUser:            siswa.IdUser,
			Nama:              siswa.Nama,
			NamaInstansi:      siswa.NamaInstansi,
			TingkatPendidikan: siswa.TingkatPendidikan,
			Alamat:            siswa.Alamat,
			NomorTelepon:      siswa.NomorTelepon,
			Email:             siswa.Email,
			TanggalLahir:      siswa.TanggalLahir,
			NomorRekening:     siswa.NomorRekening,
			NamaBank:          siswa.NamaBank,
		})
	}

	return &payload.ListSiswaResponse{
		Data: results,
		PaginateInfo: payload.PaginateInfo{
			NextPage:   nextPage,
			PrevPage:   prevPage,
			TotalPages: totalPages,
		},
	}, nil
}

func (s *siswaServiceImpl) RegisterSiswa(request payload.RegisterSiswaRequest) (*payload.LoginResponse, error) {
	siswa, err := s.siswaRepository.RegisterSiswa(&entity.SiswaDetail{
		Nama:              request.Nama,
		NamaInstansi:      request.NamaInstansi,
		TingkatPendidikan: request.TingkatPendidikan,
		Alamat:            request.Alamat,
		NomorTelepon:      request.NomorTelepon,
		Email:             request.Email,
		TanggalLahir:      request.TanggalLahir,
		NomorRekening:     request.NomorRekening,
		NamaBank:          request.NamaBank,
	}, &entity.Siswa{
		Email:    request.Email,
		Password: request.Password,
	})

	if err != nil {
		return nil, err
	}

	tokenString, err := auth.CreateJWTToken(siswa.Email, "SISWA", siswa.Id)
	if err != nil {
		return nil, err
	}

	return &payload.LoginResponse{
		Email:  siswa.Email,
		Role:   "SISWA",
		IdUser: siswa.Id,
		Token:  tokenString,
	}, nil
}
