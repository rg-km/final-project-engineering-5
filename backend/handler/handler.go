package handler

import (
	"FinalProject/api/service"
	"FinalProject/api/module"

	"github.com/gin-gonic/gin"
)

type handler struct {
	siswaService service.SiswaService
	mitraService service.MitraService
	beasiswaService service.BeasiswaService
}

func StartHandler(serviceModule module.ServiceModule) {
	handler := handler{
		siswaService: serviceModule.GetSiswaService(),
		mitraService: serviceModule.GetMitraService(),
		beasiswaService: serviceModule.GetBeasiswaService(),
	}

	router := gin.Default()

	handler.registerHandler(router)

	router.Run(":8080")
}
