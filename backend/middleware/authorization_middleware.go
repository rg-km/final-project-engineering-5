package middleware

import (
	"FinalProject/auth"
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
			}{Message: "Anda tidak diperbolehkan mengakses data ini.", Error: err.Error()})
			c.Abort()
			return
		}

		c.Next()
	}
}

func ValidateMitraRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := auth.ExtractJwtFromHeader(c.Request)

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
			}{Message: "Anda tidak diperbolehkan mengakses data ini.", Error: err.Error()})
			c.Abort()
			return
		}

		c.Next()
	}
}
