package data

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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
func (s ClaritinService) GetPollenReport(zipcode string) (PollenReport, error) {
	//	Our return value
	retval := PollenReport{}

	//	Format the url:
	apiurl := fmt.Sprintf("https://www.claritin.com/claritinapi/globalheader/getallergyforecastdata?zipcode=%s", zipcode)

	//	Create the client
	client := &http.Client{}

	//	Create our request:
	req, err := http.NewRequest("GET", apiurl, nil)
	if err != nil {
		apperr := fmt.Errorf("There was a problem creating the Claritin API request: %s", err)
		return retval, apperr
	}

	//	Execute our request:
	resp, err := client.Do(req)

	if err != nil {
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
		apperr := fmt.Errorf("There was a problem decoding the response from Claritin API: %s", err)
		return retval, apperr
	}

	//	Parse the date/time:
	layoutBaseClartin := "1/02/2006 15:04:05 PM"
	parsedStartDate, err := time.Parse(layoutBaseClartin, serviceResponse.PollenForecast.Timestamp)
	if err != nil {
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

	return retval, nil
}
