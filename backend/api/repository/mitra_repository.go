package repository

import "FinalProject/entity"

type MitraRepository interface {
	Login(email string, password string) (*entity.Mitra, error)
}