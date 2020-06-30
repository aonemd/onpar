package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

const (
	// find an entry that has `headphone jack` in it using `amixer controls`
	headphone_numid = "1"
	sound_card      = "0"

	speaker_up_icon     = ""
	speaker_mute_icon   = " MUTE"
	headphone_up_icon   = ""
	headphone_mute_icon = " MUTE"
)

func main() {
	response, err := exec.Command("amixer", "sget", "Master").Output()
	if err != nil {
		return
	}

	levelPattern, _ := regexp.Compile("[0-9]+%")
	statusPattern, _ := regexp.Compile("\\[(on|off)\\]")

	level := levelPattern.FindString(string(response))
	status := strings.Trim(statusPattern.FindString(string(response)), "[]")

	headPhoneStatusArgs := strings.Split(fmt.Sprintf("-c %s cget numid=%s", sound_card, headphone_numid), " ")
	headphoneResponse, _ := exec.Command("amixer", headPhoneStatusArgs...).Output()
	headPhoneStatusPattern, _ := regexp.Compile("values=(on|off)")
	headPhoneStatus := strings.Trim(headPhoneStatusPattern.FindString(string(headphoneResponse)), "values=")

	upIcon := speaker_up_icon
	muteIcon := speaker_mute_icon

	if headPhoneStatus == "on" {
		upIcon = headphone_up_icon
		muteIcon = headphone_mute_icon
	}

	if status == "off" {
		fmt.Println(muteIcon)
	} else {
		fmt.Println(fmt.Sprintf("%s %s", upIcon, level))
	}
}
