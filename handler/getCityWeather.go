package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hinokamikagura/weather-api-wrapper-service/connection"
	"github.com/redis/go-redis/v9"
)

func GetCityWeather(c *gin.Context) {
	city := c.Query("city")
	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "City parameter is required"})
		return
	}

	ctx := context.Background()
	rdb := connection.RedisConnect()

	cacheKey := strings.ToLower("weather_" + city)
	cacheData, err := rdb.Get(ctx, cacheKey).Result()

	if err == redis.Nil {

		fmt.Println("GetCityWeather called", weatherAPIBaseURL+"?key="+apiKey+"&q="+city)

		resp, err := http.Get(weatherAPIBaseURL + "?key=" + apiKey + "&q=" + city)

		if err != nil || resp.StatusCode != http.StatusOK {
			sendError(c, http.StatusInternalServerError, err.Error())
			return
		}

		defer resp.Body.Close()

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			sendError(c, http.StatusInternalServerError, err.Error())
			return
		}

		var apiResponse map[string]interface{}

		err = json.Unmarshal(respBody, &apiResponse)
		if err != nil {
			sendError(c, http.StatusInternalServerError, err.Error())
			return
		}

		// Placeholder for actual weather fetching logic
		weatherData := map[string]interface{}{
			"city": city,
			"data": apiResponse,
		}
		err = rdb.Set(ctx, cacheKey, respBody, 0).Err()
		if err != nil {
			log.Println("Failed to set cache:", err)
		}
		sendSuccess(c, weatherData)
	} else if err != nil {
		log.Println("Redis error : %v", err)
		sendError(c, http.StatusInternalServerError, err.Error())
	}

	var cacheResult map[string]interface{}
	json.Unmarshal([]byte(cacheData), &cacheResult)
	cashResponse := map[string]interface{}{
		"city": city,
		"data": cacheResult,
	}
	fmt.Println("Cache data loaded:")
	sendSuccess(c, cashResponse)

}
