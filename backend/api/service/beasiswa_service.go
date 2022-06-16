package service

import (
	"FinalProject/payload"
)

type BeasiswaService interface {
	GetBeasiswaById(id string) (*payload.BeasiswaResponse, error)
}
