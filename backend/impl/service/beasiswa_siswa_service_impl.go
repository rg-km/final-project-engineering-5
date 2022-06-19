package service

import (
	"FinalProject/api/repository"
	"FinalProject/entity"
	"FinalProject/payload"
	"FinalProject/utility"
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
