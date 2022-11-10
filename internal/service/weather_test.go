package service

import (
	"context"
	"testing"

	// external packages
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	// project packages
	"github.com/closetotheworld/go-weather-service/pkg/weather"
)

func TestNewWeatherService(t *testing.T) {
	mockWeatherManager := weather.NewMockWeatherApiManager(t)
	s := NewWeatherService(mockWeatherManager)
	assert.NotNil(t, s)
}

func TestWeatherService_GetWeatherSummary(t *testing.T) {
	ctx := context.Background()

	// gretting test case
	t.Run("greeting test", func(t *testing.T) {
		for i := 0; i < 8; i++ {
			mockWeatherManager := weather.NewMockWeatherApiManager(t)
			ws := NewWeatherService(mockWeatherManager)

			wc, result := GetGreetingTestCase(i)
			mockWeatherManager.On("AsyncRequest", mock.Anything, mock.Anything).Return(wc.Current, wc.Forecast, wc.Historical, nil)

			w, err := ws.GetWeatherSummary(ctx, "10", "120")
			assert.Nil(t, err)
			assert.Equal(t, result, w.Gretting)
		}
	})

	// temperture test case
	t.Run("temperture test", func(t *testing.T) {
		for i := 0; i < 6; i++ {
			mockWeatherManager := weather.NewMockWeatherApiManager(t)
			ws := NewWeatherService(mockWeatherManager)

			wc, result := GetTempertureTestCase(i)
			mockWeatherManager.On("AsyncRequest", mock.Anything, mock.Anything).Return(wc.Current, wc.Forecast, wc.Historical, nil)

			w, err := ws.GetWeatherSummary(ctx, "10", "120")
			assert.Nil(t, err)
			assert.Equal(t, result, w.Temperture)
		}
	})

	// headsUp test case
	t.Run("headsUp test", func(t *testing.T) {
		for i := 0; i < 5; i++ {
			mockWeatherManager := weather.NewMockWeatherApiManager(t)
			ws := NewWeatherService(mockWeatherManager)

			wc, result := GetHeadsUpTestCase(i)
			mockWeatherManager.On("AsyncRequest", mock.Anything, mock.Anything).Return(wc.Current, wc.Forecast, wc.Historical, nil)

			w, err := ws.GetWeatherSummary(ctx, "10", "120")
			assert.Nil(t, err)
			assert.Equal(t, result, w.HeadsUp)
		}
	})
}
