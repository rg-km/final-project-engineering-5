package middleware

import (
	"FinalProject/auth"
	"FinalProject/utility"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const (
	MITRA = "MITRA"
	SISWA = "SISWA"
)

func ValidateSiswaRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := auth.ExtractJwtFromHeader(c.Request)

		if len(tokenString) == 0 {
			c.JSON(http.StatusUnauthorized, struct {
				Message string `json:"message"`
				Error string `json:"error"`
			}{Message: "Token tidak ada.", Error: utility.ErrUnauthorized.Error()})
			c.Abort()
			return
		}

		claims, err := auth.GetClaimsFromJwt(tokenString)
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.JSON(http.StatusUnauthorized, struct {
					Message string `json:"message"`
					Error string `json:"error"`
				}{Message: "Terjadi kesalahan saat autentikasi.", Error: err.Error()})
				c.Abort()
				return
			}

			c.JSON(http.StatusBadRequest, struct {
				Message string `json:"message"`
				Error string `json:"error"`
			}{Message: "Terjadi kesalahan saat autentikasi.", Error: err.Error()})
			c.Abort()
			return
		}

		if strings.Compare(SISWA, claims.Role) != 0 {
			c.JSON(http.StatusForbidden, struct {
				Message string `json:"message"`
				Error string `json:"error"`
			}{Message: "Anda tidak diperbolehkan mengakses data ini.", Error: utility.ErrForbiddedn.Error()})
			c.Abort()
			return
		}

		var idUser int = claims.IdUser
		c.Set("idUser", idUser)
		c.Next()
	}
}

func ValidateMitraRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := auth.ExtractJwtFromHeader(c.Request)

		if len(tokenString) == 0 {
			c.JSON(http.StatusUnauthorized, struct {
				Message string `json:"message"`
				Error string `json:"error"`
			}{Message: "Token tidak ada.", Error: utility.ErrUnauthorized.Error()})
			c.Abort()
			return
		}

		claims, err := auth.GetClaimsFromJwt(tokenString)
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.JSON(http.StatusUnauthorized, struct {
					Message string `json:"message"`
					Error string `json:"error"`
				}{Message: "Terjadi kesalahan saat autentikasi.", Error: err.Error()})
				c.Abort()
				return
			}

			c.JSON(http.StatusBadRequest, struct {
				Message string `json:"message"`
				Error string `json:"error"`
			}{Message: "Terjadi kesalahan saat autentikasi.", Error: err.Error()})
			c.Abort()
			return
		}

		if strings.Compare(MITRA, claims.Role) != 0 {
			c.JSON(http.StatusForbidden, struct {
				Message string `json:"message"`
				Error string `json:"error"`
			}{Message: "Anda tidak diperbolehkan mengakses data ini.", Error: utility.ErrForbiddedn.Error()})
			c.Abort()
			return
		}


		var idUser int = claims.IdUser
		c.Set("idUser", idUser)
		c.Next()
	}
}

func ValidateRoles(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := auth.ExtractJwtFromHeader(c.Request)

		if len(tokenString) == 0 {
			c.JSON(http.StatusUnauthorized, struct {
				Message string `json:"message"`
				Error string `json:"error"`
			}{Message: "Token tidak ada.", Error: utility.ErrUnauthorized.Error()})
			c.Abort()
			return
		}

		if len(roles) == 0 {
			c.JSON(http.StatusUnauthorized, struct {
				Message string `json:"message"`
				Error string `json:"error"`
			}{Message: "Pastikan roles tidak kosong.", Error: utility.ErrBadRequest.Error()})
			return
		}

		claims, err := auth.GetClaimsFromJwt(tokenString)
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.JSON(http.StatusUnauthorized, struct {
					Message string `json:"message"`
					Error string `json:"error"`
				}{Message: "Terjadi kesalahan saat autentikasi.", Error: err.Error()})
				c.Abort()
				return
			}

			c.JSON(http.StatusBadRequest, struct {
				Message string `json:"message"`
				Error string `json:"error"`
			}{Message: "Terjadi kesalahan saat autentikasi.", Error: err.Error()})
			c.Abort()
			return
		}

		isRoleExists := false
		role := ""
		for _, item := range roles {
			if strings.Compare(item, claims.Role) == 0 {
				isRoleExists = true
				role = item
			}
		}

		log.Println("masuk?")
		if !isRoleExists {
			c.JSON(http.StatusForbidden, struct {
				Message string `json:"message"`
				Error string `json:"error"`
			}{Message: "Anda tidak diperbolehkan mengakses data ini.", Error: utility.ErrForbiddedn.Error()})
			c.Abort()
			return
		}


		var idUser int = claims.IdUser
		c.Set("idUser", idUser)
		c.Set("role", role)
		log.Println("rolenya:", role)
		c.Next()
	}
}
