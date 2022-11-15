package model

type WeatherResult struct {
	Summary Weather `json:"summary"`
}

type Weather struct {
	Greeting    string `json:"greeting"`
	Temperature string `json:"temperature"`
	HeadsUp     string `json:"heads-up"`
}
