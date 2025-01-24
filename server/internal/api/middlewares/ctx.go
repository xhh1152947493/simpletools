package middlewares

import (
	"github.com/gin-gonic/gin"
	ctx "simpletools/internal/api/context"
)

func CustomContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		customContext := &ctx.CustomContext{Ctx: c}
		c.Set("customContext", customContext)
		c.Next()
	}
}
