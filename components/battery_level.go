package main

import (
	"fmt"
	"os/exec"
	"regexp"
)

func main() {
	response, _ := exec.Command("acpi", "-b", " | ", "grep", "Battery ", "0").Output()

	// statePattern, _ := regexp.Compile("Full|Charging|Discharging")
	powerPattern, _ := regexp.Compile("[0-9]+%")
	remainingTimePattern, _ := regexp.Compile("[01][0-9]:[0-9][0-9]")

	// state := statePattern.FindString(string(response))
	power := powerPattern.FindString(string(response))
	remainingTime := remainingTimePattern.FindString(string(response))

	fmt.Println(fmt.Sprintf("BAT: %s (%s)", power, remainingTime))
}
