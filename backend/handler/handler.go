package handler

import (
	"FinalProject/api/service"
	"FinalProject/api/module"

	"github.com/gin-gonic/gin"
)

type handler struct {
	siswaService service.SiswaService
}

func StartHandler(serviceModule module.ServiceModule) {
	handler := handler{
		siswaService: serviceModule.GetSiswaService(),
	}

	router := gin.Default()

	handler.registerHandler(router)

	router.Run(":8080")
}
