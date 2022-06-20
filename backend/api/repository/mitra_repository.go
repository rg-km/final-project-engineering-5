package repository

import "FinalProject/entity"

type MitraRepository interface {
	Login(email string, password string) (*entity.Mitra, error)
	RegisterMitra(mitra *entity.MitraDetail, user *entity.Mitra) (*entity.MitraDetail, error)
}
