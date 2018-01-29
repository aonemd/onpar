package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	response, _ := exec.Command("setxkbmap", "-query").Output()
	layoutPattern, _ := regexp.Compile("layout:\\s+[a-z]+")
	layout := strings.TrimSpace(strings.Split(layoutPattern.FindString(string(response)), ":")[1])

	fmt.Println(fmt.Sprintf("%s", layout[:2]))
}
