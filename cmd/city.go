package cmd

import (
	"cli-weather-app/utils"
	"fmt"
	"github.com/spf13/cobra"
)

var cityName string

// cityCmd represents the city command
var cityCmd = &cobra.Command{
	Use:     "city",
	Short:   "Get temperature of a city",
	Long:    `This command fetches the temperature from the openWeatherMap API.`,
	Version: "0.1.0",
	Example: `wcli city <CITY_NAME>`,
	Run: func(cmd *cobra.Command, args []string) {
		var city string
		if cityName != "" {
			city = cityName
		} else {
			city = args[0] // check for validation
		}
		fmt.Printf("Weather at %s is %s", city, utils.GetCity(city, utils.LoadKey()))
	},
}

func init() {
	rootCmd.AddCommand(cityCmd)
}
