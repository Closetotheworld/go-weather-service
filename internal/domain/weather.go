package domain

import (
	"context"
	// external packages

	// project packages
	"github.com/closetotheworld/go-weather-service/internal/model"
)

type WeatherService interface {
	GetWeatherSummary(ctx context.Context) (*model.Weather, error)
}
