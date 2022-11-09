package service

import (
	"context"

	// external packages

	// project packages
	"github.com/closetotheworld/go-weather-service/internal/domain"
	"github.com/closetotheworld/go-weather-service/internal/model"
	"github.com/closetotheworld/go-weather-service/pkg/weather"
)

type WeatherService struct {
	weatherApiManager *weather.WeatherApiManager
}

func NewWeatherService(wm weather.WeatherApiManager) domain.WeatherService {
	return &WeatherService{weatherApiManager: &wm}
}

func (w WeatherService) GetWeatherSummary(ctx context.Context) (*model.Weather, error) {

	mock := model.Weather{
		Gretting:   "gretting",
		Temperture: "temperture",
		HeadsUp:    "headsUp",
	}
	return &mock, nil
}
