package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

var (
	status         = []string{}
	dateTime       string
	batteryLevel   string
	volumeLevel    string
	keyboardLayout string
)

func updateKeyboardLayout(c chan string) {
	time.Sleep(1 * time.Second)

	response, _ := exec.Command("setxkbmap", "-query").Output()
	levelPattern, _ := regexp.Compile("layout:\\s+[a-z]+")
	layout := strings.TrimSpace(strings.Split(levelPattern.FindString(string(response)), ":")[1])

	c <- fmt.Sprintf("KEY: %s", layout)
}

func updateVolumeLevel(c chan string) {
	time.Sleep(1 * time.Second)

	response, _ := exec.Command("amixer", "sget", "Master").Output()
	levelPattern, _ := regexp.Compile("[0-9]+%")
	level := levelPattern.FindString(string(response))

	c <- fmt.Sprintf("VOL: %s", level)
}

func updateBatteryLevel(c chan string) {
	time.Sleep(1 * time.Second)

	response, _ := exec.Command("acpi", "-b", "| grep", "Battery ", "0").Output()

	// statePattern, _ := regexp.Compile("Full|Charging|Discharging")
	powerPattern, _ := regexp.Compile("[0-9]+%")
	remainingTimePattern, _ := regexp.Compile("[01][0-9]:[0-9][0-9]")

	// state := statePattern.FindString(string(response))
	power := powerPattern.FindString(string(response))
	remainingTime := remainingTimePattern.FindString(string(response))

	c <- fmt.Sprintf("BAT: %s (%s)", power, remainingTime)
}

func updateDateTime(c chan string) {
	time.Sleep(1 * time.Second)

	c <- time.Now().Local().Format("Mon Jan 02 03:04 PM")
}

func main() {
	dateTimeChannel := make(chan string)
	batteryLevelChannel := make(chan string)
	volumeLevelChannel := make(chan string)
	keyboardLayoutChannel := make(chan string)

	go updateDateTime(dateTimeChannel)
	go updateBatteryLevel(batteryLevelChannel)
	go updateVolumeLevel(volumeLevelChannel)
	go updateKeyboardLayout(keyboardLayoutChannel)

	for {
		select {
		case dateTime = <-dateTimeChannel:
			status = []string{
				"",
				keyboardLayout,
				volumeLevel,
				batteryLevel,
				dateTime,
			}
			go updateDateTime(dateTimeChannel)
		case batteryLevel = <-batteryLevelChannel:
			status = []string{
				"",
				keyboardLayout,
				volumeLevel,
				batteryLevel,
				dateTime,
			}
			go updateBatteryLevel(batteryLevelChannel)
		case volumeLevel = <-volumeLevelChannel:
			status = []string{
				"",
				keyboardLayout,
				volumeLevel,
				batteryLevel,
				dateTime,
			}
			go updateVolumeLevel(volumeLevelChannel)
		case keyboardLayout = <-keyboardLayoutChannel:
			status = []string{
				"",
				keyboardLayout,
				volumeLevel,
				batteryLevel,
				dateTime,
			}
			go updateKeyboardLayout(keyboardLayoutChannel)
		}

		exec.Command("xsetroot", "-name", strings.Join(status, " ")).Run()
	}
}
