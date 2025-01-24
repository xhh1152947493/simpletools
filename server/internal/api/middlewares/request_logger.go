package middlewares

import (
	"github.com/gin-gonic/gin"
	"simpletools/internal/data"
	"time"
)

func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := data.Now()
		c.Next()
		data.Log().Info().HttpRequest(c.Request).Int64("cost", time.Since(start).Milliseconds()).Msg("http request over")
	}
}
