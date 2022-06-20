package module

import (
	"FinalProject/api/module"
	"FinalProject/api/service"
	serviceImpl "FinalProject/impl/service"
)

type serviceModuleImpl struct {
	siswaService    service.SiswaService
	mitraService    service.MitraService
	beasiswaService service.BeasiswaService
	beasiswaSiswaService service.BeasiswaSiswaService
}

func NewServiceModuleImpl(dataModule module.DataModule) *serviceModuleImpl {
	return &serviceModuleImpl{
		siswaService: serviceImpl.NewSiswaServiceImpl(
			dataModule.GetSiswaRepository(),
		),
		mitraService: serviceImpl.NewMitraServiceImpl(
			dataModule.GetMitraRepository(),
		),
		beasiswaService: serviceImpl.NewBeasiswaServiceImpl(
			dataModule.GetBeasiswaRepository(),
		),
		beasiswaSiswaService: serviceImpl.NewBeasiswaSiswaServiceImpl(
			dataModule.GetBeasiswaSiswaRepository(),
		),
	}
}

func (s *serviceModuleImpl) GetSiswaService() service.SiswaService {
	return s.siswaService
}

func (s *serviceModuleImpl) GetMitraService() service.MitraService {
	return s.mitraService
}

func (s *serviceModuleImpl) GetBeasiswaService() service.BeasiswaService {
	return s.beasiswaService
}

func (s *serviceModuleImpl) GetBeasiswaSiswaService() service.BeasiswaSiswaService {
	return s.beasiswaSiswaService
}
