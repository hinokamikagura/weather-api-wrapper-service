package router

import (
	// "weather-api-wrapper-service/handlers"

	"github.com/gin-gonic/gin"
	"github.com/hinokamikagura/weather-api-wrapper-service/handler"
)

func InitializeRoutes(router *gin.Engine) {
	basePath := "/api/v1"
	handler.Init()
	v1 := router.Group(basePath)
	{
		v1.GET("/weather", handler.GetCityWeather)
		v1.GET("/forecast", func(ctx *gin.Context) {
			ctx.String(200, "Forecast endpoint")
		})
	}
}
