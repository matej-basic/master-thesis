package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest() (string, error){
	return "Simple Golang Lambda benchmark", nil
}
func main() {
	lambda.Start(HandleRequest)
}

