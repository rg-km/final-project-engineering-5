package module

import (
	"FinalProject/api/repository"
	repositoryImpl "FinalProject/impl/repository"
	"database/sql"
	"sync"
)

type dataModuleImpl struct {
	siswaRepository    repository.SiswaRepository
	mitraRepository    repository.MitraRepository
	beasiswaRepository repository.BeasiswaRepository
	beasiswaSiswaRepository repository.BeasiswaSiswaRepostiroy
}

func NewDataModuleImpl(db *sql.DB, mu *sync.Mutex) *dataModuleImpl {
	return &dataModuleImpl{
		siswaRepository:    repositoryImpl.NewSiswaRepositoryImpl(db),
		mitraRepository:    repositoryImpl.NewMitraRepositoryImpl(db),
		beasiswaRepository: repositoryImpl.NewBeasiswaRepositoryImpl(db, mu),
		beasiswaSiswaRepository: repositoryImpl.NewBeasiswaSiswaRepositoryImpl(db),
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

func (d *dataModuleImpl) GetBeasiswaSiswaRepository() repository.BeasiswaSiswaRepostiroy {
	return d.beasiswaSiswaRepository
}
