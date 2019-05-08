# pollen [![CircleCI](https://circleci.com/gh/danesparza/pollen.svg?style=shield)](https://circleci.com/gh/danesparza/pollen)
AWS Lambda based API service to get pollen data forecast for a given zipcode.  It's designed to attempt to get pollen data from 3 sources in parallel and return the data from the service that responds first.  Instrumented with [AWS x-ray](https://aws.amazon.com/xray/).

## Quick start
To use this labmda handler, grab the `pollen_handler.zip` from the [latest release](https://github.com/danesparza/pollen/releases/latest) and [create a lambda function](https://docs.aws.amazon.com/lambda/latest/dg/lambda-app.html#lambda-app-upload-deployment-pkg) in [your AWS account](https://console.aws.amazon.com/lambda/home).  

When creating the function on AWS ... in the *Function code* section be sure to upload the `handler.zip`, select the 'Go' runtime, and select 'pollen' as the handler name:
![Screenshot of lambda creation in AWS console](lambda_setup.png?raw=true)

## Testing the service
To test the service, create a simple test event for your Lambda function.  The event should include the zipcode you want to get pollen data for:
```json
{
  "zipcode": "30019"
}
```

You should get a nice JSON response that looks like this:
```json
{
  "location": "DACULA, GA",
  "zip": "30019",
  "predominant_pollen": "Oak, Birch and Sycamore.",
  "startdate": "2019-04-18T13:00:33.548528198Z",
  "data": [
    10.2,
    1,
    7.9,
    10
  ],
  "service": "Nasacort",
  "version": "1.0.4.f3092bea655df439b7b3eaa3cfe9b628dec03cff"
}
```

## What does the data mean?
Parameter          | Description
----------         | -----------
location           | The detected city/state location for the report
zip                | The zipcode that was passed to the Lambda function
predominant_pollen | The predominant pollen currently detected in the area
startdate          | The date/time for the report
data               | An array of floats.  This indicates the pollen indices by day, starting with today.  In the case of the example above, today's pollen index is 10.2, tomorrow's pollen index is 1, the next day's index is 7.9, etc.  
service            | The reporting service
version            | The version of the pollen Lambda service being used

## How can use it outside of AWS?
Simple!  Just use [AWS API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-lambda-integrations.html) to setup a REST API that calls your new Lambda function.

## AWS X-ray?
Yep -- the service is instrumented with [AWS X-ray](https://aws.amazon.com/xray/), so you can get an idea of runtime performance.  Just navigate to X-Ray in your console to check it out.
