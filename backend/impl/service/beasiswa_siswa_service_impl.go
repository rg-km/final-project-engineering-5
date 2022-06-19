package service

import (
	"FinalProject/api/repository"
	"FinalProject/entity"
	"FinalProject/payload"
	"FinalProject/utility"
	"strings"
)

const (
	STATUS_BEASISWA_DITERIMA 	= "DITERIMA"
	STATUS_BEASISWA_DIPROSES 	= "DIPROSES"
	STATUS_BEASISWA_DITOLAK 	= "DITOLAK"
)

type beasiswaSiswaServiceImpl struct {
	beasiswaSiswaRepository repository.BeasiswaSiswaRepostiroy
}

func NewBeasiswaSiswaServiceImpl(
	beasiswaSiswaRepository repository.BeasiswaSiswaRepostiroy) *beasiswaSiswaServiceImpl {
	return &beasiswaSiswaServiceImpl{
		beasiswaSiswaRepository: beasiswaSiswaRepository,
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
		Id: request.Id,
		IdSiswa: request.IdSiswa,
		IdBeasiswa: request.IdBeasiswa,
		Status: request.Status,
	}, id)
	if err != nil {
		return nil, err
	}

	return &payload.BeasiswaSiswaStatusUpdateResponse{
		Message: "Berhasil melakukan update status beasiswa",
		BeasiswaSiswa: payload.BeasiswaSiswa{
			Id: updatedStatusBeasiswa.Id,
			IdSiswa: updatedStatusBeasiswa.IdSiswa,
			NamaSiswa: updatedStatusBeasiswa.NamaSiswa,
			IdBeasiswa: updatedStatusBeasiswa.IdBeasiswa,
			NamaBeasiswa: updatedStatusBeasiswa.NamaBeasiswa,
			IdMitra: updatedStatusBeasiswa.IdMitra,
			NamaMitra: updatedStatusBeasiswa.NamaMitra,
			Status: updatedStatusBeasiswa.Status,
			TanggalDaftar: updatedStatusBeasiswa.TanggalDaftar,
		},
	}, nil
}
