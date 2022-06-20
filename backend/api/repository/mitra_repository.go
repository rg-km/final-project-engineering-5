package repository

import "FinalProject/entity"

type MitraRepository interface {
	IsMitraExistsByEmail(email string) (bool, error)
	Login(email string, password string) (*entity.Mitra, error)
	RegisterMitra(mitra *entity.MitraDetail, user *entity.Mitra) (*entity.MitraDetail, error)
}
