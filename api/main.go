package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os/exec"
	"sathvik.ninja/hemgissa/model"
	"sathvik.ninja/hemgissa/scraper"
)

type Request struct {
	Query string
}

type Response struct {
	EstimatedPrice float64 `json:"estimatedPrice"`
}

var predictor *model.Predictor

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
	props := scraper.Scrape(req.Query)
	// TODO: Out of index check
	estimatedPrice := predictor.PredictPrice(props).Result[0]
	roundUpPrice := (math.Round(estimatedPrice*100) / 100) * 1000000
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(Response{
		roundUpPrice,
	})
}

func startDeps() {
	go func() {
		err := exec.Command("python3", "model_server.py").Start()
		if err != nil {
			log.Fatal(err)
		}
	}()
	
}

func main() {
	startDeps()
	var err error
	predictor, err = model.NewPredictor("http://localhost:5000/predict", "model/regions.json")
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", handler)
	http.HandleFunc("/predict", predict)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
