package main

import (
	"fmt"
	"os/exec"
	"regexp"
)

const (
	device_output = "Master"
)

func main() {
	response, _ := exec.Command("amixer", "sget", device_output).Output()
	levelPattern, _ := regexp.Compile("[0-9]+%")
	level := levelPattern.FindString(string(response))

	fmt.Println(fmt.Sprintf("VOL: %s", level))
}
