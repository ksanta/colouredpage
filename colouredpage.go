package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// This starts a web server and displays the page background colour as the configured colour.
// The background colour is given as the first argument to the app, and should be valid CSS colour

var templates = template.Must(template.ParseFiles("page.html"))

var colour string

func main() {
	if len(os.Args) == 1 {
		log.Fatal("Cannot start. Must specify a colour as the command argument.")
	}

	colour = os.Args[1]
	log.Printf("Starting web server with colour %s", colour)

	http.HandleFunc("/", handleAllRequests)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleAllRequests(writer http.ResponseWriter, _ *http.Request) {
	randomSleep()
	err := templates.Execute(writer, colour)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func randomSleep() {
	duration := time.Duration(rand.Int31n(100)) * time.Millisecond
	log.Println("Sleeping for", duration)
	time.Sleep(duration)
}
