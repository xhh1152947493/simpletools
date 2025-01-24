package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
	"simpletools/internal/data"
	"simpletools/internal/defs"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				tips := fmt.Sprintf("http request panic!\n%v|%s", err, string(debug.Stack()))
				data.Log().Error().HttpRequest(c.Request).UID(1).Msg(tips)
				c.AbortWithStatusJSON(http.StatusOK, gin.H{"code": defs.ErrCodeSystemPanic, "msg": "server internal panic", "data": nil})
			}
		}()
		c.Next()
	}
}
