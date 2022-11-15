package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	// external packages
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	// project packages
	"github.com/closetotheworld/go-weather-service/internal/model"
	"github.com/closetotheworld/go-weather-service/internal/service/weather"
)

func TestNewWeatherHeandler(t *testing.T) {
	mockWeatherManager := weather.NewMockWeatherService(t)
	s := NewWeatherHeandler(mockWeatherManager)
	assert.NotNil(t, s)
}

func TestWeatherHandler_GetWeatherSummary(t *testing.T) {
	//c, _ := gin.CreateTestContext()
	// handler test를 위한 func variable
	testHandlerFunc := func(engine *gin.Engine, lat *string, lon *string) *httptest.ResponseRecorder {
		httpRequest, _ := http.NewRequest(http.MethodGet, "/summary", nil)

		q := httpRequest.URL.Query()
		if lat != nil {
			q.Add("lat", *lat)
		}
		if lon != nil {
			q.Add("lon", *lon)
		}

		httpRequest.URL.RawQuery = q.Encode()

		w := httptest.NewRecorder()
		engine.ServeHTTP(w, httpRequest)
		return w
	}

	// valid param test
	t.Run("test valid param", func(t *testing.T) {
		mockWeatherService := weather.NewMockWeatherService(t)
		w := NewWeatherHeandler(mockWeatherService)
		mockResult := model.WeatherResult{
			Summary: model.Weather{
				Greeting:    "mockGreeting",
				Temperature: "mockTemperture",
				HeadsUp:     "mockHeadsUp",
			},
		}
		mockWeatherService.On("GetWeatherSummary", mock.Anything, mock.Anything, mock.Anything).Return(&mockResult, nil)

		r := gin.Default()
		r.GET("/summary", w.GetWeatherSummary)

		lat := "-45"
		lon := "-45"

		res := testHandlerFunc(r, &lat, &lon)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.NotNil(t, res.Body)
	})

	// invalid param test
	t.Run("test invalid param", func(t *testing.T) {
		t.Run("invalid param(out of range)", func(t *testing.T) {
			mockWeatherService := weather.NewMockWeatherService(t)
			w := NewWeatherHeandler(mockWeatherService)

			r := gin.Default()
			r.GET("/summary", w.GetWeatherSummary)

			lat := "-100"
			lon := "-120"

			res := testHandlerFunc(r, &lat, &lon)
			assert.Equal(t, http.StatusBadRequest, res.Code)
		})
		t.Run("invalid param(not add query param)", func(t *testing.T) {
			mockWeatherService := weather.NewMockWeatherService(t)
			w := NewWeatherHeandler(mockWeatherService)

			r := gin.Default()
			r.GET("/summary", w.GetWeatherSummary)

			res := testHandlerFunc(r, nil, nil)
			assert.Equal(t, http.StatusBadRequest, res.Code)
		})

	})

	t.Run("internal server error (service or pkg api occured error)", func(t *testing.T) {
		mockWeatherService := weather.NewMockWeatherService(t)
		w := NewWeatherHeandler(mockWeatherService)
		mockWeatherService.On("GetWeatherSummary", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("some error"))

		r := gin.Default()
		r.GET("/summary", w.GetWeatherSummary)

		lat := "-45"
		lon := "-45"

		res := testHandlerFunc(r, &lat, &lon)
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
}
