package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
	"os"
	"time"
)

func pagedUrl(minPrice int, maxPrice int, pageNumber int) string {
	url := "https://www.hemnet.se/salda/bostader?location_ids=17744&item_types%5B%5D=bostadsrat"
	minPriceQuery := fmt.Sprintf("&selling_price_min=%d", minPrice)
	maxPriceQuery := fmt.Sprintf("&selling_price_max=%d", maxPrice)
	pageNumberQuery := fmt.Sprintf("&page=%d", pageNumber)
	return url + minPriceQuery + maxPriceQuery + pageNumberQuery + "&sold_age=6m"
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func scrape(store *Storage) {
	
	// Instantiate default collector
	collector := colly.NewCollector(
		colly.AllowedDomains("hemnet.se", "www.hemnet.se"),
	)
	// TODO: Limit the number of threads started by colly to two
	// when visiting links which domains' matches "*httpbin.*" glob
	_ = collector.Limit(&colly.LimitRule{
		DomainGlob: "*hemnet.se*",
		Delay:      2 * time.Second,
	})
	
	header := []string{"Name",
		"Region",
		"Kommun",
		"SoldPrice",
		"AskPrice",
		"Rooms",
		"Size",
		"MonthlyRent",
		"YearlyRent",
		"YearBuilt",
	}
	
	beginPrice := 0
	endPrice := 1750000
	stopScrapePrice := 12000000 // 12 Million
	
	filename := "result-not-clean.csv"
	file, err := os.Create(filename)
	checkError("Cannot create file", err)
	writer := csv.NewWriter(file)
	if err := writer.Write(header); err != nil {
		log.Println("Failed to write to CSV")
		log.Println(err)
	}
	for endPrice <= stopScrapePrice {
		for pageNumber := 1; pageNumber <= 50; pageNumber++ {
			url := pagedUrl(beginPrice, endPrice, pageNumber)
			log.Println("Visiting URL")
			log.Println(url)
			if stop := visit(collector, writer, url); stop {
				break
			}
		}
		beginPrice = endPrice
		endPrice = beginPrice + 250000
	}
	
	// Wait until threads are finished
	collector.Wait()
	file.Close()
	writer.Flush()
	store.UploadToS3(filename)
}

func main() {
	currentTime := time.Now().Local()
	subPath := currentTime.Format("20060102")
	store := NewStorage("hemnet-predictor", "data/"+subPath+"/")
	scrape(store)
}
