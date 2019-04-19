package data_test

import (
	"context"
	"testing"

	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/danesparza/pollen/data"
)

func TestZyrtec_GetPollenReport_ReturnsValidData(t *testing.T) {
	//	Arrange
	service := data.ZyrtecService{}
	zipcode := "30019"
	ctx := context.Background()
	ctx, seg := xray.BeginSegment(ctx, "unit-test")
	defer seg.Close(nil)

	//	Act
	response, err := service.GetPollenReport(ctx, zipcode)

	//	Assert
	if err != nil {
		t.Errorf("Error calling GetPollenReport: %v", err)
	}

	t.Logf("Returned object: %+v", response)

}
