package handler

import (
	"FinalProject/payload"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) registerHandler(r *gin.Engine) {
	baseEndpoints := r.Group("/api")

	baseEndpoints.POST("/siswa/login", h.handleLoginSiswa)
	baseEndpoints.POST("/mitra/login", h.handleLoginMitra)
}

func (h *handler) handleLoginSiswa(c *gin.Context) {
	request := payload.LoginRequest{}
	
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
		}{Message: err.Error()})
		return
	}

	response, err := h.siswaService.Login(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
		}{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
	return
}

func (h *handler) handleLoginMitra(c *gin.Context) {
	request := payload.LoginRequest{}

	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
		}{Message: err.Error()})
		return
	}

	response, err := h.mitraService.Login(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
		}{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}