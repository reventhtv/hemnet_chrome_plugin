package model

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

func TestNewStator(t *testing.T) {
	var request [][]string
	csvLine := "Bergsgatan 8,Bergvik,Södertälje,1585000,1495000,2,53,3606,5100,1944"
	record := strings.Split(csvLine, ",")
	request = append(request, record)
	stator := NewStator(request)
	if len(stator.records) <= 0 {
		t.Fail()
	}
	result := stator.records[0]
	expected := Record{
		kommun: "Södertälje",
		rooms:  2,
		size:   53,
		rent:   3606,
	}
	if result != expected {
		t.Fail()
	}
	t.Log(result)
}

func readRecords() [][]string {
	var csvLines [][]string
	
	csvFile, err := os.Open("test.csv")
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

func TestNewStator_Overall(t *testing.T) {
	request := readRecords()
	stator := NewStator(request)
	if len(stator.records) <= 0 {
		t.Fail()
	}
	result := stator.records[0]
	expected := Record{
		kommun: "Södertälje",
		rooms:  2,
		size:   53,
		rent:   3606,
	}
	if result != expected {
		t.Fail()
	}
	t.Log(result)
}

func TestStator_Provide(t *testing.T) {

}
