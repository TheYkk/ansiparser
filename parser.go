package ansiparser

import (
	"fmt"
	"regexp"
)

var ansiRegex = regexp.MustCompile(`\x1b\[[0-9;]*m`)
var ansiBlockRegex = regexp.MustCompile(`\x1b\[[0-9;]*m(.*)\x1b\[[0-9;]*m`)

func Parser(buf []byte) string {
	matches := ansiRegex.FindAll(buf, -1)

	for i, match := range matches {
		fmt.Println(i, match)
		if match[3] == 48 {
			// Clear
		}
	}
	return ""
}
