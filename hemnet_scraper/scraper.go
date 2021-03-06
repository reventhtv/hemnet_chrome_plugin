package main

import (
	"encoding/csv"
	"github.com/gocolly/colly/v2"
	"log"
	"strconv"
	"strings"
)

type SoldProperties struct {
	Name        string
	Region      string
	Kommun      string
	SoldPrice   int
	AskPrice    int
	Rooms       int
	Size        int
	MonthlyRent int
	YearlyRent  int
	YearBuilt   int
}

func removeCurrencyAndWs(s string) string {
	s = strings.Replace(s, "kr", "", 1)
	return removeWhiteSpaces(s)
}

func removeWhiteSpaces(s string) string {
	// a weird white space. Could not figure out what
	s = strings.ReplaceAll(s, " ", "")
	// A normal space
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "\n", "")
	return s
}

func deepVisit(internalScraper *colly.Collector, url string) SoldProperties {
	var properties SoldProperties
	
	internalScraper.OnHTML("h1", func(element *colly.HTMLElement) {
		name := strings.Trim(element.Text, " ")
		name = strings.Replace(name, "Slutpris", "", 1)
		name = strings.ReplaceAll(name, ",", "")
		name = strings.Trim(name, "\n ")
		properties.Name = name
	})
	
	internalScraper.OnHTML(".sold-property__metadata", func(element *colly.HTMLElement) {
		address := element.Text
		location := strings.Split(address, "-")[1]
		loc := strings.Split(location, ",")
		if len(loc) == 2 {
			properties.Region = removeWhiteSpaces(loc[0])
			
			kommun := strings.Replace(loc[1], "kommun", "", 1)
			properties.Kommun = removeWhiteSpaces(kommun)
		}
		if len(loc) == 1 {
			kommun := strings.Replace(loc[0], "kommun", "", 1)
			properties.Kommun = removeWhiteSpaces(kommun)
		}
	})
	internalScraper.OnHTML(".sold-property__price-value", func(element *colly.HTMLElement) {
		price := element.Text
		price = removeCurrencyAndWs(price)
		priceIn, err := strconv.Atoi(price)
		if err != nil {
			// handle error
			log.Println(err)
		}
		properties.SoldPrice = priceIn
	})
	
	internalScraper.OnHTML(".sold-property__details", func(element *colly.HTMLElement) {
		element.ForEach(".sold-property__price-stats", func(_ int, internal *colly.HTMLElement) {
			labels := internal.ChildTexts(".sold-property__attribute")
			values := internal.ChildTexts(".sold-property__attribute-value")
			if labels[1] == "Begärt pris" {
				value := removeCurrencyAndWs(values[1])
				priceIn, err := strconv.Atoi(value)
				if err != nil {
					// handle error
					log.Println(err)
				}
				properties.AskPrice = priceIn
			}
			
		})
		element.ForEach(".sold-property__attributes", func(_ int, internal *colly.HTMLElement) {
			labels := internal.ChildTexts(".sold-property__attribute")
			values := internal.ChildTexts(".sold-property__attribute-value")
			
			props := make(map[string]string)
			
			for i := 0; i < len(labels); i++ {
				props[labels[i]] = values[i]
			}
			v, found := props["Antal rum"]
			if found {
				rum := strings.Replace(v, "rum", "", 1)
				rooms, _ := strconv.Atoi(removeWhiteSpaces(rum))
				properties.Rooms = rooms
			}
			v, found = props["Boarea"]
			if found {
				area := strings.Replace(v, "m²", "", 1)
				if strings.Contains(area, ",") {
					area = strings.Split(area, ",")[0]
				}
				size, _ := strconv.Atoi(removeWhiteSpaces(area))
				properties.Size = size
			}
			v, found = props["Avgift/månad"]
			if found {
				rentMon := strings.Replace(v, "mån", "", 1)
				rentMon = strings.Replace(rentMon, "/", "", 1)
				rentInInt, _ := strconv.Atoi(removeCurrencyAndWs(rentMon))
				properties.MonthlyRent = rentInInt
			}
			
			v, found = props["Driftskostnad"]
			if found {
				rentMon := strings.Replace(v, "år", "", 1)
				rentMon = strings.Replace(rentMon, "/", "", 1)
				rentInInt, _ := strconv.Atoi(removeCurrencyAndWs(rentMon))
				properties.YearlyRent = rentInInt
			}
			
			v, found = props["Byggår"]
			if found {
				year, _ := strconv.Atoi(removeWhiteSpaces(v))
				properties.YearBuilt = year
			}
		})
	})
	
	if err := internalScraper.Visit(url); err != nil {
		log.Println(err)
	}
	log.Println(properties)
	return properties
}

func visit(collector *colly.Collector, writer *csv.Writer, url string) bool {
	stop := false
	collector.OnError(func(r *colly.Response, err error) {
		if r.StatusCode == 404 {
			stop = true
			log.Println("No hits stopping the page")
			return
		}
	})
	
	stringConvIntSafe := func(i int) string {
		if i == 0 {
			return "N/A"
		}
		return strconv.Itoa(i)
	}
	
	propertiesToString := func(p SoldProperties) ([]string, bool) {
		
		if p.Kommun == "" || p.SoldPrice == 0 {
			return nil, false
		}
		
		name := p.Name
		if name == "" {
			name = "N/A"
		}
		
		kommun := p.Kommun
		if kommun == "" {
			kommun = "N/A"
		}
		
		records := []string{
			name,
			p.Region,
			p.Kommun,
			strconv.Itoa(p.SoldPrice),
			stringConvIntSafe(p.AskPrice),
			stringConvIntSafe(p.Rooms),
			stringConvIntSafe(p.Size),
			stringConvIntSafe(p.MonthlyRent),
			stringConvIntSafe(p.YearlyRent),
			stringConvIntSafe(p.YearBuilt),
		}
		return records, true
	}
	
	collector.OnHTML(".sold-results", func(e *colly.HTMLElement) {
		collectedRecordsCount := 0
		e.ForEach(".sold-property-listing", func(i int, element *colly.HTMLElement) {
			link := element.Attr("href")
			if strings.Index(link, "/salda") > -1 {
				deepCollector := collector.Clone()
				apartment := deepVisit(deepCollector, link)
				record, shouldWrite := propertiesToString(apartment)
				if shouldWrite {
					if err := writer.Write(record); err != nil {
						log.Println("Failed to write to CSV")
						log.Println(err)
					} else {
						collectedRecordsCount = collectedRecordsCount + 1
					}
				}
			}
		})
		log.Println("collectedRecordsCount %i", collectedRecordsCount)
	})
	if err := collector.Visit(url); err != nil {
		log.Println(err)
	}
	return stop
}
