package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Company struct {
	Name    string
	CEO     string
	Founded int
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Adewah := Company{
		Name:    "Adewah Global Tech.",
		CEO:     "JAMES",
		Founded: 2026,
	}
	templ.Execute(w, Adewah)
}
func main() {
	http.HandleFunc("/", HomePageHandler)
	fmt.Println("server loaded successully: visit http//localhost:8080")
	http.ListenAndServe(":8080", nil)
}
