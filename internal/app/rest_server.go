package app

import (
	"log"
	"strings"

	// external packages
	"github.com/gin-gonic/gin"

	// project packages
	"github.com/closetotheworld/go-weather-service/internal/controller"
	"github.com/closetotheworld/go-weather-service/internal/service/weather"
	"github.com/closetotheworld/go-weather-service/pkg/weather_api"
)

func StartServer(apiKey string, port string) {
	r := gin.Default()

	weatherService := weather.NewWeatherService(&weather_api.WeatherApiManagerImpl{ApiKey: apiKey})
	weatherController := controller.NewWeatherHeandler(weatherService)

	r.GET("/summary", weatherController.GetWeatherSummary)

	if err := r.Run(strings.Join([]string{":", port}, "")); err != nil {
		log.Panic(err)
	}
}
