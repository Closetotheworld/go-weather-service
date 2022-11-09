package weather

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
)

var (
	forecastHourOffset   = []string{"6", "12", "18", "24", "30", "36", "42", "48"}
	historicalHourOffset = []string{"-6", "-12", "-18", "-24"}
)

type WeatherApiCommon struct {
	Timestamp  uint    `json:"timestamp"`
	Code       int     `json:"code"`
	Temp       float32 `json:"temp"`
	Rain1h     int     `json:"rain1h"`
	HourOffset string
}

type WeatherApiForecast struct {
	Timestamp  uint    `json:"timestamp"`
	Code       int     `json:"code"`
	MinTemp    float32 `json:"min_temp"`
	MaxTemp    float32 `json:"max_temp"`
	Rain1h     int     `json:"rain1h"`
	HourOffset string
}

type WeatherApiManager interface {
	GetCurrentInfo(lat string, lon string) (*WeatherApiCommon, error)
	GetForecastInfo(lat string, lon string, hourOffset string) (*WeatherApiForecast, error)
	GetHistoricalInfo(lat string, lon string, hourOffset string) (*WeatherApiCommon, error)
}

type WeatherApiManagerImpl struct {
	ApiKey string
	Url    string
}

func (w *WeatherApiManagerImpl) GetCurrentInfo(lat string, lon string) (*WeatherApiCommon, error) {
	result := &WeatherApiCommon{}

	query := url.Values{}
	query.Add("api_key", w.ApiKey)
	query.Add("lat", lat)
	query.Add("lon", lon)

	base, err := url.Parse("https://thirdparty-weather-api-v2.droom.workers.dev/current")
	if err != nil {
		return nil, err
	}
	base.RawQuery = query.Encode()

	resp, err := http.Get(base.String())
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (w *WeatherApiManagerImpl) GetForecastInfo(lat string, lon string, hourOffset string) (*WeatherApiForecast, error) {
	result := &WeatherApiForecast{}

	query := url.Values{}
	query.Add("api_key", w.ApiKey)
	query.Add("lat", lat)
	query.Add("lon", lon)
	query.Add("hour_offset", hourOffset)

	base, err := url.Parse("https://thirdparty-weather-api-v2.droom.workers.dev/forecast/hourly")
	if err != nil {
		return nil, err
	}
	base.RawQuery = query.Encode()

	resp, err := http.Get(base.String())
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, result)
	if err != nil {
		return nil, err
	}
	result.HourOffset = hourOffset
	return result, nil
}

func (w *WeatherApiManagerImpl) GetHistoricalInfo(lat string, lon string, hourOffset string) (*WeatherApiCommon, error) {
	result := &WeatherApiCommon{}

	query := url.Values{}
	query.Add("api_key", w.ApiKey)
	query.Add("lat", lat)
	query.Add("lon", lon)
	query.Add("hour_offset", hourOffset)

	base, err := url.Parse("https://thirdparty-weather-api-v2.droom.workers.dev/historical/hourly")
	if err != nil {
		return nil, err
	}
	base.RawQuery = query.Encode()

	resp, err := http.Get(base.String())
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, result)
	if err != nil {
		return nil, err
	}

	result.HourOffset = hourOffset
	return result, nil
}

func (w *WeatherApiManagerImpl) AsyncRequest(lat string, lon string) (*WeatherApiCommon, []*WeatherApiForecast, []*WeatherApiCommon) {
	var currentWeather *WeatherApiCommon
	var forecastWeather []*WeatherApiForecast
	var historicalWeather []*WeatherApiCommon

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		currentWeather, _ = w.GetCurrentInfo(lat, lon)
	}()

	for fh := range forecastHourOffset {
		wg.Add(1)
		go func(hourOffset string) {
			defer wg.Done()
			fw, _ := w.GetForecastInfo(lat, lon, hourOffset)
			forecastWeather = append(forecastWeather, fw)
		}(forecastHourOffset[fh])
	}

	for hh := range historicalHourOffset {
		wg.Add(1)
		go func(hourOffset string) {
			defer wg.Done()
			hw, _ := w.GetHistoricalInfo(lat, lon, hourOffset)
			historicalWeather = append(historicalWeather, hw)
		}(historicalHourOffset[hh])
	}
	wg.Wait()

	return currentWeather, forecastWeather, historicalWeather

}
