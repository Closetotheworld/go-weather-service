package model

type WeatherResult struct {
	Summary Weather `json:"summary"`
}

type Weather struct {
	Greeting   string `json:"gretting"`
	Temperture string `json:"temperture"`
	HeadsUp    string `json:"heads_up"`
}
