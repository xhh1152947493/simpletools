package main

import (
	"github.com/gin-gonic/gin"
	ctx "simpletools/internal/api/context"
	"simpletools/internal/api/handlers"
	"simpletools/internal/api/middlewares"
	"simpletools/internal/bootstrap"
	"simpletools/internal/data"
)

const (
	useConfig = "server_simpletools_debug.json"
	pidFile   = "simpletools.pid"
)

func wrapHandler(cb ctx.CustomHandlerFunc) gin.HandlerFunc {
	return ctx.WrapHandler(cb)
}

func registerHandlers(r *gin.Engine) {
	publicRoutes := r.Group("/api/", middlewares.Validate(false))
	{
		publicRoutes.POST("/aitranslate", wrapHandler(handlers.OnAITranslateHandler))
		publicRoutes.POST("/ainamed", wrapHandler(handlers.OnAINamedHandler))
	}
}

func main() {
	if err := data.InitGlobal(useConfig); err != nil {
		panic(err)
	}

	r := gin.New()
	r.Use(middlewares.Recover())
	r.Use(middlewares.Exit())
	r.Use(middlewares.Cors())
	r.Use(middlewares.Log())
	r.Use(middlewares.CustomContext())

	registerHandlers(r)

	bootstrap.Bootstrap(pidFile, r)
}
