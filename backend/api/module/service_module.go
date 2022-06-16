package module

import (
	"FinalProject/api/service"
)

type ServiceModule interface {
	GetSiswaService() service.SiswaService
	GetMitraService() service.MitraService
	GetBeasiswaService() service.BeasiswaService
}