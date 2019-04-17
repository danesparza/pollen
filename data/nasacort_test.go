package data_test

import (
	"testing"

	"github.com/danesparza/pollen/data"
)

func TestNasacort_GetPollenReport_ReturnsValidData(t *testing.T) {
	//	Arrange
	service := data.NasacortService{}
	zipcode := "30019"

	//	Act
	response, err := service.GetPollenReport(zipcode)

	//	Assert
	if err != nil {
		t.Errorf("Error calling GetPollenReport: %v", err)
	}

	t.Logf("Returned object: %+v", response)

}