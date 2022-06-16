package module

import (
	"FinalProject/api/repository"
)

type DataModule interface {
	GetSiswaRepository() repository.SiswaRepository
	GetMitraRepository() repository.MitraRepository
	GetBeasiswaRepository() repository.BeasiswaRepository
}