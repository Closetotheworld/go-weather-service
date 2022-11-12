package weather

import (
	"context"
	"net/http"

	// external packages

	// project packages
	"github.com/closetotheworld/go-weather-service/internal/domain"
	e "github.com/closetotheworld/go-weather-service/internal/err"
	"github.com/closetotheworld/go-weather-service/internal/model"
	"github.com/closetotheworld/go-weather-service/pkg/weather_api"
)

type WeatherService struct {
	WeatherApiManager weather_api.WeatherApiManager
}

func NewWeatherService(wm weather_api.WeatherApiManager) domain.WeatherService {
	return &WeatherService{WeatherApiManager: wm}
}

func (w *WeatherService) GetWeatherSummary(ctx context.Context, lat float32, lon float32) (*model.WeatherResult, error) {
	current, forecast, historical, err := w.WeatherApiManager.AsyncRequest(lat, lon)
	if err != nil {
		return nil, e.ErrorByStatus(http.StatusInternalServerError)
	}

	parseResult := model.Weather{
		Greeting:   ParseGreeting(*current),
		Temperture: ParseTemperture(current.Temp, historical),
		HeadsUp:    ParseHeadsUp(forecast),
	}
	return &model.WeatherResult{Summary: parseResult}, nil
}
