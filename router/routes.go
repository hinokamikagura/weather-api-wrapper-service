package router

import (
	// "weather-api-wrapper-service/handlers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	basePath := "/api/v1"
	v1 := router.Group(basePath)
	{
		v1.GET("/weather", func(ctx *gin.Context) {
			ctx.String(200, "Weather endpoint")
		})
		v1.GET("/forecast", func(ctx *gin.Context) {
			ctx.String(200, "Forecast endpoint")
		})
	}
}
