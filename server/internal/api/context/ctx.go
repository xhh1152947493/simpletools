package ctx

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simpletools/internal/data"
	"simpletools/internal/defs"
	"simpletools/internal/utils"
)

type CustomHandlerFunc func(*CustomContext) *defs.CustomError

type CustomContext struct {
	Ctx *gin.Context

	decodeParams utils.AnyJson
}

func WrapHandler(f CustomHandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		cc, _ := c.Get("customContext")
		ctx := cc.(*CustomContext)
		if cErr := f(ctx); cErr != nil {
			ctx.Abort(cErr, nil)
			data.GLog.Error().CErr(cErr).Msg("server report error")
		}
	}
}

func (cc *CustomContext) Abort(cErr *defs.CustomError, data any) {
	cc.Ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"code": cErr.GetCode(), "msg": cErr.GetErr(), "data": data})
}

func (cc *CustomContext) AnswerOK(data any) {
	cc.Ctx.JSON(http.StatusOK, gin.H{"code": defs.ErrCodeOK, "msg": "ok", "data": data})
}

func (cc *CustomContext) tryInitParams() {
	if cc.decodeParams != nil {
		return
	}
	decodeParams := make(utils.AnyJson)
	_ = cc.Ctx.ShouldBindJSON(&decodeParams)
	cc.decodeParams = decodeParams
}

func (cc *CustomContext) GetString(key string) string {
	cc.tryInitParams()
	return cc.decodeParams.GetString(key)
}

func (cc *CustomContext) GetInt64(key string) int64 {
	cc.tryInitParams()
	return cc.decodeParams.GetInt64(key)
}

func (cc *CustomContext) Username() string {
	return cc.Ctx.GetString("username")
}

func (cc *CustomContext) Platform() string {
	return cc.Ctx.GetString("platform")
}
