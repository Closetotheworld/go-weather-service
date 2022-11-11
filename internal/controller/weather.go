package controller

import (
	// external packages
	"github.com/gin-gonic/gin"

	// project packages
	"github.com/closetotheworld/go-weather-service/internal/domain"
)

type WeatherHandler struct {
	weatherService domain.WeatherService
}

func NewWeatherHeandler(weatherService domain.WeatherService) *WeatherHandler {
	return &WeatherHandler{weatherService: weatherService}
}

func (w *WeatherHandler) GetWeatherSummary(c *gin.Context) {
	lat := c.Query("lat")
	lon := c.Query("lon")
	result, err := w.weatherService.GetWeatherSummary(c, lat, lon)

	if err != nil {
		c.JSON(400, "something")
	}
	c.JSON(200, result)
}
