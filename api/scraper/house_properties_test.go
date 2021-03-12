package scraper

import (
	"github.com/gocolly/colly/v2"
	"os"
	"path/filepath"
	"testing"
)

func TestProperties(t *testing.T) {
	//testURl := "https://www.hemnet.se/bostad/lagenhet-2rum-farmarstigen-tyreso-kommun-kanelgrand-23-17320817"
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	
	c := colly.NewCollector()
	
	scraper := NewScraper(c)
	result := scraper.Scrape(dir)
	
	expected := Properties{
		Name:        "Kanelgränd 23",
		Kommun:      "Farmarstigen",
		Rooms:       "2",
		Size:        "62",
		YearBuilt:   "1986",
		MonthlyRent: "4882",
	}
	if result != expected {
		t.Fail()
	}
	t.Log(result)
}
