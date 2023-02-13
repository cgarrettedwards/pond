package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	apiKey := os.Getenv("OPENWEATHER_APIKEY")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What city would you like to grab weather for?")
	city, err := reader.ReadString('\n')
	city = strings.Trim(city, "\n")
	if err != nil {
		panic(err)
	}

	getCoordinates(city, apiKey)
	getWeather(43.444, 32.3222, apiKey)
	getWeather(43.323, 2.34, apiKey)
	getWeather(99.99, 102.45, apiKey)
	getWeather(63, 11, apiKey)

}

func getWeather(lat float64, long float64, apiKey string) {
	fmt.Printf("Calling OpenWeather API for %f, %f\n", lat, long)

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat={%f}&lon={%f}&appid={%s}", lat, long, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	if resp.StatusCode != 200 {
		fmt.Println(resp.Status)
		return
	}

	var responseData interface{}

	err = json.NewDecoder(resp.Body).Decode(responseData)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", responseData)
}

func getCoordinates(city string, apiKey string) {
	fmt.Printf("Calling Geocoding API for %s\n", city)

	url := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q={%s}&appid={%s}", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	if resp.StatusCode != 200 {
		fmt.Println(resp.Status)
		return
	}

	var responseData interface{}

	err = json.NewDecoder(resp.Body).Decode(responseData)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", responseData)
}
