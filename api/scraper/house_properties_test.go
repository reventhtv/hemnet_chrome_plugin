package scraper

import "testing"

func TestProperties(t *testing.T) {
	// TODO: Provide self owned file to avoid flaky test
	testURl := "https://www.hemnet.se/bostad/lagenhet-2rum-farmarstigen-tyreso-kommun-kanelgrand-23-17320817"
	result := Scrape(testURl)
	
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
