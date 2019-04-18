package data

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/aws/aws-xray-sdk-go/xray"
	"golang.org/x/net/context/ctxhttp"
)

// NasacortService is a pollen service for Zyrtec formatted data
type NasacortService struct{}

// NasacortResponse is the native service return format
type NasacortResponse struct {
	Response struct {
		Status        string `json:"status"`
		Location      string `json:"location"`
		Today         string `json:"today"`
		Tomorrow      string `json:"tomorrow"`
		AfterTomorrow string `json:"after_tomorrow"`
		Day4          string `json:"day_4"`
		Source        string `json:"source"`
		City          string `json:"city"`
		State         string `json:"state"`
		Raw           string `json:"raw"`
	} `json:"response"`
}

// GetPollenReport gets the pollen report
func (s NasacortService) GetPollenReport(ctx context.Context, zipcode string) (PollenReport, error) {
	//	Start the service segment
	ctx, seg := xray.BeginSegment(ctx, "nasacort-service")

	//	Our return value
	retval := PollenReport{}

	//	Format the url:
	apiurl := "https://www.nasacort.com/wp-json/pollen/get/"

	resp, err := ctxhttp.PostForm(ctx, xray.Client(nil), apiurl, url.Values{
		"zipcode": {zipcode},
	})

	if err != nil {
		seg.AddError(err)
		apperr := fmt.Errorf("There was a problem calling Nasacort API: %s", err)
		return retval, apperr
	}
	defer resp.Body.Close()

	//	If the HTTP status code indicates an error, report it and get out
	if resp.StatusCode >= 400 {
		apperr := fmt.Errorf("There was an error getting information from Nasacort API: %s", resp.Status)
		return retval, apperr
	}

	//	Decode the return object
	serviceResponse := NasacortResponse{}
	err = json.NewDecoder(resp.Body).Decode(&serviceResponse)
	if err != nil {
		seg.AddError(err)
		apperr := fmt.Errorf("There was a problem decoding the response from Nasacort API: %s", err)
		return retval, apperr
	}

	//	Parse the data items:
	dataitems := []float64{}
	parsedToday, _ := strconv.ParseFloat(serviceResponse.Response.Today, 64)
	dataitems = append(dataitems, parsedToday)

	parsedTomorrow, _ := strconv.ParseFloat(serviceResponse.Response.Tomorrow, 64)
	dataitems = append(dataitems, parsedTomorrow)

	parsedAfterTomorrrow, _ := strconv.ParseFloat(serviceResponse.Response.AfterTomorrow, 64)
	dataitems = append(dataitems, parsedAfterTomorrrow)

	parsedDay4, _ := strconv.ParseFloat(serviceResponse.Response.Day4, 64)
	dataitems = append(dataitems, parsedDay4)

	//	Set the properties in the return object:
	retval = PollenReport{
		ReportingService:  "Nasacort",
		PredominantPollen: serviceResponse.Response.Source,
		Zipcode:           zipcode,
		Location:          fmt.Sprintf("%s, %s", serviceResponse.Response.City, serviceResponse.Response.State),
		StartDate:         time.Now(),
		Data:              dataitems,
	}

	return retval, nil
}
