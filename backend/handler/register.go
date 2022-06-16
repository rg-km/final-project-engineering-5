package handler

import (
	"FinalProject/payload"
	"FinalProject/utility"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) registerHandler(r *gin.Engine) {
	baseEndpoints := r.Group("/api")

	baseEndpoints.POST("/siswa/login", h.handleLoginSiswa)
	baseEndpoints.GET("/siswa", h.handleGetListSiswa)

	baseEndpoints.POST("/mitra/login", h.handleLoginMitra)

	baseEndpoints.GET("beasiswa/:id", h.handleGetBeasiswaById)
}

func (h *handler) handleLoginSiswa(c *gin.Context) {
	request := payload.LoginRequest{}

	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
			Error   string `json:"error"`
		}{Message: err.Error(), Error: utility.ErrBadRequest.Error()})
		return
	}

	response, err := h.siswaService.Login(request)
	if err != nil {
		if err == utility.ErrNoDataFound {
			c.JSON(http.StatusNotFound, struct {
				Message string `json:"message"`
				Error   string `json:"error"`
			}{Message: "Tidak dapat melayani permintaan anda saat ini.", Error: err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, struct {
			Message string `json:"message"`
			Error   string `json:"error"`
		}{Message: "Tidak dapat melayani permintaan anda saat ini.", Error: err.Error()})
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
			Error   string `json:"error"`
		}{Message: err.Error(), Error: utility.ErrBadRequest.Error()})
		return
	}

	response, err := h.mitraService.Login(request)
	if err != nil {
		if err == utility.ErrNoDataFound {
			c.JSON(http.StatusNotFound, struct {
				Message string `json:"message"`
				Error   string `json:"error"`
			}{Message: "Tidak dapat melayani permintaan anda saat ini.", Error: err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, struct {
			Message string `json:"message"`
			Error   string `json:"error"`
		}{Message: "Tidak dapat melayani permintaan anda saat ini.", Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
	return
}

func (h *handler) handleGetListSiswa(c *gin.Context) {
	request := payload.ListSiswaRequest{}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
			Error   string `json:"error"`
		}{Message: err.Error(), Error: utility.ErrBadRequest.Error()})
		return
	}
	request.Page = page

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
			Error   string `json:"error"`
		}{Message: err.Error(), Error: utility.ErrBadRequest.Error()})
		return
	}
	request.Limit = limit

	request.Nama = c.Query("nama")

	response, err := h.siswaService.GetListSiswa(request)
	if err != nil {
		if err == utility.ErrNoDataFound {
			c.JSON(http.StatusNotFound, struct {
				Message string `json:"message"`
				Error   string `json:"error"`
			}{Message: "Tidak ada data.", Error: err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, struct {
			Message string `json:"message"`
			Error   string `json:"error"`
		}{Message: "Tidak dapat melayani permintaan anda saat ini.", Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
	return
}

func (h *handler) handleGetBeasiswaById(c *gin.Context) {
	id := c.Param("id")
	response, err := h.beasiswaService.GetBeasiswaById(id)
	if err != nil {
		if err == utility.ErrNoDataFound {
			c.JSON(http.StatusNotFound, struct {
				Message string `json:"message"`
				Error   string `json:"error"`
			}{Message: "Tidak ada data.", Error: err.Error()})
			return
		}

	}

	log.Println(response)

	c.JSON(http.StatusOK, response)
}
