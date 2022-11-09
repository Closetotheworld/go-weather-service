package weather

type MockWeatherApiImpl struct {
}

func (m *MockWeatherApiImpl) GetCurrentInfo(lat string, lon string) (*WeatherApiCommon, error) {
	return nil, nil
}

func (m *MockWeatherApiImpl) GetForecastInfo(lat string, lon string, hourOffset string) (*WeatherApiForecast, error) {
	return nil, nil
}

func (m *MockWeatherApiImpl) GetHistoricalInfo(lat string, lon string, hourOffset string) (*WeatherApiCommon, error) {
	return nil, nil
}

func (m *MockWeatherApiImpl) AsyncRequest(lat string, lon string) (*WeatherApiCommon, []*WeatherApiForecast, []*WeatherApiCommon) {
	return nil, nil, nil
}
