package main

import (
	"log"
	"os/exec"
	"strings"
	"time"
)

var (
	status = []string{}
)

type Component struct {
	path     string
	Output   string
	interval int
	Channel  chan string
}

func NewComponent(path string, interval int) *Component {
	return &Component{path: path, interval: interval, Channel: make(chan string)}
}

func (c *Component) Run() {
	time.Sleep(time.Duration(c.interval) * time.Second)
	response, err := exec.Command(c.path).Output()
	if err != nil {
		log.Fatal(err)
	}

	c.Channel <- strings.TrimSpace(string(response))
}

func main() {
	dateTime := NewComponent("components/date_time.sh", 1)
	batteryLevel := NewComponent("components/battery_level", 1)
	volumeLevel := NewComponent("components/volume_level", 1)
	darkSkyWeather := NewComponent("components/dark_sky_weather", 600)
	keyboardLayout := NewComponent("components/keyboard_layout", 1)
	// Initialize new components here ...

	go dateTime.Run()
	go batteryLevel.Run()
	go volumeLevel.Run()
	go darkSkyWeather.Run()
	go keyboardLayout.Run()
	// Call new components Run() here ...

	for {
		select {
		case dateTime.Output = <-dateTime.Channel:
			go dateTime.Run()
		case darkSkyWeather.Output = <-darkSkyWeather.Channel:
			go darkSkyWeather.Run()
		case keyboardLayout.Output = <-keyboardLayout.Channel:
			go keyboardLayout.Run()
		case batteryLevel.Output = <-batteryLevel.Channel:
			go batteryLevel.Run()
		case volumeLevel.Output = <-volumeLevel.Channel:
			go volumeLevel.Run()
			// Call new components Run() here ...
		}

		status = []string{
			"",
			// Add new components here ...
			keyboardLayout.Output,
			darkSkyWeather.Output,
			volumeLevel.Output,
			batteryLevel.Output,
			dateTime.Output,
		}

		exec.Command("xsetroot", "-name", strings.Join(status, " ")).Run()
	}
}
