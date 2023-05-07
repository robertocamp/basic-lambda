package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

func hello() (string, error) {
	return "Hello λ!", nil
}

func main() {
	// make the handler available for Remote Procedure Calls by AWS Lambda
	lambda.Start(hello)
}