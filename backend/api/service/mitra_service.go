package service

import "FinalProject/payload"

type MitraService interface {
	Login(request payload.LoginRequest) (*payload.LoginResponse, error)
	RegisterMitra(request payload.RegisterMitraDetailRequest) (*payload.LoginResponse, error)
}
