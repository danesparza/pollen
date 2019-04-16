package data

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// ZyrtecService is a pollen service for Zyrtec formatted data
type ZyrtecService struct{}

// ZyrtecResponse is the native service return format
type ZyrtecResponse struct {
	PredominantPollen string `json:"predominantPollen"`
	Location          string `json:"location"`
	Zipcode           string `json:"zipcode"`
	Date              string `json:"date"`
	AllergyScore      string `json:"allergyScore"`
	Forecast          struct {
		Extended []struct {
			PollenScore float64 `json:"pollenScore,omitempty"`
		} `json:"extended"`
	} `json:"forecast"`
}

// GetPollenReport gets the pollen report
func (s ZyrtecService) GetPollenReport(zipcode string) (PollenReport, error) {

	//	Our return value
	retval := PollenReport{}

	//	Format the url:
	url := fmt.Sprintf("https://api.allergycastapp.com/allergies/dashboard/%s", zipcode)

	//	Create the client
	client := &http.Client{}

	//	Create our request:
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	//	Execute our request:
	resp, err := client.Do(req)

	if err != nil {
		apperr := fmt.Errorf("There was a problem calling Zyrtec API: %s", err)
		return retval, apperr
	}
	defer resp.Body.Close()

	//	If the HTTP status code indicates an error, report it and get out
	if resp.StatusCode >= 400 {
		apperr := fmt.Errorf("There was an error getting information from Zyrtec API: %s", resp.Status)
		return retval, apperr
	}

	//	Decode the return object
	serviceResponse := ZyrtecResponse{}
	err = json.NewDecoder(resp.Body).Decode(&serviceResponse)
	if err != nil {
		apperr := fmt.Errorf("There was a problem decoding the response from Zyrtec API: %s", err)
		return retval, apperr
	}

	//	Parse the date/time:
	layoutBaseZyrtec := "2006-01-02"
	parsedStartDate, err := time.Parse(layoutBaseZyrtec, serviceResponse.Date)
	if err != nil {
		apperr := fmt.Errorf("There was a problem parsing the date in the response from Zyrtec API: %s", err)
		return retval, apperr
	}

	//	Parse the data items:
	dataitems := []float64{}
	for _, item := range serviceResponse.Forecast.Extended {
		if item.PollenScore != 0 {
			dataitems = append(dataitems, item.PollenScore)
		}
	}

	//	Set the properties in the return object:
	retval = PollenReport{
		PredominantPollen: serviceResponse.PredominantPollen,
		Zipcode:           serviceResponse.Zipcode,
		City:              serviceResponse.Location,
		StartDate:         parsedStartDate,
		Data:              dataitems,
	}

	return retval, nil
}
