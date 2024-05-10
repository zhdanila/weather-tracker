package models

type ApiConfigData struct {
	OpenWeatherMapApiKey string `json:"openWeatherMapApiKey"`
}

type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Celsius float64 `json:"temp"`
	} `json:"main"`
}
