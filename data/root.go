package data

import (
	"time"
)

// PollenReport represents the report of pollen data
type PollenReport struct {
	Location          string    `json:"location"`           // The location for the report
	Zipcode           string    `json:"zip"`                // The zipcode for the report
	PredominantPollen string    `json:"predominant_pollen"` // The predominant pollen in the report period
	StartDate         time.Time `json:"startdate"`          // The start time for this report
	Data              []float64 `json:"data"`               //	Pollen data indices -- one for today and each future day
	ReportingService  string    `json:"service"`            // The reporting service
}

// PollenService is the interface for all services that can fetch pollen data
type PollenService interface {
	// GetPollenReport gets the pollen report
	GetPollenReport(zipcode string) (PollenReport, error)
}
