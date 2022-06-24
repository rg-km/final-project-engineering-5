package service

import (
	"FinalProject/payload"
)

type BeasiswaSiswaService interface {
	UpdateStatusBeasiswa(request payload.BeasiswaSiswaStatusUpdateRequest, id int) (*payload.BeasiswaSiswaStatusUpdateResponse, error)
	ApplyBeasiswa(request payload.BeasiswaSiswaApplyRequest) (*payload.BeasiswaSiswaApplyResponse, error)
	GetListBeassiwaSiswaByIdMitra(request payload.ListBeasiswaSiswaByMitraIdRequest) (*payload.ListBeasiswaSiswaByMitraIdResponse, error)
}
