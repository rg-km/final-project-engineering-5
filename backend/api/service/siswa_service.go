package service

import "FinalProject/payload"

type SiswaService interface {
	Login(request payload.LoginRequest) (*payload.LoginResponse, error)
	GetListSiswa(request payload.ListSiswaRequest) (*payload.ListSiswaResponse, error)
	RegisterSiswa(request payload.RegisterSiswaRequest) (*payload.LoginResponse, error)
}
