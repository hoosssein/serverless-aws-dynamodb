package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"io/ioutil"
	"moghimi/myservice/src/api/deviceHandler"
	"moghimi/myservice/src/dao"
	"os"
)

func main2() {
	lambda.Start(deviceHandler.DeviceHandler{Dao: dao.DynamoDeviceDao{}}.Get)
}

func main() {
	fileName := os.Args[1]
	content, _ := ioutil.ReadFile(fileName)
	request := events.APIGatewayProxyRequest{}
	json.Unmarshal(content, &request)
	post, err := deviceHandler.DeviceHandler{Dao: dao.DynamoDeviceDao{}}.Get(request)
	fmt.Println(post, err)
}
