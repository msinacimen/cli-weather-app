package cmd

import (
	"fmt"
	"github.com/brettski/go-termtables"
	"github.com/spf13/cobra"
	"math"
	"strconv"
	"wcli/utils"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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
		result := utils.GetCity(city, utils.LoadKey())

		description := cases.Title(language.Und).String(result.Weather[0].Description)

		temp := strconv.Itoa(int(math.Round(result.Main.Temp)))
		feels_like := strconv.Itoa(int(math.Round(result.Main.FeelsLike)))

		min := strconv.Itoa(int(math.Round(result.Main.TempMin)))
		max := strconv.Itoa(int(math.Round(result.Main.TempMax)))

		humidity := strconv.Itoa(int(math.Round(result.Main.Humidity)))
		pressure := strconv.Itoa(int(math.Round(result.Main.Pressure)))

		table := termtables.CreateTable()

		table.AddHeaders("City", "Weather", "Temperature", "Feels Like", "Min/Max Temperature", "Humidity", "Pressure")
		table.AddRow(city, description, temp+"°C", feels_like+"°C", min+" - "+max+"°C", humidity+"%", pressure+"mbar")

		fmt.Println(table.Render())

	},
}

func init() {
	rootCmd.AddCommand(cityCmd)
}
