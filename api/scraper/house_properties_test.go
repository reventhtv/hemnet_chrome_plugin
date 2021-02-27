package scraper

import "testing"

func TestProperties(t *testing.T) {
	// TODO: Provide self owned file to avoid flaky test
	testURl := "https://www.hemnet.se/bostad/lagenhet-2rum-farmarstigen-tyreso-kommun-kanelgrand-23-17320817"
	result := Scrape(testURl)
	if result.Kommun != "Farmarstigen" {
		t.Fail()
	}
	t.Log(result)
}
