// requires `acpi`

package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
)

const (
	battery        = "0"
	critical_level = 10

	charging_icon      = ""
	full_icon          = " "
	three_quarter_icon = ""
	half_icon          = ""
	empty_icon         = ""
)

func main() {
	response, _ := exec.Command("acpi", "-b", " | ", "grep", "Battery ", battery).Output()

	iconPattern, _ := regexp.Compile("Full|Charging|Discharging")
	powerPattern, _ := regexp.Compile("[0-9]+%")
	remainingTimePattern, _ := regexp.Compile("[01][0-9]:[0-9][0-9]")

	icon := iconPattern.FindString(string(response))
	power, _ := strconv.Atoi(powerPattern.FindString(string(response)))
	remainingTime := remainingTimePattern.FindString(string(response))

	switch icon {
	case "Full":
		icon = full_icon
	case "Charging":
		icon = charging_icon
	case "Discharging":
		if power <= critical_level {
			icon = empty_icon
		} else if power <= 50 {
			icon = half_icon
		} else if power <= 75 {
			icon = three_quarter_icon
		} else {
			icon = full_icon
		}
	}

	fmt.Println(fmt.Sprintf("%s %v (%s)", icon, power, remainingTime))
}
