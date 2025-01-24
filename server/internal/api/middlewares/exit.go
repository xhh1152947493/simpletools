package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"simpletools/internal/data"
)

func Exit() gin.HandlerFunc {
	return func(c *gin.Context) {
		if data.GIsExiting.Load() { // 服务器正在退出，拒绝新的请求，记录相关url参数及和body的内容
			data.Log().Warn().HttpRequest(c.Request).Msg(fmt.Sprintf("http recv request when server exit"))
			c.AbortWithStatus(http.StatusServiceUnavailable)
		}
	}
}
