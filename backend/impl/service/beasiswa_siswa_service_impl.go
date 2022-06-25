package service

import (
	"FinalProject/api/repository"
	"FinalProject/entity"
	"FinalProject/payload"
	"FinalProject/utility"
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
	mitraRepository repository.MitraRepository
}

func NewBeasiswaSiswaServiceImpl(
	beasiswaSiswaRepository repository.BeasiswaSiswaRepostiroy,
	beasiswaRepository repository.BeasiswaRepository,
	siswaRepository repository.SiswaRepository,
	mitraRepository repository.MitraRepository) *beasiswaSiswaServiceImpl {
	return &beasiswaSiswaServiceImpl{
		beasiswaSiswaRepository: beasiswaSiswaRepository,
		beasiswaRepository: beasiswaRepository,
		siswaRepeository: siswaRepository,
		mitraRepository: mitraRepository,
	}
}

func (b *beasiswaSiswaServiceImpl) GetBeasiswaSiswaById(id int) (*payload.BeasiswaSiswa, error) {
	isThere, err := b.beasiswaSiswaRepository.IsBeasiswaSiswaExistsById(id)
	if err != nil {
		return nil, err
	}

	if !isThere {
		return nil, utility.ErrNoDataFound
	}

	beasiswaSiswa, err := b.beasiswaSiswaRepository.GetBeasiswaSiswaById(id)
	if err != nil {
		return nil, err
	}

	return &payload.BeasiswaSiswa{
		Id: beasiswaSiswa.Id,
		IdSiswa: beasiswaSiswa.IdSiswa,
		NamaSiswa: beasiswaSiswa.NamaSiswa,
		IdBeasiswa: beasiswaSiswa.IdBeasiswa,
		NamaBeasiswa: beasiswaSiswa.NamaBeasiswa,
		IdMitra: beasiswaSiswa.IdMitra,
		NamaMitra: beasiswaSiswa.NamaMitra,
		Status: beasiswaSiswa.Status,
		TanggalDaftar: beasiswaSiswa.TanggalDaftar,
	}, nil
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

func (b *beasiswaSiswaServiceImpl) GetListBeassiwaSiswaByIdMitra(request payload.ListBeasiswaSiswaByMitraIdRequest) (*payload.ListBeasiswaSiswaByMitraIdResponse, error) {
	isThere, err := b.mitraRepository.IsMitraExistsById(request.IdMitra)
	if err != nil {
		return nil, err
	}

	if !isThere {
		return nil, utility.ErrNoDataFound
	}

	totalBeasiswaSiswa, err := b.beasiswaSiswaRepository.GetTotalBeasiswaSiswa(request.Nama)
	if err != nil {
		return nil, err
	}
	nextPage, prevPage, totalPages := utility.GetPaginateURL("api/beasiswa-siswa", &request.Page, &request.Limit, totalBeasiswaSiswa)

	listBeasiswaSiswa, err := b.beasiswaSiswaRepository.GetListBeasiswaSiswaByIdMitra(request.IdMitra, request.Page, request.Limit, request.Nama)
	if err != nil {
		return nil, err
	}

	lenListBeasiswSiswa := len(listBeasiswaSiswa)
	if lenListBeasiswSiswa == 0 {
		return nil, utility.ErrNoDataFound
	}

	results := make([]payload.BeasiswaSiswa, 0)
	for i := 0; i < lenListBeasiswSiswa; i++ {
		beasiswaSiswa := listBeasiswaSiswa[i]
		results = append(results, payload.BeasiswaSiswa{
			Id: beasiswaSiswa.Id,
			IdSiswa: beasiswaSiswa.IdSiswa,
			NamaSiswa: beasiswaSiswa.NamaSiswa,
			IdBeasiswa: beasiswaSiswa.IdBeasiswa,
			NamaBeasiswa: beasiswaSiswa.NamaBeasiswa,
			IdMitra: beasiswaSiswa.IdMitra,
			NamaMitra: beasiswaSiswa.NamaMitra,
			Status: beasiswaSiswa.Status,
			TanggalDaftar: beasiswaSiswa.TanggalDaftar,
		})
	}

	return &payload.ListBeasiswaSiswaByMitraIdResponse{
		Data: results,
		PaginateInfo: payload.PaginateInfo{
			NextPage: nextPage,
			PrevPage: prevPage,
			TotalPages: totalPages,
		},
	}, nil
}
