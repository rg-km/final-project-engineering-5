package module

import (
	"FinalProject/api/service"
	"FinalProject/api/module"
	serviceImpl "FinalProject/impl/service"
)

type serviceModuleImpl struct {
	siswaService service.SiswaService
}

func NewServiceModuleImpl(dataModule module.DataModule) *serviceModuleImpl {
	return &serviceModuleImpl{
		siswaService: serviceImpl.NewSiswaServiceImpl(
			dataModule.GetSiswaRepository(),
		),
	}
}

func (s *serviceModuleImpl) GetSiswaService() service.SiswaService {
	return s.siswaService
}