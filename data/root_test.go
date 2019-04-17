package data_test

import (
	"testing"

	"github.com/danesparza/pollen/data"
)

func TestMultipleServices_GetPollenData_ReturnsValidData(t *testing.T) {
	//	Arrange
	services := []data.PollenService{
		data.ClaritinService{},
		data.NasacortService{},
		data.ZyrtecService{},
	}
	zipcode := "30019"

	//	Act
	response := data.GetPollenReport(services, zipcode)

	//	Assert
	t.Logf("Returned object: %+v", response)

}
