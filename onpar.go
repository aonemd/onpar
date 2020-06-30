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
	dateTimeComponent := NewComponent("date_time_component.sh", 1)
	batteryLevelComponent := NewComponent("battery_level_component", 1)
	volumeLevelComponent := NewComponent("volume_level_component", 1)
	darkSkyWeatherComponent := NewComponent("dark_sky_weather_component", 600)
	keyboardLayoutComponent := NewComponent("keyboard_layout_component", 1)
	musicComponent := NewComponent("music_component", 2)
	// Initialize new components here ...

	go dateTimeComponent.Run()
	go batteryLevelComponent.Run()
	go volumeLevelComponent.Run()
	go darkSkyWeatherComponent.Run()
	go keyboardLayoutComponent.Run()
	go musicComponent.Run()
	// Call new components Run() here ...

	for {
		select {
		case dateTimeComponent.Output = <-dateTimeComponent.Channel:
			go dateTimeComponent.Run()
		case darkSkyWeatherComponent.Output = <-darkSkyWeatherComponent.Channel:
			go darkSkyWeatherComponent.Run()
		case keyboardLayoutComponent.Output = <-keyboardLayoutComponent.Channel:
			go keyboardLayoutComponent.Run()
		case batteryLevelComponent.Output = <-batteryLevelComponent.Channel:
			go batteryLevelComponent.Run()
		case volumeLevelComponent.Output = <-volumeLevelComponent.Channel:
			go volumeLevelComponent.Run()
		case musicComponent.Output = <-musicComponent.Channel:
			go musicComponent.Run()
			// Call new components Run() here ...
		}

		status = []string{
			"",
			// Add new components here ...
			musicComponent.Output,
			keyboardLayoutComponent.Output,
			darkSkyWeatherComponent.Output,
			volumeLevelComponent.Output,
			batteryLevelComponent.Output,
			dateTimeComponent.Output,
		}

		exec.Command("xsetroot", "-name", strings.Join(status, " ")).Run()
	}
}
