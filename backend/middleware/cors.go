package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
			c.JSON(http.StatusNoContent, "")
			return
		}

		c.Request.Header.Set("Access-Control-Allow-Origin", "*")
		c.Request.Header.Set("Access-Control-Allow-Credential", "true")
		c.Request.Header.Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Origin, Authorization")
		c.Request.Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, HEAD")
		c.Next()
	}
}