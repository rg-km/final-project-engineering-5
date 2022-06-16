package service

import "FinalProject/payload"

type BeasiswaService interface {
	GetListBeasiswa(request payload.ListBeasiswaRequest) (*payload.ListBeasiswaResponse, error)
}