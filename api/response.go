package main

import "sathvik.ninja/hemgissa/model"

type Summary struct {
	Min    int
	Max    int
	Median int
	Mean   int
	Values []int
}

type Property struct {
	Name  string
	Size  Summary
	Rent  Summary
	Rooms Summary
}

type Response struct {
	EstimatedPrice float64 `json:"estimatedPrice"`
	CountySummary  Property
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
