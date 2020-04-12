package data

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/context/ctxhttp"

	"github.com/aws/aws-xray-sdk-go/xray"
)

// PollencomService is a pollen service for Pollen.com formatted data
type PollencomService struct{}

// PollencomResponse is the native service return format
type PollencomResponse struct {
	Type         string `json:"Type"`
	ForecastDate string `json:"ForecastDate"`
	Location     struct {
		ZIP     string `json:"ZIP"`
		City    string `json:"City"`
		State   string `json:"State"`
		Periods []struct {
			Period string  `json:"Period"`
			Index  float64 `json:"Index"`
		} `json:"periods"`
		DisplayLocation string `json:"DisplayLocation"`
	} `json:"Location"`
}

// GetPollenReport gets the pollen report
func (s PollencomService) GetPollenReport(ctx context.Context, zipcode string) (PollenReport, error) {
	//	Start the service segment
	ctx, seg := xray.BeginSubsegment(ctx, "pollencom-service")

	//	Our return value
	retval := PollenReport{}

	//	Format the url:
	apiurl := fmt.Sprintf("https://www.pollen.com/api/forecast/extended/pollen/%s", zipcode)

	req, _ := http.NewRequest("GET", apiurl, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.146 Safari/537.36")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Referer", apiurl)

	resp, err := ctxhttp.Do(ctx, xray.Client(nil), req)

	if err != nil {
		seg.AddError(err)
		apperr := fmt.Errorf("There was a problem calling Pollen.com API: %s", err)
		return retval, apperr
	}
	defer resp.Body.Close()

	//	If the HTTP status code indicates an error, report it and get out
	if resp.StatusCode >= 400 {
		apperr := fmt.Errorf("There was an error getting information from Pollen.com API: %s", resp.Status)
		return retval, apperr
	}

	//	Decode the return object
	serviceResponse := PollencomResponse{}
	err = json.NewDecoder(resp.Body).Decode(&serviceResponse)
	if err != nil {
		seg.AddError(err)
		apperr := fmt.Errorf("There was a problem decoding the response from Pollen.com API: %s", err)
		return retval, apperr
	}

	//	Parse the data items:
	dataitems := []float64{}

	parsedToday := serviceResponse.Location.Periods[0].Index
	dataitems = append(dataitems, parsedToday)

	parsedTomorrow := serviceResponse.Location.Periods[1].Index
	dataitems = append(dataitems, parsedTomorrow)

	parsedAfterTomorrrow := serviceResponse.Location.Periods[2].Index
	dataitems = append(dataitems, parsedAfterTomorrrow)

	parsedDay4 := serviceResponse.Location.Periods[3].Index
	dataitems = append(dataitems, parsedDay4)

	//	Set the properties in the return object:
	retval = PollenReport{
		ReportingService:  "Pollen.com",
		PredominantPollen: "Not in this service call -- need to look at current forecast",
		Zipcode:           zipcode,
		Location:          fmt.Sprintf("%s, %s", serviceResponse.Location.City, serviceResponse.Location.State),
		StartDate:         time.Now(),
		Data:              dataitems,
	}

	xray.AddMetadata(ctx, "PollencomResult", retval)

	// Close the segment
	seg.Close(nil)

	return retval, nil
}
