/*
Copyright © 2025 Abhinav Sinha abhisinha0601@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/abhisinha2001/weathercli/models"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var units string

// weatherCmd represents the weather command
var weatherCmd = &cobra.Command{
	Use:   "weather <city>",
	Short: "Get current weather for a city",

	Run: func(cmd *cobra.Command, args []string) {
		_ = godotenv.Load()
		apiKey := os.Getenv("API_KEY")

		if apiKey == "" {
			fmt.Println("API Key does not exist. Please check the .env file")
			os.Exit(1)
		}

		city := strings.Join(args, " ")

		fmt.Printf("Fetching weather for %s...\n", city)

		weather, err := models.GetWeather(city, apiKey, units)
		if err != nil {
			fmt.Printf("Error: %v/n", err)
			os.Exit(1)
		}

		if len(weather.Weather) == 0 {
			fmt.Println("No weather data found")
			os.Exit(1)
		}

		fmt.Printf("Weather in %s: %.1f°C, %s \n", weather.Name, weather.Main.Temp, weather.Weather[0].Description)
	},
}

func init() {
	rootCmd.AddCommand(weatherCmd)

	weatherCmd.Flags().StringVarP(&units, "units", "u", "metric", "Unists of measurement: standard, metric,imperial")

}
