package data_test

import (
	"context"
	"testing"

	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/danesparza/pollen/data"
)

func TestMultipleServices_GetPollenData_ReturnsValidData(t *testing.T) {
	//	Arrange
	services := []data.PollenService{
		data.ClaritinService{},
		data.NasacortService{},
		data.ZyrtecService{},
		data.PollencomService{},
	}
	zipcode := "30019"
	ctx := context.Background()
	ctx, seg := xray.BeginSegment(ctx, "unit-test")
	defer seg.Close(nil)

	//	Act
	response := data.GetPollenReport(ctx, services, zipcode)

	//	Assert
	t.Logf("Returned object: %+v", response)

}
