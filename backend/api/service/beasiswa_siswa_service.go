package service

import (
	"FinalProject/payload"
)

type BeasiswaSiswaService interface {
	GetBeasiswaSiswaById(id int) (*payload.BeasiswaSiswa, error)
	UpdateStatusBeasiswa(request payload.BeasiswaSiswaStatusUpdateRequest, id int) (*payload.BeasiswaSiswaStatusUpdateResponse, error)
	ApplyBeasiswa(request payload.BeasiswaSiswaApplyRequest) (*payload.BeasiswaSiswaApplyResponse, error)
	GetListBeassiwaSiswaByIdMitra(request payload.ListBeasiswaSiswaByMitraIdRequest) (*payload.ListBeasiswaSiswaByMitraIdResponse, error)
	GetListBeasiswaSiswaByIdSiswa(id int) (*payload.ListBeasiswaSiswaBySiswaIdResponse, error)
}
