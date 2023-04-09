package main

import (
	"cli-weather-app/utils"
	"fmt"
)

func main() {
	//cmd.Execute()
	city := "New York"
	rtr := utils.GetCity(city, utils.LoadKey())
	fmt.Printf("%s", rtr)
}
