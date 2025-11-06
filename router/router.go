package router

import (
	// "weather-api-wrapper-service/handlers"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func Init() {
	router = gin.Default()

	InitializeRoutes(router)

	router.Run("0.0.0.0:8080")
}
