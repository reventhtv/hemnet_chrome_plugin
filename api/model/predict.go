package model

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sathvik.ninja/hemgissa/scraper"
	"strconv"
)

type Regions struct {
	Region map[string]int
}

type Predictor struct {
	ModelUrl string
	Encoder  map[string]int
}

func NewPredictor(modelUrl string, encoderPath string) (*Predictor, error) {
	data, err := ioutil.ReadFile(encoderPath)
	var encoderMap Regions
	//err := json.NewDecoder(data).Decode(&encoderMap)
	err = json.Unmarshal(data, &encoderMap)
	if err != nil {
		return nil, err
	}
	
	return &Predictor{
		ModelUrl: modelUrl,
		Encoder:  encoderMap.Region,
	}, nil
}

func strToFloatIgnoreErr(prop string) float64 {
	if s, err := strconv.ParseFloat(prop, 64); err == nil {
		return s
	}
	return 0.0
}

func (p *Predictor) transformQuery(props scraper.Properties) []float64 {
	query := make([]float64, 4)
	
	// Populate region
	regionAsCode := p.Encoder[props.Kommun]
	query[0] = float64(regionAsCode)
	// Populate Size
	size := strToFloatIgnoreErr(props.Size)
	query[1] = size
	
	//Rooms
	rooms := strToFloatIgnoreErr(props.Rooms)
	query[2] = rooms
	
	//Rent regularized
	rent := strToFloatIgnoreErr(props.MonthlyRent)
	query[3] = rent / 1000
	
	return query
}

type PredictResponse struct {
	Result []float64
}

func (p *Predictor) PredictPrice(props scraper.Properties) PredictResponse {
	preparedPros := p.transformQuery(props)
	log.Println(preparedPros)
	postBody, _ := json.Marshal(map[string][][]float64{
		"queries": {
			preparedPros,
		},
	})
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(p.ModelUrl, "application/json", responseBody)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var predictions PredictResponse
	err = decoder.Decode(&predictions)
	if err != nil {
		log.Fatalln(err)
	}
	return predictions
}
