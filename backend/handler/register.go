package handler

import (
	"FinalProject/middleware"
	"FinalProject/payload"
	"FinalProject/utility"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) registerHandler(r *gin.Engine) {
	baseEndpoints := r.Group("/api")

	baseEndpoints.POST("/siswa/login", h.handleLoginSiswa)
	baseEndpoints.GET("/siswa", middleware.ValidateMitraRole(), h.handleGetListSiswa)
	baseEndpoints.POST("/siswa/signup", h.handleRegisterSiswa)
	
	baseEndpoints.POST("/mitra/login", h.handleLoginMitra)
	baseEndpoints.POST("/mitra/signup", h.handleRegisterMitra)

	baseEndpoints.GET("/beasiswa", h.handleGetListBeasiswa)
	baseEndpoints.POST("beasiswa", middleware.ValidateMitraRole(), h.handleCreateBeasiswa)
	baseEndpoints.GET("/beasiswa/:id", h.handleGetBeasiswaById)
	baseEndpoints.PUT("/beasiswa/:id", middleware.ValidateMitraRole(), h.handleUpdateBeasiswa)

	baseEndpoints.PUT("/beasiswa-siswa/:id", middleware.ValidateMitraRole(), h.handleUpdateStatusBeasiswa)
}

func (h *handler) handleLoginSiswa(c *gin.Context) {
	email, password, ok := c.Request.BasicAuth()
	if !ok {
		c.JSON(http.StatusUnauthorized, struct {
			Message string `json:"message"`
			Error string `json:"erorr"`
		}{Message: "Invalid Auth", Error: utility.ErrUnauthorized.Error()})
		return
	}

	response, err := h.siswaService.Login(payload.LoginRequest{
		Email: email,
		Password: password,
	})
	if err != nil {
		if err == utility.ErrNoDataFound {
			c.JSON(http.StatusNotFound, struct {
				Message string `json:"message"`
				Error   string `json:"error"`
			}{Message: "User belum terdaftar.", Error: err.Error()})
			return
		}

		if err == utility.ErrUnauthorized {
			c.JSON(http.StatusUnauthorized, struct {
				Message string `json:"message"`
				Error string `json:"error"`
			}{Message: "Email atau Password tidak valid.", Error: err.Error()})
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
	email, password, ok := c.Request.BasicAuth()
	if !ok {
		c.JSON(http.StatusUnauthorized, struct {
			Message string `json:"message"`
			Error string `json:"erorr"`
		}{Message: "Invalid Auth", Error: utility.ErrUnauthorized.Error()})
		return
	}

	response, err := h.mitraService.Login(payload.LoginRequest{
		Email: email,
		Password: password,
	})
	if err != nil {
		if err == utility.ErrNoDataFound {
			c.JSON(http.StatusNotFound, struct {
				Message string `json:"message"`
				Error   string `json:"error"`
			}{Message: "User belum terdaftar.", Error: err.Error()})
			return
		}

		if err == utility.ErrUnauthorized {
			c.JSON(http.StatusUnauthorized, struct {
				Message string `json:"message"`
				Error string `json:"error"`
			}{Message: "Email atau Password tidak valid.", Error: err.Error()})
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
			Error   string `json:"error"`
		}{Message: "Pastikan id yang valid.", Error: utility.ErrBadRequest.Error()})
		return
	}

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
			Error   string `json:"error"`
		}{Message: "Pastikan id yang valid.", Error: utility.ErrBadRequest.Error()})
		return
	}

	response, err := h.beasiswaService.UpdateBeasiswa(request, id)
	if err != nil {
		if err == utility.ErrBadRequest {
			c.JSON(http.StatusBadRequest, struct {
				Message string `json:"message"`
				Error   string `json:"error"`
			}{Message: "Pastikan data valid.", Error: utility.ErrBadRequest.Error()})
			return
		}

		if err == utility.ErrNoDataFound {
			c.JSON(http.StatusBadRequest, struct {
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
func (h *handler) handleRegisterMitra(c *gin.Context) {
	request := payload.MitraDetail{}
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
			Error   string `json:"error"`
		}{Message: err.Error(), Error: utility.ErrBadRequest.Error()})
	}

	response, err := h.mitraService.RegisterMitra(request)
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

func (h *handler) handleRegisterSiswa(c *gin.Context) {
	request := payload.Siswa{}
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
			Error   string `json:"error"`
		}{Message: err.Error(), Error: utility.ErrBadRequest.Error()})
		return
	}

	response, err := h.siswaService.RegisterSiswa(request)
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


func (h *handler) handleUpdateStatusBeasiswa(c *gin.Context) {
	request := payload.BeasiswaSiswaStatusUpdateRequest{}
	if err := c.Bind(&request); err != nil {
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

	response, err := h.beasiswaSiswaService.UpdateStatusBeasiswa(request, id)
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
