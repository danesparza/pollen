package data

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// ZyrtecService is a pollen service for Zyrtec formatted data
type ZyrtecService struct{}

// ZyrtecResponse is the native service return format
type ZyrtecResponse struct {
	PredominantPollen string `json:"predominantPollen"`
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

	log.Printf("Response: %+v", serviceResponse)

	return retval, nil
}
