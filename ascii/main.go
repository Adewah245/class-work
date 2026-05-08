package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println(`usage: go run . "your text" bannerfiles names`)
		return
	}
	arg := os.Args[1]
	if arg == "" {
		return
	}

	bannerfiles := "standard.txt"
	if len(os.Args) == 3 {
		bannerfiles = os.Args[2]
	}

	filename := bannerfiles

	banner, err := GetBannerLines(filename)
	if err != nil {
		fmt.Println("Error: bannerfiles not found!", err)
		return
	}

	word := inputtext(arg)
	PrintBanner(word, banner)

}
