package app

import (
	"log"
	"strings"

	// external packages
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// project packages
	"github.com/closetotheworld/go-weather-service/docs"
	"github.com/closetotheworld/go-weather-service/internal/app/middleware"
	"github.com/closetotheworld/go-weather-service/internal/controller"
	"github.com/closetotheworld/go-weather-service/internal/service/weather"
	"github.com/closetotheworld/go-weather-service/pkg/weather_api"
)

func StartServer(apiKey string, port string) {
	r := gin.Default()
	r.Use(middleware.TimeOutHandler())

	docs.SwaggerInfo.Title = "Weather Summary api"
	docs.SwaggerInfo.Description = "API documentation for Weather Summary api"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	weatherService := weather.NewWeatherService(&weather_api.WeatherApiManagerImpl{ApiKey: apiKey})
	weatherController := controller.NewWeatherHeandler(weatherService)
	r.GET("/summary", weatherController.GetWeatherSummary)

	if err := r.Run(strings.Join([]string{":", port}, "")); err != nil {
		log.Panic(err)
	}
}
