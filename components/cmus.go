package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	response, err := exec.Command("cmus-remote", "-Q").Output()
	if err != nil {
		return
	}

	artistPattern, _ := regexp.Compile("tag artist .*")
	titlePattern, _ := regexp.Compile("tag title .*")

	artist := strings.TrimPrefix(artistPattern.FindString(string(response)), "tag artist ")
	title := strings.TrimPrefix(titlePattern.FindString(string(response)), "tag title ")

	if artist != "" && title != "" {
		fmt.Println(fmt.Sprintf(" %s - %s", artist, title))
	}
}
