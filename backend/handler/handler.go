package handler

import (
	"FinalProject/api/module"
	"FinalProject/api/service"
	"FinalProject/middleware"

	"github.com/gin-gonic/gin"
)

type handler struct {
	siswaService    service.SiswaService
	mitraService    service.MitraService
	beasiswaService service.BeasiswaService
}

func StartHandler(serviceModule module.ServiceModule) {
	handler := handler{
		siswaService:    serviceModule.GetSiswaService(),
		mitraService:    serviceModule.GetMitraService(),
		beasiswaService: serviceModule.GetBeasiswaService(),
	}

	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	handler.registerHandler(router)

	router.Run(":8080")
}
