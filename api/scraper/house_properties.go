package scraper

import (
	"github.com/gocolly/colly/v2"
	"log"
	"strings"
)

type Properties struct {
	Name        string
	Kommun      string
	Rooms       string
	Size        string
	YearBuilt   string
	MonthlyRent string
}

func Scrape(url string) Properties {
	// Instantiate default collector
	scraper := colly.NewCollector(
		colly.AllowedDomains("hemnet.se", "www.hemnet.se"),
	)
	var properties Properties
	scraper.OnHTML("div[class=property-address]", func(e *colly.HTMLElement) {
		properties.Name = e.ChildText(".qa-property-heading")
		fullKommun := e.ChildText(".property-address__area")
		finalValue := strings.Split(fullKommun, ",")[0]
		properties.Kommun = finalValue
	})
	
	scraper.OnHTML(".qa-living-area-attribute", func(element *colly.HTMLElement) {
		label := element.ChildText(".property-attributes-table__label")
		value := element.ChildText(".property-attributes-table__value")
		if label == "Boarea" {
			properties.Size = strings.Replace(value, " m²", "", 1)
		}
	})
	
	scraper.OnHTML("div[class=property-attributes-table]", func(e *colly.HTMLElement) {
		e.ForEach("dl[class=property-attributes-table__area]", func(_ int, element *colly.HTMLElement) {
			e.ForEach("div[class=property-attributes-table__row]", func(_ int, internal *colly.HTMLElement) {
				label := internal.ChildText(".property-attributes-table__label")
				value := internal.ChildText(".property-attributes-table__value")
				
				switch label {
				case "Antal rum":
					properties.Rooms = strings.Replace(value, " rum", "", 1)
				case "Byggår":
					properties.YearBuilt = value
				case "Avgift":
					clearKr := strings.Replace(value, "kr", "", 1)
					clearKr = strings.Replace(clearKr, "mån", "", 1)
					clearKr = strings.Replace(clearKr, "/", "", 1)
					finalValue := strings.Replace(clearKr, " ", "", 2)
					properties.MonthlyRent = finalValue
				}
				
			})
		})
	})
	
	if err := scraper.Visit(url); err != nil {
		log.Println(err)
	}
	return properties
}
