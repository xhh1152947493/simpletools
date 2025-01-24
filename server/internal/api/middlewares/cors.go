package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simpletools/internal/data"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		if c.Request.Method == http.MethodOptions {
			data.Log().Info().HttpRequest(c.Request).Msg("http options request end")
			c.AbortWithStatus(http.StatusNoContent)
		}
	}
}
