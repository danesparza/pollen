# pollen [![CircleCI](https://circleci.com/gh/danesparza/pollen.svg?style=shield)](https://circleci.com/gh/danesparza/pollen)
AWS Lambda based service to get pollen data forecast for a given zipcode

## Quick start
To use this labmda handler, get the [latest release](https://github.com/danesparza/pollen/releases/latest) and [create a lambda function](https://docs.aws.amazon.com/lambda/latest/dg/lambda-app.html#lambda-app-upload-deployment-pkg) in [your AWS account](https://console.aws.amazon.com/lambda/home).  

In the *Function code* section be sure to upload the handler, select the 'Go' runtime, and select 'pollen' as the handler name:
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
