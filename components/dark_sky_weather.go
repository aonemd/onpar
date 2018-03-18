package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
)

const (
	api_key = "api_key"
	lon_lat = "lon,lat"
	units   = "si"

	clear_day_icon           = ""
	clear_night_icon         = ""
	partly_cloudy_day_icon   = ""
	partly_cloudy_night_icon = ""
	cloudy_icon              = ""
	rain_icon                = ""
	sleet_icon               = ""
	snow_icon                = ""
	wind_icon                = ""
	fog_icon                 = ""
	celsius_icon             = "°C"
	fahrenheit_icon          = "°F"
)

func main() {
	api_uri := fmt.Sprintf("https://api.darksky.net/forecast/%s/%s?exclude=minutely,hourly,daily,alerts,flags&units=%s", api_key, lon_lat, units)
	response, _ := http.Get(api_uri)
	data, _ := ioutil.ReadAll(response.Body)

	var forecast interface{}
	json.Unmarshal(data, &forecast)
	forecastMap := forecast.(map[string]interface{})
	currentForcast := forecastMap["currently"].(map[string]interface{})

	summary := currentForcast["summary"]
	icon := currentForcast["icon"]
	raw_temperature := currentForcast["temperature"].(float64)
	temperature := math.Round(raw_temperature)

	switch icon {
	case "clear-day":
		icon = clear_day_icon
	case "clear-night":
		icon = clear_night_icon
	case "partly-cloudy-day":
		icon = partly_cloudy_day_icon
	case "partly-cloudy-night":
		icon = partly_cloudy_night_icon
	case "cloudy":
		icon = cloudy_icon
	case "rain":
		icon = rain_icon
	case "sleet":
		icon = sleet_icon
	case "snow":
		icon = snow_icon
	case "wind":
		icon = wind_icon
	case "fog":
		icon = fog_icon
	}

	if units == "si" {
		fmt.Println(fmt.Sprintf("%s %v%s %s", icon, temperature, celsius_icon, summary))
	} else {
		fmt.Println(fmt.Sprintf("%s %v%s %s", icon, temperature, fahrenheit_icon, summary))
	}
}
