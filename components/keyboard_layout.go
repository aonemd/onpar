package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	response, err := exec.Command("setxkbmap", "-query").Output()
	if err != nil {
		return
	}

	layoutPattern, _ := regexp.Compile("layout:\\s+[a-z]+")
	layout := strings.TrimSpace(strings.Split(layoutPattern.FindString(string(response)), ":")[1])

	fmt.Println(fmt.Sprintf(" %s", layout[:2]))
}
