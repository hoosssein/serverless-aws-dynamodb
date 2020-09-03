package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"moghimi/myservice/src/api/deviceHandler"
)

func main() {
	lambda.Start(deviceHandler.Handler.Get)
}
