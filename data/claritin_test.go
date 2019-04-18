package data_test

import (
	"context"
	"testing"

	"github.com/danesparza/pollen/data"
)

func TestClaritin_GetPollenReport_ReturnsValidData(t *testing.T) {
	//	Arrange
	service := data.ClaritinService{}
	zipcode := "30019"

	//	Act
	response, err := service.GetPollenReport(context.Background(), zipcode)

	//	Assert
	if err != nil {
		t.Errorf("Error calling GetPollenReport: %v", err)
	}

	t.Logf("Returned object: %+v", response)

}
