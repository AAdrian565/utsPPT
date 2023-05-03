package main

import (
	"net/http"
	"pages"
)

type Person struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func main() {
	http.HandleFunc("/input", pages.InputJSON)
	http.HandleFunc("/raw", pages.Raw)
	http.HandleFunc("/", pages.View)
	http.ListenAndServe(":8080", nil)
}
