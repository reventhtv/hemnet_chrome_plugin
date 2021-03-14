package main

import "sathvik.ninja/hemgissa/model"

type Summary struct {
	Min    int   `json:"min"`
	Max    int   `json:"max"`
	Median int   `json:"median"`
	Mean   int   `json:"mean"`
	Values []int `json:"values"`
}

type Property struct {
	Name  string  `json:"name"`
	Size  Summary `json:"size"`
	Rent  Summary `json:"rent"`
	Rooms Summary `json:"rooms"`
}

type Response struct {
	EstimatedPrice float64  `json:"estimatedPrice"`
	CountySummary  Property `json:"countySummary"`
}

func transformToResponse(statsInternal model.Summary) Summary {
	return Summary{
		Min:    statsInternal.Min,
		Max:    statsInternal.Max,
		Mean:   statsInternal.Mean,
		Median: statsInternal.Median,
		Values: statsInternal.Values,
	}
}
