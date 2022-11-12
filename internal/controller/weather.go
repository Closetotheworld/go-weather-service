package controller

import (
	"log"
	"net/http"

	// external packages
	"github.com/gin-gonic/gin"

	// project packages
	"github.com/closetotheworld/go-weather-service/internal/domain"
)

type InputWeatherSummary struct {
	Lat float32 `form:"lat" binding:"required,min=-90.0,max=90.0"`
	Lon float32 `form:"lon" binding:"required,min=-180.0,max=180.0"`
}

type WeatherHandler struct {
	weatherService domain.WeatherService
}

func NewWeatherHeandler(weatherService domain.WeatherService) *WeatherHandler {
	return &WeatherHandler{weatherService: weatherService}
}

// GetWeatherSummary
// @Summary GetWeatherSummary
// @Description weather에 대한 summary를 반환환다.
// @Accept  json
// @Produce  json
// @Param default query InputWeatherSummary true "summary inputs"
// @Router /summary [get]
// @Success 200 {object} model.WeatherResult "Success"
// @Failure 400 "Invalid parameter requested"
// @Failure 408 "request timeout"
// @Failure 500 "Internal Server error"
func (w *WeatherHandler) GetWeatherSummary(c *gin.Context) {
	inputQuery := InputWeatherSummary{}

	if err := c.ShouldBindQuery(&inputQuery); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid parameter requested")
		log.Print(err.Error())
		return
	}

	result, err := w.weatherService.GetWeatherSummary(c, inputQuery.Lat, inputQuery.Lon)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server error")
		return
	}
	c.JSON(http.StatusOK, result)
	return
}
