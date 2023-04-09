package utils

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
)

const baseUrl = "https://api.openweathermap.org"

type OWMResponse struct {
	Weather []struct {
		Description string
	}
	Main struct {
		Temp float64
	}
}

func LoadKey() string {
	// loading apiKey from .env file
	dotenverr := godotenv.Load()
	if dotenverr != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("OWM_API_KEY")
	return apiKey
}

func GetCity(CITY string, APIKEY string) string {
	URL := fmt.Sprintf("%s/data/2.5/weather?q=%s&units=metric&appid=%s", baseUrl, CITY, APIKEY)
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)

	}
	defer resp.Body.Close()

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var response OWMResponse
	jsonErr := json.Unmarshal(body, &response)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return response.Weather[0].Description
	//fmt.Println(response.Main.Temp)
}
