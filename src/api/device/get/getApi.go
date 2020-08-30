package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"io/ioutil"
	"moghimi/myservice/src/api"
	"moghimi/myservice/src/dao"
	"os"
)

const IdPrefix = "/devices/"

func Get(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := request.PathParameters["id"]
	id = IdPrefix + id

	device, err := dao.GetDevice(id)
	if err != nil {
		return api.SendError(&request, err)
	}

	return api.SendJson(&request, &device, 200)
}

func main2() {
	lambda.Start(Get)
}

func main() {
	fileName := os.Args[1]
	content, _ := ioutil.ReadFile(fileName)
	request := events.APIGatewayProxyRequest{}
	json.Unmarshal(content, &request)
	post, err := Get(request)
	fmt.Println(post, err)
}
