package service

import "FinalProject/payload"

type SiswaService interface {
	GetSiswaById(id int) (*payload.SiswaDetailResponse, error)
	Login(request payload.LoginRequest) (*payload.LoginResponse, error)
	GetListSiswa(request payload.ListSiswaRequest) (*payload.ListSiswaResponse, error)
	RegisterSiswa(request payload.RegisterSiswaRequest) (*payload.LoginResponse, error)
}
