package router

import (
	// "weather-api-wrapper-service/handlers"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/hinokamikagura/weather-api-wrapper-service/handler"
	"github.com/khaaleoo/gin-rate-limiter/core"
)

func InitializeRoutes(router *gin.Engine) {

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	rateLimiterOption := core.RateLimiterOption{
		Limit: 1,
		Burst: 50,
		Len:   1 * time.Minute,
	}

	rateLimiterMiddleware := core.RequireRateLimiter(core.RateLimiter{
		RateLimiterType: core.IPRateLimiter,
		Key:             "iplimiter_maximum_requests_for_ip_test",
		Option:          rateLimiterOption,
	})
	basePath := "/api/v1"
	handler.Init()
	v1 := router.Group(basePath)
	{
		v1.GET("/weather", rateLimiterMiddleware, handler.GetCityWeather)
		v1.GET("/forecast", func(ctx *gin.Context) {
			ctx.String(200, "Forecast endpoint")
		})
	}
}
