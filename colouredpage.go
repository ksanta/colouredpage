package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

// This starts a web server and displays the page background colour as the configured colour.
// The background coloud is given as the first argument to the app, and should be valid CSS colour

var templates = template.Must(template.ParseFiles("page.html"))

var colour string

func main() {
	colour = os.Args[1]
	log.Printf("Starting web server with colour %s", colour)
	http.HandleFunc("/", webHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func webHandler(writer http.ResponseWriter, _ *http.Request) {
	err := templates.Execute(writer, colour)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
