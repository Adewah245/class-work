package main

import (
	"os"
	"strings"
)

func GetBannerLines(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	step := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(step, "\n")

	return lines, nil
}
