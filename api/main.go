package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sathvik.ninja/hemgissa/scraper"
)

type Request struct {
	Query string
}

func handler(w http.ResponseWriter, r *http.Request) {
	
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func predict(w http.ResponseWriter, r *http.Request) {
	var req Request
	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	scraper.Scrape(req.Query)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/predict", predict)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
