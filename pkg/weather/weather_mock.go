package weather

type MockWeatherApiImpl struct {
}

func (m *MockWeatherApiImpl) GetCurrentInfo(lat float32, lon float32) (*WeatherApiCommon, error) {
	return nil, nil
}

func (m *MockWeatherApiImpl) GetForecastInfo(lat float32, lon float32, hourOffset int) (*WeatherApiForecast, error) {
	return nil, nil
}

func (m *MockWeatherApiImpl) GetHistoricalInfo(lat float32, lon float32, hourOffset int) (*WeatherApiCommon, error) {
	return nil, nil
}
