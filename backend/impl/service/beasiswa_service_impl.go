package service

import (
	"FinalProject/api/repository"
	"FinalProject/entity"
	"FinalProject/payload"
	"FinalProject/utility"
	"log"
)

type beasiswaServiceImpl struct {
	beasiswaRepository repository.BeasiswaRepository
}

func NewBeasiswaServiceImpl(beasiswaRepository repository.BeasiswaRepository) *beasiswaServiceImpl {
	return &beasiswaServiceImpl{
		beasiswaRepository: beasiswaRepository,
	}
}

func (b *beasiswaServiceImpl) GetBeasiswaById(id string) (*payload.BeasiswaResponse, error) {
	// beasiswa, err := b.beasiswaRepository.GetBeasiswaById(id)
	// if err != nil {
	// 	return nil, err
	// }

	// return beasiswa, nil

	beasiswa, err := b.beasiswaRepository.GetBeasiswaById(id)
	if err != nil {
		return nil, err
	}

	results := make([]payload.Beasiswa, 0)
	for _, beasiswaItem := range beasiswa {
		results = append(results, payload.Beasiswa{
			Id:               beasiswaItem.Id,
			IdMitra:          beasiswaItem.IdMitra,
			JudulBeasiswa:    beasiswaItem.JudulBeasiswa,
			Deskripsi:        beasiswaItem.Deskripsi,
			TanggalPembukaan: beasiswaItem.TanggalPembukaan,
			TanggalPenutupan: beasiswaItem.TanggalPenutupan,
			Benefits:         beasiswaItem.Benefits,
		})
	}

	return &payload.BeasiswaResponse{
		Data: results,
	}, nil

}

func (s *beasiswaServiceImpl) GetListBeasiswa(request payload.ListBeasiswaRequest) (*payload.ListBeasiswaResponse, error) {
	totalBeasiswa, err := s.beasiswaRepository.GetTotalBeasiswa(request.Nama)
	if err != nil {
		return nil, err
	}
	log.Println("totalBeasiswa:", totalBeasiswa)

	nextPage, prevPage, totalPages := utility.GetPaginateURL("api/beasiswa", &request.Page, &request.Limit, totalBeasiswa)

	listBeasiswa, err := s.beasiswaRepository.GetListBeasiswa(request.Page, request.Limit, request.Nama)
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

func (s *beasiswaServiceImpl) CreateBeasiswa(request payload.Beasiswa) (*payload.Beasiswa, error) {
	beasiswa, err := s.beasiswaRepository.CreateBeasiswa(&entity.Beasiswa{
		Id:               request.Id,
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

	log.Println(beasiswa)
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
