package weather

import (
	"fmt"
	"sort"
	"strings"

	// external packages
	"github.com/closetotheworld/go-weather-service/pkg/weather_api"
	// project packages
)

func ParseGreeting(currentWeather weather_api.WeatherApiCommon) string {
	var result string
	isRainOver := false

	if currentWeather.Rain1h >= WEATHER_GREETING_RAIN_STANDARD {
		isRainOver = true
	}

	switch currentWeather.Code {
	case CODE_SUNNY:
		if currentWeather.Temp >= WEATHER_GREETING_CELSIUS_STANDARD {
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

func ParseTemperture(currentTemp float32, historicalWeather []*weather_api.WeatherApiCommon) string {
	var minTemp float32
	var maxTemp float32
	var yDayTemp float32
	answer := ""

	for i := range historicalWeather {
		if historicalWeather[i].HourOffset == -24 {
			yDayTemp = historicalWeather[i].Temp
		}

		if i == 0 {
			minTemp = historicalWeather[i].Temp
			maxTemp = historicalWeather[i].Temp
			continue
		}
		if minTemp < historicalWeather[i].Temp {
			minTemp = historicalWeather[i].Temp
		}
		if maxTemp > historicalWeather[i].Temp {
			maxTemp = historicalWeather[i].Temp
		}
	}

	tempDiff := currentTemp - yDayTemp
	switch {
	case tempDiff < 0:
		tempDiff = tempDiff * -1
		if currentTemp > WEATHER_TEMPERTURE_CELSIUS_STANDARD {
			answer = celsiusLowerOver(tempDiff)
		} else {
			answer = celsiusLowerUnder(tempDiff)
		}
	case tempDiff > 0:
		if currentTemp > WEATHER_TEMPERTURE_CELSIUS_STANDARD {
			answer = celsiusUpperOver(tempDiff)
		} else {
			answer = celsiusUpperUnder(tempDiff)
		}
	default:
		answer = celsiusSame(currentTemp)
	}

	return strings.Join([]string{answer, minMaxCelsius(minTemp, maxTemp)}, " ")
}

func ParseHeadsUp(forecastWeather []*weather_api.WeatherApiForecast) string {
	mapWeather := map[int]int{
		CODE_SNOWY: 0,
		CODE_RAINY: 0,
	}
	answer := ""
	fw := forecastWeather

	sort.Slice(fw, func(i, j int) bool {
		return fw[i].HourOffset < fw[j].HourOffset
	})
	for i := range fw {
		fmt.Println(fw[i].HourOffset)
		mapWeather[fw[i].Code] += 6
		if fw[i].HourOffset == 24 {
			if mapWeather[CODE_SNOWY] >= 12 {
				answer = WEATHER_EXPECT_HEAVY_SNOW
				return answer
			}
			if mapWeather[CODE_RAINY] >= 12 {
				answer = WEATHER_EXPECT_HEAVY_RAIN
			}
		}
	}
	if mapWeather[CODE_SNOWY] >= 12 {
		answer = WEATHER_EXPECT_SNOW
	}
	if answer == "" {
		if mapWeather[CODE_RAINY] >= 12 {
			answer = WEATHER_EXPECT_RAIN
		} else {
			answer = WEATHER_EXPECT_COMMON
		}
	}
	return answer
}
