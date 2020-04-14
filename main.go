package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/danesparza/pollen/data"
)

var (
	// BuildVersion contains the version information for the app
	BuildVersion = "Unknown"

	// CommitID is the git commitId for the app.  It's filled in as
	// part of the automated build
	CommitID string
)

// Message is a custom struct event type to handle the Lambda input
type Message struct {
	Zipcode string `json:"zipcode"`
}

// HandleRequest handles the AWS lambda request
func HandleRequest(ctx context.Context, msg Message) (data.PollenReport, error) {
	xray.Configure(xray.Config{LogLevel: "trace"})
	ctx, seg := xray.BeginSegment(ctx, "pollen-lambda-handler")

	//	Set the services to call with
	services := []data.PollenService{
		data.NasacortService{},
		data.ZyrtecService{},
		data.PollencomService{},
	}

	//	Call the helper method to get the report:
	response := data.GetPollenReport(ctx, services, msg.Zipcode)

	//	Set the service version information:
	response.Version = fmt.Sprintf("%s.%s", BuildVersion, CommitID)

	//	Close the segment
	seg.Close(nil)

	//	Return our response
	return response, nil
}

func main() {
	//	Immediately forward to Lambda
	lambda.Start(HandleRequest)
}
