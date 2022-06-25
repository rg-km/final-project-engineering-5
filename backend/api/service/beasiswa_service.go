package service

import (
	"FinalProject/payload"
)

type BeasiswaService interface {
	GetBeasiswaById(id int) (*payload.BeasiswaResponse, error)
	GetListBeasiswa(request payload.ListBeasiswaRequest) (*payload.ListBeasiswaResponse, error)
	CreateBeasiswa(request payload.CreateBeasiswaRequest) (*payload.Beasiswa, error)
	UpdateBeasiswa(request payload.UpdateBeasiswaRequest, id int) (*payload.BeasiswaResponse, error)
	DeleteBeasiswa(id int) (*payload.DeleteBeasiswaResponse, error)
}
