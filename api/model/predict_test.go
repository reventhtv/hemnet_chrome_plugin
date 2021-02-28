package model

import (
	"sathvik.ninja/hemgissa/scraper"
	"testing"
)

func TestNewPredictor(t *testing.T) {
	SUT, err := NewPredictor("", "regions.json")
	if err != nil {
		t.Fail()
	}
	//{Kanelgränd 23 Farmarstigen 2  1986 4882}
	// 278.0,35.0,2.0,4882
	testValue := scraper.Properties{
		Name:        "Kanelgränd 23",
		Kommun:      "Farmarstigen",
		Rooms:       "2",
		Size:        "35",
		YearBuilt:   "1986",
		MonthlyRent: "4882",
	}
	query := SUT.transformQuery(testValue)
	
	expected := [4]float64{278.0, 35.0, 2.0, 4.882}
	
	for i := 0; i < 4; i++ {
		if query[i] != expected[i] {
			t.Fail()
		}
		t.Log(query[i])
	}
}

func TestPredictor_PredictPrice(t *testing.T) {
	SUT, err := NewPredictor("http://localhost:5000/predict", "regions.json")
	if err != nil {
		t.Fail()
	}
	
	testValue := scraper.Properties{
		Name:        "Kanelgränd 23",
		Kommun:      "Farmarstigen",
		Rooms:       "2",
		Size:        "65",
		YearBuilt:   "1986",
		MonthlyRent: "4882",
	}
	result := SUT.PredictPrice(testValue)
	if len(result.Result) != 1 {
		t.Fail()
	}
	
	t.Log(result.Result)
}
