package weather

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	API_KEY = ""
	w       = WeatherApiManagerImpl{}
)

func TestWeatherApiManagerImpl_GetCurrentInfo(t *testing.T) {
	t.Run("test valid request", func(t *testing.T) {
		w.ApiKey = API_KEY
		result, err := w.GetCurrentInfo("10", "-120")
		assert.NotNil(t, result)
		assert.Nil(t, err, nil)
		t.Log(result)
	})
	t.Run("test bad request", func(t *testing.T) {
		w.ApiKey = API_KEY
		result, err := w.GetCurrentInfo("-100", "-200")
		assert.Nil(t, result)
		assert.NotNil(t, err)
	})
	t.Run("test unauthrized", func(t *testing.T) {
		w.ApiKey = ""
		_, err := w.GetCurrentInfo("-100", "-200")
		assert.NotNil(t, err)
	})
}

func TestWeatherApiManagerImpl_GetForecastInfo(t *testing.T) {
	t.Run("test valid request", func(t *testing.T) {
		w.ApiKey = API_KEY
		result, err := w.GetForecastInfo("10", "-120", "12")
		t.Log(result)
		assert.NotNil(t, result)
		assert.Nil(t, err, nil)
	})
	t.Run("test bad request", func(t *testing.T) {
		w.ApiKey = API_KEY
		result, err := w.GetForecastInfo("-100", "-200", "5")
		assert.Nil(t, result)
		assert.NotNil(t, err)
	})
	t.Run("test unauthorized", func(t *testing.T) {
		w.ApiKey = ""
		_, err := w.GetForecastInfo("-100", "-200", "16")
		assert.NotEqual(t, err, nil)
	})
}

func TestWeatherApiManagerImpl_GetHistoricalInfo(t *testing.T) {
	t.Run("test valid request", func(t *testing.T) {
		w.ApiKey = API_KEY
		result, err := w.GetHistoricalInfo("10", "-120", "-6")
		t.Log(result)
		assert.NotNil(t, result)
		assert.Nil(t, err, nil)
	})
	t.Run("test bad request", func(t *testing.T) {
		w.ApiKey = API_KEY
		result, err := w.GetHistoricalInfo("-100", "-200", "-10")
		assert.Nil(t, result)
		assert.NotNil(t, err)
	})
	t.Run("test unauthorized", func(t *testing.T) {
		_, err := w.GetHistoricalInfo("-100", "-200", "2")
		assert.NotEqual(t, err, nil)
	})
}

func TestWeatherApiManagerImpl_AsyncRequest(t *testing.T) {
	t.Run("run", func(t *testing.T) {
		w.ApiKey = API_KEY
		current, forecast, historical, err := w.AsyncRequest("10", "-120")
		t.Log(current)
		for i := range forecast {
			t.Log(forecast[i])
		}
		for i := range historical {
			t.Log(historical[i])
		}
		assert.Nil(t, err)
	})
}
