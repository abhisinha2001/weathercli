package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/abhisinha2001/weathercli/models"
)

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

	weather, err := models.GetWeather(city, apiKey)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)

	}
	fmt.Printf("Weather in %s: %.1f°C, %s\n", weather.Name, weather.Main.Temp, weather.Weather[0].Description)

}
