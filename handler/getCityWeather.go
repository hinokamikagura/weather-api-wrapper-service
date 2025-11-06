package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCityWeather(c *gin.Context) {
	fmt.Println("GetCityWeather called")
	city := c.Query("city")
	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "City parameter is required"})
		return
	}

	// client := &http.Client{}
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

	// data, err := json.MarshalIndent(apiResponse, "", "  ")

	// fmt.Println("Fetching weather data for city:", string(data))

	// Placeholder for actual weather fetching logic
	weatherData := map[string]interface{}{
		"city": city,
		"data": apiResponse,
	}

	sendSuccess(c, weatherData)
}
