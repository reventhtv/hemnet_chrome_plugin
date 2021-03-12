package scraper

import (
	"github.com/gocolly/colly/v2"
	"log"
	"strings"
)

type Properties struct {
	Name        string
	Type        string
	Kommun      string
	Rooms       string
	Size        string
	YearBuilt   string
	MonthlyRent string
}

type Scraper struct {
	colly *colly.Collector
}

func NewScraper(colly *colly.Collector) *Scraper {
	return &Scraper{
		colly,
	}
	
}

func (s *Scraper) Scrape(url string) Properties {
	// Instantiate default collector
	scraper := s.colly
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
			value = strings.Replace(value, " m²", "", 1)
			properties.Size = strings.Split(value, ",")[0]
		}
	})
	
	scraper.OnHTML("div[class=property-attributes-table]", func(e *colly.HTMLElement) {
		e.ForEach("dl[class=property-attributes-table__area]", func(_ int, element *colly.HTMLElement) {
			e.ForEach("div[class=property-attributes-table__row]", func(_ int, internal *colly.HTMLElement) {
				label := internal.ChildText(".property-attributes-table__label")
				value := internal.ChildText(".property-attributes-table__value")
				
				switch label {
				case "Bostadstyp":
					properties.Type = value
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
	log.Println(properties)
	return properties
}
