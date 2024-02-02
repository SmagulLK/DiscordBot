package models

type WeatherResponse struct {
	Main struct {
		Temperature float64 `json:"temp"`
	} `json:"main"`
}
