package service

import (
	"context"

	// external packages

	// project packages
	"github.com/closetotheworld/go-weather-service/internal/domain"
	"github.com/closetotheworld/go-weather-service/internal/model"
)

type WeatherService struct {
}

func NewWeatherService() domain.WeatherService {
	return &WeatherService{}
}

func (w WeatherService) GetWeatherSummary(ctx context.Context) (*model.Weather, error) {
	mock := model.Weather{
		Gretting:   "gretting",
		Temperture: "temperture",
		HeadsUp:    "headsUp",
	}
	return &mock, nil
}
