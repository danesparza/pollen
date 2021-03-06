package data

import (
	"context"
	"time"

	"github.com/aws/aws-xray-sdk-go/xray"
)

// PollenReport represents the report of pollen data
type PollenReport struct {
	Location          string    `json:"location"`           // The location for the report
	Zipcode           string    `json:"zip"`                // The zipcode for the report
	PredominantPollen string    `json:"predominant_pollen"` // The predominant pollen in the report period
	StartDate         time.Time `json:"startdate"`          // The start time for this report
	Data              []float64 `json:"data"`               //	Pollen data indices -- one for today and each future day
	ReportingService  string    `json:"service"`            // The reporting service
	Version           string    `json:"version"`            // Service version information
}

// PollenService is the interface for all services that can fetch pollen data
type PollenService interface {
	// GetPollenReport gets the pollen report
	GetPollenReport(ctx context.Context, zipcode string) (PollenReport, error)
}

// GetPollenReport calls all services in parallel and returns the first result
func GetPollenReport(ctx context.Context, services []PollenService, zipcode string) PollenReport {

	ch := make(chan PollenReport, 1)

	//	Start the service segment
	ctx, seg := xray.BeginSubsegment(ctx, "pollen-report")
	defer seg.Close(nil)

	//	For each passed service ...
	for _, service := range services {

		//	Launch a goroutine for each service...
		go func(c context.Context, s PollenService, zip string) {

			//	Get its pollen report ...
			result, err := s.GetPollenReport(c, zip)

			//	As long as we don't have an error, return what we found on the result channel
			if err == nil {
				
				//	Make sure we also have more than one datapoint!
				if len(result.Data) > 1 {
					select {
					case ch <- result:
					default:
					}
				}				
			}
		}(ctx, service, zipcode)

	}

	//	Return the first result passed on the channel
	return <-ch
}
