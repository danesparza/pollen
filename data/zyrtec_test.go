package data_test

import (
	"testing"

	"github.com/danesparza/pollen/data"
)

func TestZyrtec_GetPollenReport_ReturnsValidData(t *testing.T) {
	//	Arrange
	service := data.ZyrtecService{}
	zipcode := "30019"

	//	Act
	_, err := service.GetPollenReport(zipcode)

	//	Assert
	if err != nil {
		t.Errorf("Error calling GetPollenReport: %v", err)
	}

}