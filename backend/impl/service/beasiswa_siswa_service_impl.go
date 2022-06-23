package service

import (
	"FinalProject/api/repository"
	"FinalProject/entity"
	"FinalProject/payload"
	"FinalProject/utility"
	"log"
	"strings"
)

const (
	STATUS_BEASISWA_DITERIMA = "DITERIMA"
	STATUS_BEASISWA_DIPROSES = "DIPROSES"
	STATUS_BEASISWA_DITOLAK  = "DITOLAK"
)

type beasiswaSiswaServiceImpl struct {
	beasiswaSiswaRepository repository.BeasiswaSiswaRepostiroy
	beasiswaRepository repository.BeasiswaRepository
	siswaRepeository repository.SiswaRepository
}

func NewBeasiswaSiswaServiceImpl(
	beasiswaSiswaRepository repository.BeasiswaSiswaRepostiroy,
	beasiswaRepository repository.BeasiswaRepository,
	siswaRepository repository.SiswaRepository) *beasiswaSiswaServiceImpl {
	return &beasiswaSiswaServiceImpl{
		beasiswaSiswaRepository: beasiswaSiswaRepository,
		beasiswaRepository: beasiswaRepository,
		siswaRepeository: siswaRepository,
	}
}

func (b *beasiswaSiswaServiceImpl) UpdateStatusBeasiswa(
	request payload.BeasiswaSiswaStatusUpdateRequest, id int) (*payload.BeasiswaSiswaStatusUpdateResponse, error) {
	if request.Id != id {
		return nil, utility.ErrBadRequest
	}

	isThere, err := b.beasiswaSiswaRepository.IsBeasiswaSiswaExistsById(id)
	if err != nil {
		return nil, err
	}

	if !isThere {
		return nil, utility.ErrNoDataFound
	}

	if strings.Compare(STATUS_BEASISWA_DITERIMA, request.Status) != 0 &&
		strings.Compare(STATUS_BEASISWA_DIPROSES, request.Status) != 0 &&
		strings.Compare(STATUS_BEASISWA_DITOLAK, request.Status) != 0 {
		return nil, utility.ErrBadRequest
	}

	updatedStatusBeasiswa, err := b.beasiswaSiswaRepository.UpdateStatusBeasiswa(entity.BeasiswaSiswaStatusUpdate{
		Id:         request.Id,
		IdSiswa:    request.IdSiswa,
		IdBeasiswa: request.IdBeasiswa,
		Status:     request.Status,
	}, id)
	if err != nil {
		return nil, err
	}

	return &payload.BeasiswaSiswaStatusUpdateResponse{
		Message: "Berhasil melakukan update status beasiswa",
		BeasiswaSiswa: payload.BeasiswaSiswa{
			Id:            updatedStatusBeasiswa.Id,
			IdSiswa:       updatedStatusBeasiswa.IdSiswa,
			NamaSiswa:     updatedStatusBeasiswa.NamaSiswa,
			IdBeasiswa:    updatedStatusBeasiswa.IdBeasiswa,
			NamaBeasiswa:  updatedStatusBeasiswa.NamaBeasiswa,
			IdMitra:       updatedStatusBeasiswa.IdMitra,
			NamaMitra:     updatedStatusBeasiswa.NamaMitra,
			Status:        updatedStatusBeasiswa.Status,
			TanggalDaftar: updatedStatusBeasiswa.TanggalDaftar,
		},
	}, nil
}

func (b *beasiswaSiswaServiceImpl) ApplyBeasiswa(request payload.BeasiswaSiswaApplyRequest) (*payload.BeasiswaSiswaApplyResponse, error) {
	isThereBeasiswa, err := b.beasiswaRepository.IsBeasiswaExistsById(request.IdBeasiswa)
	if err != nil {
		return nil, err
	}
	if !isThereBeasiswa {
		return nil, utility.ErrNoDataFound
	}


	isThereUser, err := b.siswaRepeository.IsSiswaExistsById(request.IdSiswa)
	if err != nil {
		return nil, err
	}
	if !isThereUser {
		return nil, utility.ErrNoDataFound
	}

	beasiswaSiswa, err := b.beasiswaSiswaRepository.ApplyBeasiswa(request.IdBeasiswa, request.IdSiswa)

	if err != nil {
		return nil, err
	}
	log.Println("beasiswaSiswa:", beasiswaSiswa)

	return &payload.BeasiswaSiswaApplyResponse{
		Message: "Berhasil melakukan pendaftaran beasiswa",
		BeasiswaSiswa: payload.BeasiswaSiswa{
			Id:            beasiswaSiswa.Id,
			IdSiswa:       beasiswaSiswa.IdSiswa,
			NamaSiswa:     beasiswaSiswa.NamaSiswa,
			IdBeasiswa:    beasiswaSiswa.IdBeasiswa,
			NamaBeasiswa:  beasiswaSiswa.NamaBeasiswa,
			IdMitra:       beasiswaSiswa.IdMitra,
			NamaMitra:     beasiswaSiswa.NamaMitra,
			Status:        beasiswaSiswa.Status,
			TanggalDaftar: beasiswaSiswa.TanggalDaftar,
		},
	}, nil

}
