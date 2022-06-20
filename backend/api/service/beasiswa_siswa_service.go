package service

import (
	"FinalProject/payload"
)

type BeasiswaSiswaService interface {
	UpdateStatusBeasiswa(request payload.BeasiswaSiswaStatusUpdateRequest, id int) (*payload.BeasiswaSiswaStatusUpdateResponse, error)
}
