package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"io/ioutil"
	"moghimi/myservice/src/api"
	"moghimi/myservice/src/dao"
	"moghimi/myservice/src/model"
	"moghimi/myservice/src/utils"
	"os"
)

func Post(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	device := &model.Device{}
	err := json.Unmarshal([]byte(request.Body), &device)
	if err != nil {
		return api.SendError(&request, utils.HttpError{OriginalError: err, Code: 400, Message: err.Error()})
	}

	err = device.Validate()
	if err != nil {
		return api.SendError(&request, utils.HttpError{OriginalError: err, Code: 400, Message: err.Error()})
	}

	device, err = dao.SaveDevice(device)
	if err != nil {
		return api.SendError(&request, err)
	}

	return api.SendJson(&request, &device, 201)
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
