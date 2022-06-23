package service

import (
	"FinalProject/payload"
)

type BeasiswaService interface {
	GetBeasiswaById(id int) (*payload.BeasiswaResponse, error)
	GetListBeasiswa(request payload.ListBeasiswaRequest) (*payload.ListBeasiswaResponse, error)
	CreateBeasiswa(request payload.Beasiswa) (*payload.Beasiswa, error)
	UpdateBeasiswa(request payload.Beasiswa, id int) (*payload.BeasiswaResponse, error)
	DeleteBeasiswa(id int) (*payload.DeleteBeasiswaResponse, error)
}
