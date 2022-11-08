package app

import (
	"log"

	// external packages
	"github.com/gin-gonic/gin"

	// project packages
	"github.com/closetotheworld/go-weather-service/internal/controller"
	"github.com/closetotheworld/go-weather-service/internal/service"
)

func StartServer() {
	r := gin.New()

	weatherService := service.NewWeatherService()
	weatherController := controller.NewWeatherHeandler(weatherService)

	r.GET("/summary", weatherController.GetWeatherSummary)
	if err := r.Run(); err != nil {
		log.Panic(err)
	}
}
