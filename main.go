package main

import (
	"fmt"

	"github.com/hinokamikagura/weather-api-wrapper-service/router"
)

func main() {
	router.Init()
	// Additional routes and handlers would be defined here
	fmt.Println("This is weather api wrapper service.")
}
