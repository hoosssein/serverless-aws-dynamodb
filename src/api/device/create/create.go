package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"moghimi/myservice/src/dao"
	"moghimi/myservice/src/model"
)

func Post(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println(request.Body)
	device := model.Device{}
	err := json.Unmarshal([]byte(request.Body), &device)
	if err != nil {
		return sendError(err)
	}

	err = device.Validate()
	if err != nil {
		return sendError(err)
	}

	device, err = dao.SaveDevice(device)
	if err != nil {
		return sendError(err)
	}

	marshal, err := json.Marshal(device)
	if err != nil {
		return sendError(err)
	}

	return events.APIGatewayProxyResponse{
		Body:       string(marshal),
		StatusCode: 201,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func sendError(err error) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 400}, nil
}

func main() {
	lambda.Start(Post)
}
