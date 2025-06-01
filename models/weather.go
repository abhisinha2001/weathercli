package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WeatherModel struct {
	Name string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func GetWeather(city, apiKey string) (*WeatherModel, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to get weather: status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	var weatherResp WeatherModel
	if err := json.Unmarshal(body, &weatherResp); err != nil {
		return nil, err
	}
	return &weatherResp, nil
}
