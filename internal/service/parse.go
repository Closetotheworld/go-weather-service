package service

import (
	"github.com/closetotheworld/go-weather-service/pkg/weather"
)

func ParseGreeting(currentWeather weather.WeatherApiCommon) string {
	var result string
	isRainOver := false

	if currentWeather.Rain1h >= 100 {
		isRainOver = true
	}

	switch currentWeather.Code {
	case CODE_SUNNY:
		if currentWeather.Temp >= 30 {
			result = WEATHER_SUNNY_CELSIUS_OVER
		} else if currentWeather.Temp <= 0 {
			result = WEATHER_SUNNY_CELSIUS_UNDER
		}
	case CODE_BLUR:
		result = WEATHER_BLUR
	case CODE_RAINY:
		if isRainOver {
			result = WEATHER_RAINY_PRECIPITATION_OVER
		} else {
			result = WEATHER_RAINY
		}
	case CODE_SNOWY:
		if isRainOver {
			result = WEATHER_SNOWY_PRECIPITATION_OVER
		} else {
			result = WEATHER_SNOWY
		}
	}
	if result == "" {
		result = WEATHER_SUNNY
	}
	return result
}

func ParseTemperture() string {
	return "2"
}

func ParseHeadsUp() string {
	return "3"
}
