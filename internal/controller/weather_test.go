package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	// external packages
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	// project packages
	"github.com/closetotheworld/go-weather-service/internal/service"
)

func TestNewWeatherHeandler(t *testing.T) {
	mockWeatherManager := service.NewMockWeatherService()
	s := NewWeatherHeandler(mockWeatherManager)
	assert.Equal(t, s, WeatherHandler{})
}

func TestWeatherHandler_GetWeatherSummary(t *testing.T) {
	w := NewWeatherHeandler(service.NewMockWeatherService())
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	r := gin.Default()
	r.GET("/summary", w.GetWeatherSummary)

	// handler test를 위한 func variable
	testHandlerFunc := func(engine *gin.Engine, lat string, lon string) *httptest.ResponseRecorder {
		httpRequest, _ := http.NewRequest(http.MethodGet, "/summary", nil)

		q := httpRequest.URL.Query()
		q.Add("lat", lat)
		q.Add("lon", lon)
		httpRequest.URL.RawQuery = q.Encode()

		w := httptest.NewRecorder()
		engine.ServeHTTP(w, httpRequest)
		return w
	}

	// valid param test
	t.Run("test valid param", func(t *testing.T) {
		res := testHandlerFunc(r, "-45", "-45")
		w.GetWeatherSummary(c)
		assert.Equal(t, http.StatusOK, res.Code)
	})

	// invalid param test
	t.Run("test invalid param", func(t *testing.T) {
		res := testHandlerFunc(r, "-180", "-45")
		w.GetWeatherSummary(c)
		assert.Equal(t, http.StatusBadRequest, res.Code)
	})
}
