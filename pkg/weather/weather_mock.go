package weather

type MockWeatherApiImpl struct {
}

func (m *MockWeatherApiImpl) GetCurrentInfo(lat float32, lon float32) (*WeatherApiCommon, error) {

}

func (m *MockWeatherApiImpl) GetForecastInfo(lat float32, lon float32, hourOffset int) (*WeatherApiForecast, error) {

}

func (m *MockWeatherApiImpl) GetHistoricalInfo(lat float32, lon float32, hourOffset int) (*WeatherApiCommon, error) {

}
