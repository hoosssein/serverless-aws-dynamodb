package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"serverless-aws-dynamodb/src/api/deviceHandler"
)

func main() {
	lambda.Start(deviceHandler.Handler.Post)
}
