package domain

import (
	"context"
	// external packages

	// project packages
	"github.com/closetotheworld/go-weather-service/internal/model"
)

type WeatherService interface {
	GetWeatherSummary(ctx context.Context, lat float32, lon float32) (*model.WeatherResult, error)
}
