package service

import (
	"FinalProject/api/repository"
	"FinalProject/entity"
	"FinalProject/payload"
	"FinalProject/utility"
)

type beasiswaServiceImpl struct {
	beasiswaRepository repository.BeasiswaRepository
	mitraRepository    repository.MitraRepository
}

func NewBeasiswaServiceImpl(
	beasiswaRepository repository.BeasiswaRepository,
	mitraRepository repository.MitraRepository) *beasiswaServiceImpl {
	return &beasiswaServiceImpl{
		beasiswaRepository: beasiswaRepository,
		mitraRepository:    mitraRepository,
	}
}

func (b *beasiswaServiceImpl) GetBeasiswaById(id int) (*payload.BeasiswaResponse, error) {
	isExists, err := b.beasiswaRepository.IsBeasiswaExistsById(id)
	if err != nil {
		return nil, err
	}

	if !isExists {
		return nil, utility.ErrNoDataFound
	}

	beasiswa, err := b.beasiswaRepository.GetBeasiswaById(id)
	if err != nil {
		return nil, err
	}

	return &payload.BeasiswaResponse{
		Message: "Berhasil mendapatkan data beasiswa.",
		Beasiswa: payload.Beasiswa{
			Id:               beasiswa.Id,
			IdMitra:          beasiswa.IdMitra,
			JudulBeasiswa:    beasiswa.JudulBeasiswa,
			Deskripsi:        beasiswa.Deskripsi,
			TanggalPembukaan: beasiswa.TanggalPembukaan,
			TanggalPenutupan: beasiswa.TanggalPenutupan,
			Benefits:         beasiswa.Benefits,
		},
	}, nil
}

func (b *beasiswaServiceImpl) GetListBeasiswa(request payload.ListBeasiswaRequest) (*payload.ListBeasiswaResponse, error) {
	totalBeasiswa, err := b.beasiswaRepository.GetTotalBeasiswa(request.Nama)
	if err != nil {
		return nil, err
	}

	nextPage, prevPage, totalPages := utility.GetPaginateURL("api/beasiswa", &request.Page, &request.Limit, totalBeasiswa)

	listBeasiswa, err := b.beasiswaRepository.GetListBeasiswa(request.Page, request.Limit, request.Nama)
	if err != nil {
		return nil, err
	}

	lenListBeasiswa := len(listBeasiswa)
	if lenListBeasiswa == 0 {
		return nil, utility.ErrNoDataFound
	}

	results := make([]payload.Beasiswa, 0)
	for i := 0; i < lenListBeasiswa; i++ {
		beasiswa := listBeasiswa[i]
		results = append(results, payload.Beasiswa{
			Id:               beasiswa.Id,
			IdMitra:          beasiswa.IdMitra,
			JudulBeasiswa:    beasiswa.JudulBeasiswa,
			Benefits:         beasiswa.Benefits,
			Deskripsi:        beasiswa.Deskripsi,
			TanggalPembukaan: beasiswa.TanggalPembukaan,
			TanggalPenutupan: beasiswa.TanggalPenutupan,
		})
	}

	return &payload.ListBeasiswaResponse{
		Data: results,
		PaginateInfo: payload.PaginateInfo{
			NextPage:   nextPage,
			PrevPage:   prevPage,
			TotalPages: totalPages,
		},
	}, nil
}

func (b *beasiswaServiceImpl) CreateBeasiswa(request payload.CreateBeasiswaRequest) (*payload.Beasiswa, error) {
	isThere, err := b.mitraRepository.IsMitraExistsById(request.IdMitra)
	if err != nil {
		return nil, err
	}

	if !isThere {
		return nil, utility.ErrNoDataFound
	}

	beasiswa, err := b.beasiswaRepository.CreateBeasiswa(&entity.Beasiswa{
		IdMitra:          request.IdMitra,
		JudulBeasiswa:    request.JudulBeasiswa,
		Deskripsi:        request.Deskripsi,
		TanggalPembukaan: request.TanggalPembukaan,
		TanggalPenutupan: request.TanggalPenutupan,
		Benefits:         request.Benefits,
	})

	if err != nil {
		return nil, err
	}

	return &payload.Beasiswa{
		Id:               beasiswa.Id,
		IdMitra:          beasiswa.IdMitra,
		JudulBeasiswa:    beasiswa.JudulBeasiswa,
		Deskripsi:        beasiswa.Deskripsi,
		TanggalPembukaan: beasiswa.TanggalPembukaan,
		TanggalPenutupan: beasiswa.TanggalPenutupan,
		Benefits:         beasiswa.Benefits,
	}, nil
}

func (b *beasiswaServiceImpl) UpdateBeasiswa(request payload.UpdateBeasiswaRequest, id int) (*payload.BeasiswaResponse, error) {
	if request.Id != id {
		return nil, utility.ErrBadRequest
	}

	isThere, err := b.beasiswaRepository.IsBeasiswaExistsById(id)
	if err != nil {
		return nil, err
	}

	if !isThere {
		return nil, utility.ErrNoDataFound
	}

	updatedBesiswa, err := b.beasiswaRepository.UpdateBeasiswa(entity.Beasiswa{
		Id:               request.Id,
		IdMitra:          request.IdMitra,
		JudulBeasiswa:    request.JudulBeasiswa,
		Deskripsi:        request.Deskripsi,
		TanggalPembukaan: request.TanggalPembukaan,
		TanggalPenutupan: request.TanggalPenutupan,
		Benefits:         request.Benefits,
	}, id)
	if err != nil {
		return nil, err
	}

	return &payload.BeasiswaResponse{
		Message: "Berhasil melakukan update data beasiswa.",
		Beasiswa: payload.Beasiswa{
			Id:               updatedBesiswa.Id,
			IdMitra:          updatedBesiswa.IdMitra,
			JudulBeasiswa:    updatedBesiswa.JudulBeasiswa,
			Deskripsi:        updatedBesiswa.Deskripsi,
			TanggalPembukaan: updatedBesiswa.TanggalPembukaan,
			TanggalPenutupan: updatedBesiswa.TanggalPenutupan,
			Benefits:         updatedBesiswa.Benefits,
		},
	}, nil
}

func (b *beasiswaServiceImpl) DeleteBeasiswa(id int) (*payload.DeleteBeasiswaResponse, error) {

	isThere, err := b.beasiswaRepository.IsBeasiswaExistsById(id)
	if err != nil {
		return nil, err
	}

	if !isThere {
		return nil, utility.ErrNoDataFound
	}

	if err != nil {
		return nil, err
	}

	b.beasiswaRepository.DeleteBeasiswa(id)

	return &payload.DeleteBeasiswaResponse{
		Message: "Berhasil melakukan Delete data beasiswa."}, nil

}

func (b *beasiswaServiceImpl) GetListBeasiswaByMitraId(request payload.ListBeasiswaByMitraIdRequest) (*payload.ListBeasiswaResponse, error) {
	isThere, err := b.mitraRepository.IsMitraExistsById(request.IdMitra)
	if err != nil {
		return nil, err
	}

	if !isThere {
		return nil, utility.ErrNoDataFound
	}

	totalBeasiswaSiswa, err := b.beasiswaRepository.GetTotalBeasiswa(request.Nama)
	if err != nil {
		return nil, err
	}
	nextPage, prevPage, totalPages := utility.GetPaginateURL("api/mitra/beasiswa", &request.Page, &request.Limit, totalBeasiswaSiswa)

	listBeasiswa, err := b.beasiswaRepository.GetListBeasiswaByMitraId(request.IdMitra, request.Page, request.Limit, request.Nama)
	if err != nil {
		return nil, err
	}

	lenListBeasiswa := len(listBeasiswa)
	if lenListBeasiswa == 0 {
		return nil, utility.ErrNoDataFound
	}

	results := make([]payload.Beasiswa, 0)
	for i := 0; i < lenListBeasiswa; i++ {
		beasiswa := listBeasiswa[i]
		results = append(results, payload.Beasiswa{
			Id:               beasiswa.Id,
			IdMitra:          beasiswa.IdMitra,
			JudulBeasiswa:    beasiswa.JudulBeasiswa,
			Benefits:         beasiswa.Benefits,
			Deskripsi:        beasiswa.Deskripsi,
			TanggalPembukaan: beasiswa.TanggalPembukaan,
			TanggalPenutupan: beasiswa.TanggalPenutupan,
		})
	}

	return &payload.ListBeasiswaResponse{
		Data: results,
		PaginateInfo: payload.PaginateInfo{
			NextPage:   nextPage,
			PrevPage:   prevPage,
			TotalPages: totalPages,
		},
	}, nil
}
