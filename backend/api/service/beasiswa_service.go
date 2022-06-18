package service

import (
	"FinalProject/payload"
)

type BeasiswaService interface {
	GetBeasiswaById(id string) (*payload.BeasiswaResponse, error)
	GetListBeasiswa(request payload.ListBeasiswaRequest) (*payload.ListBeasiswaResponse, error)
	CreateBeasiswa(request payload.Beasiswa) (*payload.Beasiswa, error)
}
