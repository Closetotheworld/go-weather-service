package weather_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	// external pacakges
	"golang.org/x/sync/errgroup"
	// project packages
)

var (
	forecastHourOffset   = []int{6, 12, 18, 24, 30, 36, 42, 48}
	historicalHourOffset = []int{-6, -12, -18, -24}
)

type WeatherApiCommon struct {
	Timestamp  uint    `json:"timestamp"`
	Code       int     `json:"code"`
	Temp       float32 `json:"temp"`
	Rain1h     int     `json:"rain1h"`
	HourOffset int
}

type WeatherApiForecast struct {
	Timestamp  uint    `json:"timestamp"`
	Code       int     `json:"code"`
	MinTemp    float32 `json:"min_temp"`
	MaxTemp    float32 `json:"max_temp"`
	Rain1h     int     `json:"rain1h"`
	HourOffset int
}

type WeatherApiManager interface {
	AsyncRequest(lat float32, lon float32) (*WeatherApiCommon, []*WeatherApiForecast, []*WeatherApiCommon, error)
}

type WeatherApiManagerImpl struct {
	ApiKey string
	Url    string
}

func (w *WeatherApiManagerImpl) GetCurrentInfo(lat float32, lon float32) (*WeatherApiCommon, error) {
	result := &WeatherApiCommon{}

	query := url.Values{}
	query.Add("api_key", w.ApiKey)
	query.Add("lat", fmt.Sprintf("%.1f", lat))
	query.Add("lon", fmt.Sprintf("%.1f", lon))

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

func (w *WeatherApiManagerImpl) GetForecastInfo(lat float32, lon float32, hourOffset int) (*WeatherApiForecast, error) {
	result := &WeatherApiForecast{}

	query := url.Values{}
	query.Add("api_key", w.ApiKey)
	query.Add("lat", fmt.Sprintf("%.1f", lat))
	query.Add("lon", fmt.Sprintf("%.1f", lon))
	query.Add("hour_offset", strconv.Itoa(hourOffset))

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

func (w *WeatherApiManagerImpl) GetHistoricalInfo(lat float32, lon float32, hourOffset int) (*WeatherApiCommon, error) {
	result := &WeatherApiCommon{}

	query := url.Values{}
	query.Add("api_key", w.ApiKey)
	query.Add("lat", fmt.Sprintf("%.1f", lat))
	query.Add("lon", fmt.Sprintf("%.1f", lon))
	query.Add("hour_offset", strconv.Itoa(hourOffset))

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

func (w *WeatherApiManagerImpl) AsyncRequest(lat float32, lon float32) (*WeatherApiCommon, []*WeatherApiForecast, []*WeatherApiCommon, error) {
	var currentWeather *WeatherApiCommon
	var forecastWeather []*WeatherApiForecast
	var historicalWeather []*WeatherApiCommon

	var eg errgroup.Group
	eg.Go(func() error {
		var err error
		currentWeather, err = w.GetCurrentInfo(lat, lon)
		if err != nil {
			return err
		}
		return nil
	})

	for fh := range forecastHourOffset {
		index := fh
		eg.Go(func() error {
			fw, err := w.GetForecastInfo(lat, lon, forecastHourOffset[index])
			if err != nil {
				return err
			}
			forecastWeather = append(forecastWeather, fw)
			return nil
		})
	}

	for hh := range historicalHourOffset {
		index := hh
		eg.Go(func() error {
			hw, err := w.GetHistoricalInfo(lat, lon, historicalHourOffset[index])
			if err != nil {
				return err
			}
			historicalWeather = append(historicalWeather, hw)
			return nil
		})
	}

	// request 중 error 발생 시 모든 고루틴 종료를 위함.
	if err := eg.Wait(); err != nil {
		return nil, nil, nil, err
	}

	return currentWeather, forecastWeather, historicalWeather, nil

}
