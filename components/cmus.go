package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	response, _ := exec.Command("cmus-remote", "-Q").Output()

	artistPattern, _ := regexp.Compile("tag artist .*")
	titlePattern, _ := regexp.Compile("tag title .*")

	artist := strings.Trim(artistPattern.FindString(string(response)), "tag artist")
	title := strings.Trim(titlePattern.FindString(string(response)), "tag title")

	if artist != "" && title != "" {
		fmt.Println(fmt.Sprintf("%s - %s", artist, title))
	}
}