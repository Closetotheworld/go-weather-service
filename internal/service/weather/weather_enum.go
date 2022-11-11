package weather

import "fmt"

const (
	CODE_SUNNY = 0
	CODE_BLUR  = 1
	CODE_RAINY = 2
	CODE_SNOWY = 3

	WEATHER_GREETING_RAIN_STANDARD      = 100
	WEATHER_GREETING_CELSIUS_STANDARD   = 30
	WEATHER_TEMPERTURE_CELSIUS_STANDARD = 15

	WEATHER_SNOWY_PRECIPITATION_OVER = "폭설이 내리고 있어요."
	WEATHER_SNOWY                    = "눈이 포슬포슬 내립니다."
	WEATHER_RAINY_PRECIPITATION_OVER = "폭우가 내리고 있어요."
	WEATHER_RAINY                    = "비가 오고 있습니다."
	WEATHER_BLUR                     = "날씨가 약간은 칙칙해요."
	WEATHER_SUNNY_CELSIUS_OVER       = "따사로운 햇살을 맞으세요."
	WEATHER_SUNNY_CELSIUS_UNDER      = "날이 참 춥네요."
	WEATHER_SUNNY                    = "날씨가 참 맑습니다."

	WEATHER_EXPECT_HEAVY_SNOW = "내일 폭설이 내릴 수도 있으니 외출 시 주의하세요."
	WEATHER_EXPECT_SNOW       = "눈이 내릴 예정이니 외출 시 주의하세요."
	WEATHER_EXPECT_HEAVY_RAIN = "폭우가 내릴 예정이에요. 우산을 미리 챙겨두세요."
	WEATHER_EXPECT_RAIN       = "며칠동안 비 소식이 있어요."
	WEATHER_EXPECT_COMMON     = "날씨는 대체로 평온할 예정이에요."
)

var (
	minMaxCelsius = func(max float32, min float32) string {
		return fmt.Sprintf("최고기온은 %.1f도, 최저기온은 %.1f도 입니다.", max, min)
	}
	celsiusLowerOver  = func(celsius float32) string { return fmt.Sprintf("어제보다 %.1f도 덜 덥습니다.", celsius) }
	celsiusLowerUnder = func(celsius float32) string { return fmt.Sprintf("어제보다 %.1f도 더 춥습니다.", celsius) }
	celsiusUpperOver  = func(celsius float32) string { return fmt.Sprintf("어제보다 %.1f도 더 덥습니다.", celsius) }
	celsiusUpperUnder = func(celsius float32) string { return fmt.Sprintf("어제보다 %.1f도 덜 춥습니다.", celsius) }
	celsiusSame       = func(celsius float32) string {
		if celsius > 15 {
			return "어제와 비슷하게 덥습니다."
		} else {
			return "어제와 비슷하게 춥습니다."
		}
	}
)
