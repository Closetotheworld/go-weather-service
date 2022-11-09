package weather

type WeatherApiCommon struct {
	Timestamp uint    `json:"timestamp"`
	Code      int     `json:"code"`
	Temp      float32 `json:"temp"`
	Rain1h    int     `json:"rain1h"`
}

type WeatherApiForecast struct {
	Timestamp uint    `json:"timestamp"`
	Code      int     `json:"code"`
	MinTemp   float32 `json:"min_temp"`
	MaxTemp   float32 `json:"max_temp"`
	Rain1h    int     `json:"rain1h"`
}

type WeatherApiManager interface {
	GetCurrentInfo(lat float32, lon float32) (*WeatherApiCommon, error)
	GetForecastInfo(lat float32, lon float32, hourOffset int) (*WeatherApiForecast, error)
	GetHistoricalInfo(lat float32, lon float32, hourOffset int) (*WeatherApiCommon, error)
}

type WeatherApiManagerImpl struct {
	ApiKey string
	Lat    float32
	Lon    float32
}

func (w *WeatherApiManagerImpl) GetCurrentInfo(lat float32, lon float32) (*WeatherApiCommon, error) {

}

func (w *WeatherApiManagerImpl) GetForecastInfo(lat float32, lon float32, hourOffset int) (*WeatherApiForecast, error) {

}

func (w *WeatherApiManagerImpl) GetHistoricalInfo(lat float32, lon float32, hourOffset int) (*WeatherApiCommon, error) {

}
