package app

import (
	"log"
	"strings"

	// external packages
	"github.com/gin-gonic/gin"

	// project packages
	"github.com/closetotheworld/go-weather-service/internal/controller"
	"github.com/closetotheworld/go-weather-service/internal/service"
	"github.com/closetotheworld/go-weather-service/pkg/weather"
)

func StartServer(apiKey string, port string) {
	r := gin.Default()

	weatherService := service.NewWeatherService(&weather.WeatherApiManagerImpl{ApiKey: apiKey})
	weatherController := controller.NewWeatherHeandler(weatherService)

	r.GET("/summary", weatherController.GetWeatherSummary)

	if err := r.Run(strings.Join([]string{":", port}, "")); err != nil {
		log.Panic(err)
	}
}
