package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly/v2"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"os/exec"
	"sathvik.ninja/hemgissa/model"
	"sathvik.ninja/hemgissa/scraper"
)

type Request struct {
	Query string
}

type HttpHandle struct {
	scraper   *scraper.Scraper
	predictor *model.Predictor
	stator    *model.Stator
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func (h *HttpHandle) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there")
}

func (h *HttpHandle) predict(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	var req Request
	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	props := h.scraper.Scrape(req.Query)
	if props.Type != "Lägenhet" {
		http.Error(w, "Only apartments are supported", http.StatusBadRequest)
		return
	}
	
	predictions := h.predictor.PredictPrice(props)
	countyStats := h.stator.ProvideCounty(props)
	
	estimatedPrice := predictions.Result[0]
	roundUpPrice := (math.Round(estimatedPrice*100) / 100) * 1000000
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(Response{
		roundUpPrice,
		Property{
			"Stockholm",
			transformToResponse(countyStats.Size),
			transformToResponse(countyStats.Rent),
			transformToResponse(countyStats.Rooms),
		},
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

func readRecords() [][]string {
	var csvLines [][]string
	
	csvFile, err := os.Open("result.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	r := csv.NewReader(csvFile)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		csvLines = append(csvLines, record)
		
	}
	return csvLines
}

func main() {
	startDeps()
	predictor, err := model.NewPredictor("http://localhost:5000/predict", "model/regions.json")
	if err != nil {
		log.Fatal(err)
	}
	records := readRecords()
	
	httpHandle := &HttpHandle{
		scraper.NewScraper(colly.NewCollector(
			colly.AllowedDomains("hemnet.se", "www.hemnet.se"),
			colly.AllowURLRevisit(),
		)),
		predictor,
		model.NewStator(records),
	}
	
	http.HandleFunc("/", httpHandle.handler)
	http.HandleFunc("/predict", httpHandle.predict)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
