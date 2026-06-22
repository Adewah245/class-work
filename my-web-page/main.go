package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HomePageHandler)
	fmt.Println("server: Runing Successfully!")
	http.ListenAndServe(":8080", nil)
}
