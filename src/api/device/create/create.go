package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"io/ioutil"
	"moghimi/myservice/src/dao"
	"moghimi/myservice/src/model"
	"moghimi/myservice/src/utils"
	"os"
)

func Post(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println(request.Body)
	device := &model.Device{}
	err := json.Unmarshal([]byte(request.Body), &device)
	if err != nil {
		return sendError(utils.HttpError{Err: err, Code: 400})
	}

	err = device.Validate()
	if err != nil {
		return sendError(utils.HttpError{Err: err, Code: 400})
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
	code := findStatus(err)
	return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: code}, nil
}

func findStatus(err error) int {
	httpError, ok := err.(utils.HttpError)
	if ok {
		return int(httpError.Code)
	}
	return 500
}

func main2() {
	lambda.Start(Post)
}

func main() {
	fileName := os.Args[1]
	content, _ := ioutil.ReadFile(fileName)
	request := events.APIGatewayProxyRequest{}
	json.Unmarshal(content, &request)
	post, err := Post(request)
	fmt.Println(post, err)
}
