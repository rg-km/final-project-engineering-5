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

	baseEndpoints.GET("/beasiswa", h.handleGetListBeasiswa)
	baseEndpoints.POST("/add/beasiswa", h.handleCreateBeasiswa)
	baseEndpoints.GET("/beasiswa/:id", h.handleGetBeasiswaById)
	baseEndpoints.PUT("/beasiswa/:id", h.handleUpdateBeasiswa)
	
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
	requestId := c.Param("id")
	
	id, err := strconv.Atoi(requestId)
	if err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
			Error string `json:"error"`
		}{Message: "Pastikan id yang valid.", Error: utility.ErrBadRequest.Error()})
		return
	}
	log.Println("id:",id)

	response, err := h.beasiswaService.GetBeasiswaById(id)
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

	log.Println(response)

	c.JSON(http.StatusOK, response)
}

func (h *handler) handleGetListBeasiswa(c *gin.Context) {
	request := payload.ListBeasiswaRequest{}

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

	log.Println(page, limit, request.Nama, request)

	response, err := h.beasiswaService.GetListBeasiswa(request)
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

func (h *handler) handleCreateBeasiswa(c *gin.Context) {
	request := payload.Beasiswa{}

	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
			Error   string `json:"error"`
		}{Message: err.Error(), Error: utility.ErrBadRequest.Error()})
		return
	}

	response, err := h.beasiswaService.CreateBeasiswa(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, struct {
			Message string `json:"message"`
			Error   string `json:"error"`
		}{Message: "Tidak dapat melayani permintaan anda saat ini.", Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
	return
}

func (h *handler) handleUpdateBeasiswa(c *gin.Context) {
	request := payload.Beasiswa{}
	if err := c.Bind(&request); err != nil {
		log.Println("masuk?")
		c.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
			Error   string `json:"error"`
		}{Message: err.Error(), Error: utility.ErrBadRequest.Error()})
		return
	}

	requestId := c.Param("id")
	id, err := strconv.Atoi(requestId)
	if err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
			Error string `json:"error"`
		}{Message: "Pastikan id yang valid.", Error: utility.ErrBadRequest.Error()})
		return
	}

	response, err := h.beasiswaService.UpdateBeasiswa(request, id)
	if err != nil {
		if err == utility.ErrBadRequest {
			c.JSON(http.StatusBadRequest, struct {
				Message string `json:"message"`
				Error string `json:"error"`
			}{Message: "Pastikan data valid.", Error: utility.ErrBadRequest.Error()})
			return
		}

		if err == utility.ErrNoDataFound {
			c.JSON(http.StatusBadRequest, struct {
				Message string `json:"message"`
				Error string `json:"error"`
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
