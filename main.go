package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/danesparza/pollen/data"
)

// Message is a custom struct event type to handle the Lambda input
type Message struct {
	Zipcode string `json:"zipcode"`
}

// HandleRequest handles the AWS lambda request
func HandleRequest(ctx context.Context, msg Message) (data.PollenReport, error) {

	//	Our return value:
	retval := data.PollenReport{}

	//	Return our return response
	return retval, nil
}

func main() {
	//	Immediately forward to Lambda
	lambda.Start(HandleRequest)
}
