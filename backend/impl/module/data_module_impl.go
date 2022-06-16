package module

import (
	"FinalProject/api/repository"
	repositoryImpl "FinalProject/impl/repository"
	"database/sql"
)

type dataModuleImpl struct {
	siswaRepository    repository.SiswaRepository
	mitraRepository    repository.MitraRepository
	beasiswaRepository repository.BeasiswaRepository
}

func NewDataModuleImpl(db *sql.DB) *dataModuleImpl {
	return &dataModuleImpl{
		siswaRepository:    repositoryImpl.NewSiswaRepositoryImpl(db),
		mitraRepository:    repositoryImpl.NewMitraRepositoryImpl(db),
		beasiswaRepository: repositoryImpl.NewBeasiswaRepositoryImpl(db),
	}
}

func (d *dataModuleImpl) GetSiswaRepository() repository.SiswaRepository {
	return d.siswaRepository
}

func (d *dataModuleImpl) GetMitraRepository() repository.MitraRepository {
	return d.mitraRepository
}

func (d *dataModuleImpl) GetBeasiswaRepository() repository.BeasiswaRepository {
	return d.beasiswaRepository
}
