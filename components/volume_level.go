package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	response, _ := exec.Command("amixer", "sget", "Master").Output()

	levelPattern, _ := regexp.Compile("[0-9]+%")
	statusPattern, _ := regexp.Compile("\\[(on|off)\\]")

	level := levelPattern.FindString(string(response))
	status := strings.Trim(statusPattern.FindString(string(response)), "[]")

	if status == "off" {
		fmt.Println(" MUTE")
	} else {
		fmt.Println(fmt.Sprintf(" %s", level))
	}
}
