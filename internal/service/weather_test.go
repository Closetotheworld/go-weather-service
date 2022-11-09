package service

import (
	"context"
	"testing"

	// external packages
	"github.com/stretchr/testify/assert"

	// project packages
	"github.com/closetotheworld/go-weather-service/pkg/weather"
)

func TestNewWeatherService(t *testing.T) {
	mockWeatherManager := weather.MockWeatherApiImpl{}
	s := NewWeatherService(&mockWeatherManager)
	assert.Equal(t, s, WeatherService{})
}

func TestWeatherService_GetWeatherSummary(t *testing.T) {
	ctx := context.Background()
	mockWeatherManager := weather.MockWeatherApiImpl{}
	ws := NewWeatherService(&mockWeatherManager)

	w, err := ws.GetWeatherSummary(ctx)
	assert.Nil(t, err)
	t.Log(w)
	// gretting test case

	// temperture test case

	// headsUp test case
}
