package model

import (
	"log"
	"sathvik.ninja/hemgissa/scraper"
	"sort"
	"strconv"
)

type Record struct {
	kommun string
	rooms  int
	size   int
	rent   int
}

type Stator struct {
	records     []Record
	countyStats Statistics
}

type Summary struct {
	Min    int
	Max    int
	Median int
	Mean   int
	Values []int
}

type Statistics struct {
	Size  Summary
	Rent  Summary
	Rooms Summary
}

func newSummary(values []int, summary chan Summary) {
	mean := make(chan int)
	go func() {
		var total int
		for _, value := range values {
			total = total + value
		}
		mean <- total / len(values)
	}()
	
	sort.Ints(values)
	max := values[len(values)-1]
	min := values[0]
	median := max - min
	summary <- Summary{
		Max:    max,
		Min:    min,
		Mean:   <-mean,
		Median: median,
		Values: values,
	}
}

func newRecord(csvLine []string) (Record, error) {
	entity := csvLine[0]
	
	rooms, err := strconv.Atoi(csvLine[5])
	if err != nil {
		log.Printf("Could not get Rooms for %s", entity)
		return Record{}, err
	}
	
	size, err := strconv.Atoi(csvLine[6])
	if err != nil {
		log.Printf("Could not get size %s", entity)
		return Record{}, err
	}
	
	rentMonthly, err := strconv.Atoi(csvLine[7])
	if err != nil {
		log.Printf("Could not get rentMonthly %s", entity)
		return Record{}, err
	}
	
	return Record{
		kommun: csvLine[2],
		rooms:  rooms,
		size:   size,
		rent:   rentMonthly,
	}, nil
}

func NewStator(rawRecords [][]string) *Stator {
	
	var records []Record
	var collectedSizes []int
	var collectedRooms []int
	var collectedRents []int
	
	for i, v := range rawRecords {
		
		// Skipping header
		if i == 0 {
			continue
		}
		
		re, err := newRecord(v)
		if err != nil {
			continue
		}
		collectedSizes = append(collectedSizes, re.size)
		collectedRooms = append(collectedRooms, re.rooms)
		collectedRents = append(collectedRents, re.rent)
		records = append(records, re)
	}
	overallStats := overallStats(collectedSizes, collectedRooms, collectedRents)
	return &Stator{
		records:     records,
		countyStats: overallStats,
	}
}

func overallStats(collectedSizes []int, collectedRooms []int, collectedRents []int) Statistics {
	sizeSummary := make(chan Summary)
	go newSummary(collectedSizes, sizeSummary)
	
	roomsSummary := make(chan Summary)
	go newSummary(collectedRooms, roomsSummary)
	
	rentsSummary := make(chan Summary)
	go newSummary(collectedRents, rentsSummary)
	
	overallStats := Statistics{
		Size:  <-sizeSummary,
		Rent:  <-rentsSummary,
		Rooms: <-roomsSummary,
	}
	return overallStats
}

func (s *Stator) ProvideCounty(properties scraper.Properties) Statistics {
	return s.countyStats
}

func (s *Stator) ProvideRegion(properties scraper.Properties) {

}
