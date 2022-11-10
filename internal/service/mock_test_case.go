package service

import (
	"github.com/closetotheworld/go-weather-service/pkg/weather"
)

var (
	forecastHourOffset   = []string{"6", "12", "18", "24", "30", "36", "42", "48"}
	historicalHourOffset = []string{"-6", "-12", "-18", "-24"}
)

type WeatherCase struct {
	Current    *weather.WeatherApiCommon
	Forecast   []*weather.WeatherApiForecast
	Historical []*weather.WeatherApiCommon
}

func (w *WeatherCase) CreateForTestCurrent() {
	for i := range forecastHourOffset {
		w.Forecast = append(w.Forecast, &weather.WeatherApiForecast{
			Code:       CODE_RAINY,
			HourOffset: forecastHourOffset[i],
		})
	}
	for j := range historicalHourOffset {
		w.Historical = append(w.Historical, &weather.WeatherApiCommon{
			Temp:       27.5,
			HourOffset: historicalHourOffset[j],
		})
	}
}

func (w *WeatherCase) CreateForTestForecast() {
	w.Current = &weather.WeatherApiCommon{
		Code:   CODE_SUNNY,
		Temp:   20.5,
		Rain1h: 0,
	}

	for i := range forecastHourOffset {
		w.Forecast = append(w.Forecast, &weather.WeatherApiForecast{
			Code:       CODE_SNOWY,
			HourOffset: forecastHourOffset[i],
		})
	}

	for j := range historicalHourOffset {
		w.Historical = append(w.Historical, &weather.WeatherApiCommon{
			Temp:       27.5,
			HourOffset: historicalHourOffset[j],
		})
	}
}

func (w *WeatherCase) CreateForTestHistorical() {
	w.Current = &weather.WeatherApiCommon{
		Code:   CODE_SUNNY,
		Temp:   20.5,
		Rain1h: 0,
	}

	for i := range forecastHourOffset {
		w.Forecast = append(w.Forecast, &weather.WeatherApiForecast{
			Code:       CODE_RAINY,
			HourOffset: forecastHourOffset[i],
		})
	}
	for j := range historicalHourOffset {
		if historicalHourOffset[j] == "-24" {
			break
		}
		w.Historical = append(w.Historical, &weather.WeatherApiCommon{
			Temp:       27.5,
			HourOffset: historicalHourOffset[j],
		})
	}
	w.Historical[1].Temp = 10.1
}

func GetGreetingTestCase(caseNum int) (*WeatherCase, string) {
	wc := WeatherCase{}
	wc.CreateForTestCurrent()
	answer := ""

	switch caseNum {
	case 0:
		wc.Current = &weather.WeatherApiCommon{Code: CODE_SNOWY, Temp: -10, Rain1h: 100}
		answer = WEATHER_SNOWY_PRECIPITATION_OVER
	case 1:
		wc.Current = &weather.WeatherApiCommon{Code: CODE_SNOWY, Temp: -10, Rain1h: 50}
		answer = WEATHER_SNOWY
	case 2:
		wc.Current = &weather.WeatherApiCommon{Code: CODE_RAINY, Temp: -10, Rain1h: 100}
		answer = WEATHER_RAINY_PRECIPITATION_OVER
	case 3:
		wc.Current = &weather.WeatherApiCommon{Code: CODE_RAINY, Temp: -10, Rain1h: 50}
		answer = WEATHER_RAINY
	case 4:
		wc.Current = &weather.WeatherApiCommon{Code: CODE_BLUR, Temp: -10, Rain1h: 50}
		answer = WEATHER_BLUR
	case 5:
		wc.Current = &weather.WeatherApiCommon{Code: CODE_SUNNY, Temp: 30, Rain1h: 100}
		answer = WEATHER_SUNNY_CELSIUS_OVER
	case 6:
		wc.Current = &weather.WeatherApiCommon{Code: CODE_SUNNY, Temp: -10, Rain1h: 100}
		answer = WEATHER_SUNNY_CELSIUS_UNDER
	case 7:
		wc.Current = &weather.WeatherApiCommon{Code: CODE_SUNNY, Temp: 15, Rain1h: 100}
		answer = WEATHER_SUNNY
	}
	return &wc, answer
}

func GetTempertureTestCase(caseNum int) (*WeatherCase, string) {
	wc := WeatherCase{}
	wc.CreateForTestHistorical()
	answer := ""

	switch caseNum {
	case 0:
		wc.Historical = append(wc.Historical, &weather.WeatherApiCommon{Temp: 30.5, HourOffset: "-24"})
		answer = celsiusLowerOver(10)
	case 1:
		wc.Current.Temp = 10.5
		wc.Historical = append(wc.Historical, &weather.WeatherApiCommon{Temp: 20.5, HourOffset: "-24"})
		answer = celsiusLowerUnder(10)
	case 2:
		wc.Historical = append(wc.Historical, &weather.WeatherApiCommon{Temp: 10.5, HourOffset: "-24"})
		answer = celsiusUpperOver(10)
	case 3:
		wc.Current.Temp = 10.5
		wc.Historical = append(wc.Historical, &weather.WeatherApiCommon{Temp: 0.5, HourOffset: "-24"})
		answer = celsiusUpperUnder(10)
	case 4:
		wc.Historical = append(wc.Historical, &weather.WeatherApiCommon{Temp: 20.5, HourOffset: "-24"})
		answer = celsiusSame(wc.Current.Temp)
	case 5:
		wc.Current.Temp = 10.5
		wc.Historical = append(wc.Historical, &weather.WeatherApiCommon{Temp: 10.5, HourOffset: "-24"})
		answer = celsiusSame(wc.Current.Temp)
	}
	answer = answer + " " + minMaxCelsius(27.5, 10.1)
	return &wc, answer
}

func GetHeadsUpTestCase(caseNum int) (*WeatherCase, string) {
	wc := WeatherCase{}
	wc.CreateForTestForecast()
	answer := ""

	switch caseNum {
	case 0:
		answer = WEATHER_EXPECT_HEAVY_SNOW
	case 1:
		for i := 0; i < 4; i++ {
			wc.Forecast[i].Code = CODE_SUNNY
		}
		answer = WEATHER_EXPECT_SNOW
	case 2:
		for i := 0; i < 4; i++ {
			wc.Forecast[i].Code = CODE_RAINY
			wc.Forecast[i+4].Code = CODE_SUNNY
		}
		answer = WEATHER_EXPECT_HEAVY_RAIN
	case 3:
		for i := 0; i < 4; i++ {
			wc.Forecast[i].Code = CODE_SUNNY
			wc.Forecast[i+4].Code = CODE_RAINY
		}
		answer = WEATHER_EXPECT_RAIN
	case 4:
		for i := 0; i < 7; i++ {
			wc.Forecast[i].Code = CODE_SUNNY
		}
		answer = WEATHER_EXPECT_COMMON
	}
	return &wc, answer
}
