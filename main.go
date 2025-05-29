package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type weatherModel struct {
	Name string `json:"name`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"temp`
	Weather []struct {
		Description string `json:"description`
	} `json:"weather"`
}

func getWeather(city, apiKey string) (*weatherModel, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Failed to get weather: status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	var weatherResp weatherModel
	if err := json.Unmarshal(body, &weatherResp); err != nil {
		return nil, err
	}
	return &weatherResp, nil
}

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	apiKey := os.Getenv("API_KEY")

	if apiKey == "" {
		fmt.Println("API_KEY is not set")
		os.Exit(1)
	}
	if len(os.Args) < 2 {
		fmt.Println("Usage: weathercli <city>")
		os.Exit(1)
	}

	city := os.Args[1]
	fmt.Printf("Fetching weather for %s\n", city)

	weather, err := getWeather(city, apiKey)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)

	}
	fmt.Printf("Weather in %s: %.1f°C, %s\n", weather.Name, weather.Main.Temp, weather.Weather[0].Description)

}
