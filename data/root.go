package data

import (
	"time"
)

// PollenReport represents the report of pollen data
type PollenReport struct {
	State             string    `json:"state"`              // The state for the report period
	City              string    `json:"city"`               // The city for the report period
	Zipcode           string    `json:"zip"`                // The zipcode for the report period
	PredominantPollen string    `json:"predominant_pollen"` // The predominant pollen in the report period
	StartDate         time.Time `json:"startdate"`          // The start time for this report
	Data              []float64 `json:"data"`               //	Pollen data indices -- one for today and each future day
}

// PollenService is the interface for all services that can fetch pollen data
type PollenService interface {
	// GetPollenReport gets the pollen report
	GetPollenReport(zipcode string) (PollenReport, error)
}
