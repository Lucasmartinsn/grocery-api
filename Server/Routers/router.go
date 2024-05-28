package Routers

import (
	"github.com/Lucasmartinsn/grocery-api/Server/Middlewares"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	router.Use(Middlewares.CORSMiddleware())
	router.Use(gin.Recovery())
	main := router.Group("api")
	{
		Login := main.Group("auth")
		{
			Login.POST("/",)
		}
	}
	return router
}
