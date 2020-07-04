// requires `playerctl`

package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

var (
	status         []string
	filteredStatus []string
	_response      []byte
	_err           error
	player         string
	artist         string
	title          string
)

func main() {
	_response, _err = exec.Command("playerctl", "metadata").Output()
	if _err != nil {
		return
	}

	playerPattern, _ := regexp.Compile("^([\\w]+)")
	player = playerPattern.FindString(string(_response))

	artistPattern, _ := regexp.Compile(":artist\\s+(.*)")
	artistFoundPattern := artistPattern.FindStringSubmatch(string(_response))
	if artistFoundPattern != nil && player != "chrome" {
		artist = strings.TrimSpace(artistFoundPattern[1])
	}

	titlePattern, _ := regexp.Compile(":title\\s+(.*)")
	titleFoundPattern := titlePattern.FindStringSubmatch(string(_response))
	if titleFoundPattern != nil {
		title = strings.TrimSpace(titleFoundPattern[1])
	}

	status = []string{
		artist,
		title,
	}

	for _, stat := range status {
		if stat != "" {
			filteredStatus = append(filteredStatus, stat)
		}
	}

	fmt.Println(fmt.Sprintf(" %s", strings.Join(filteredStatus, " - ")))
}
