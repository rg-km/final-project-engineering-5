package module

import (
	"database/sql"
	"FinalProject/api/repository"
	repositoryImpl "FinalProject/impl/repository"
)

type dataModuleImpl struct {
	siswaRepository repository.SiswaRepository
}

func NewDataModuleImpl(db *sql.DB) *dataModuleImpl {
	return &dataModuleImpl{
		siswaRepository: repositoryImpl.NewSiswaRepositoryImpl(db),
	}
}

func (d *dataModuleImpl) GetSiswaRepository() repository.SiswaRepository {
	return d.siswaRepository
}
