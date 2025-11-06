package handler

import (
	"github.com/gin-gonic/gin"
)

var (
	weatherAPIBaseURL string
)

func Init(c *gin.Context) {
	weatherAPIBaseURL = "http://api.weatherapi.com/v1/current.json"
}
