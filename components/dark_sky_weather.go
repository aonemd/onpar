package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	api_key = "api_key"
	lon_lat = "lon,lat"
	units   = "si"
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
	temperature := currentForcast["temperature"]

	fmt.Println(fmt.Sprintf("%s %v°C", summary, temperature))
}
