package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCityWeather(c *gin.Context) {
	city := c.Query("city")
	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "City parameter is required"})
		return
	}

	// Placeholder for actual weather fetching logic
	weatherData := map[string]interface{}{
		"city":        city,
		"temperature": "25°C",
		"condition":   "Sunny",
	}

	c.JSON(http.StatusOK, weatherData)
}
