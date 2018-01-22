package main

import (
	"fmt"
	"os/exec"
	"regexp"
)

func main() {
	response, _ := exec.Command("amixer", "sget", "Master").Output()
	levelPattern, _ := regexp.Compile("[0-9]+%")
	level := levelPattern.FindString(string(response))

	fmt.Println(fmt.Sprintf("VOL: %s", level))
}