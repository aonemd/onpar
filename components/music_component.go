package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	var _response []byte
	var _err error

	_response, _err = exec.Command("playerctl", "metadata", "artist").Output()
	if _err != nil {
		return
	}
	artist := strings.TrimSpace(string(_response))

	_response, _err = exec.Command("playerctl", "metadata", "title").Output()
	if _err != nil {
		return
	}
	title := strings.TrimSpace(string(_response))

	if artist != "" && title != "" {
		fmt.Println(fmt.Sprintf(" %s - %s", artist, title))
	}
}
