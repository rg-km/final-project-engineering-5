package service

import (
	"FinalProject/api/repository"
	"FinalProject/payload"
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
