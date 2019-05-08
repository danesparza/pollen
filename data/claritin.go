package data

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-xray-sdk-go/xray"
	"golang.org/x/net/context/ctxhttp"
)

// ClaritinService is a pollen service for Zyrtec formatted data
type ClaritinService struct{}

// ClaritinResponse is the native service return format
type ClaritinResponse struct {
	PollenForecast struct {
		Zip       string    `json:"zip"`
		City      string    `json:"city"`
		State     string    `json:"state"`
		Forecast  []float64 `json:"forecast"`
		Pp        string    `json:"pp"`
		Timestamp string    `json:"timestamp"`
	} `json:"pollenForecast"`
	Result bool `json:"result"`
}

// GetPollenReport gets the pollen report
func (s ClaritinService) GetPollenReport(ctx context.Context, zipcode string) (PollenReport, error) {
	//	Start the service segment
	ctx, seg := xray.BeginSubsegment(ctx, "claritin-service")

	//	Our return value
	retval := PollenReport{}

	//	Format the url:
	apiurl := fmt.Sprintf("https://www.claritin.com/claritinapi/globalheader/getallergyforecastdata?zipcode=%s", zipcode)

	resp, err := ctxhttp.Get(ctx, xray.Client(nil), apiurl)
	if err != nil {
		seg.AddError(err)
		apperr := fmt.Errorf("There was a problem calling Claritin API: %s", err)
		return retval, apperr
	}

	defer resp.Body.Close()

	//	If the HTTP status code indicates an error, report it and get out
	if resp.StatusCode >= 400 {
		apperr := fmt.Errorf("There was an error getting information from Claritin API: %s", resp.Status)
		return retval, apperr
	}

	//	Decode the return object
	serviceResponse := ClaritinResponse{}
	err = json.NewDecoder(resp.Body).Decode(&serviceResponse)
	if err != nil {
		seg.AddError(err)
		apperr := fmt.Errorf("There was a problem decoding the response from Claritin API: %s", err)
		return retval, apperr
	}

	//	Parse the date/time:
	layoutBaseClartin := "1/2/2006 15:04:05 PM"
	parsedStartDate, err := time.Parse(layoutBaseClartin, serviceResponse.PollenForecast.Timestamp)
	if err != nil {
		seg.AddError(err)
		apperr := fmt.Errorf("There was a problem parsing the date in the response from Claritin API: %s", err)
		return retval, apperr
	}

	//	Parse the data items:
	dataitems := serviceResponse.PollenForecast.Forecast

	//	Set the properties in the return object:
	retval = PollenReport{
		ReportingService:  "Claritin",
		PredominantPollen: serviceResponse.PollenForecast.Pp,
		Zipcode:           serviceResponse.PollenForecast.Zip,
		Location:          fmt.Sprintf("%s, %s", serviceResponse.PollenForecast.City, serviceResponse.PollenForecast.State),
		StartDate:         parsedStartDate,
		Data:              dataitems,
	}

	xray.AddMetadata(ctx, "ClaritinResult", retval)

	// Close the segment
	seg.Close(nil)

	return retval, nil
}
