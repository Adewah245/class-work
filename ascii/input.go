package main

import (
	"strings"
)

func inputtext(input string) []string {
	if input == "\\n" {
		return []string{""}
	}
	word := strings.Split(input, "\\n")
	return word
}
