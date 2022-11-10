package service

import (
	"context"
	"net/http"

	// external packages

	// project packages
	"github.com/closetotheworld/go-weather-service/internal/domain"
	e "github.com/closetotheworld/go-weather-service/internal/err"
	"github.com/closetotheworld/go-weather-service/internal/model"
	"github.com/closetotheworld/go-weather-service/pkg/weather"
)

type WeatherService struct {
	WeatherApiManager weather.WeatherApiManager
}

func NewWeatherService(wm weather.WeatherApiManager) domain.WeatherService {
	return &WeatherService{WeatherApiManager: wm}
}

func (w *WeatherService) GetWeatherSummary(ctx context.Context, lat string, lon string) (*model.Weather, error) {
	current, _, _, err := w.WeatherApiManager.AsyncRequest(lat, lon)
	if err != nil {
		return nil, e.ErrorByStatus(http.StatusInternalServerError)
	}
	mock := model.Weather{
		Gretting:   ParseGreeting(*current),
		Temperture: ParseTemperture(),
		HeadsUp:    ParseHeadsUp(),
	}
	return &mock, nil
}
