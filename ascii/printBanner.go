package main

import (
	"fmt"
)

func PrintBanner(words []string, Bannerlines []string) {
	for _, word := range words {
		if word == "" {
			fmt.Println()
			continue
		}
		for i := 0; i < 8; i++ {
			for _, ch := range word {
				if ch < 32 || ch > 126 {
					ch = '?'
				}
				id := (int(ch-32)*9 + 1)
				fmt.Print(Bannerlines[id+i])
			}
			fmt.Println()
		}
	}
}
