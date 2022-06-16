package service

import (
	"FinalProject/api/repository"
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
			Id: beasiswa.Id,
			IdMitra: beasiswa.IdMitra,
			NamaMitra: beasiswa.NamaMitra,
			JudulBeasiwa: beasiswa.JudulBeasiwa,
			Benefist: beasiswa.Benefist,
			Deskripsi: beasiswa.Deskripsi,
			TanggalPembukaan: beasiswa.TanggalPembukaan,
			TanggalPenutupan: beasiswa.TanggalPenutupan,
		})
	}

	return &payload.ListBeasiswaResponse{
		Data: results,
		PaginateInfo: payload.PaginateInfo{
			NextPage: nextPage,
			PrevPage: prevPage,
			TotalPages: totalPages,
		},
	}, nil
}