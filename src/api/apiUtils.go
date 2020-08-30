package api

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"moghimi/myservice/src/utils"
)

func SendError(request *events.APIGatewayProxyRequest, err error) (events.APIGatewayProxyResponse, error) {
	code := findStatus(err)
	message := findMessage(err)
	logError(request, err)
	return events.APIGatewayProxyResponse{Body: message, StatusCode: code}, nil
}

func logError(request *events.APIGatewayProxyRequest, err error) {
	fmt.Printf("error happend for Request: %v, Error: %v\n", request, err)
}

func findMessage(err error) string {
	httpError, ok := err.(utils.HttpError)
	if ok {
		return httpError.UserError()
	}
	return "INTERNAL SERVER ERROR"
}

func findStatus(err error) int {
	httpError, ok := err.(utils.HttpError)
	if ok {
		return int(httpError.Code)
	}
	return 500
}

func SendJson(request *events.APIGatewayProxyRequest, object interface{}, status int) (events.APIGatewayProxyResponse, error) {
	marshal, err := json.Marshal(object)
	if err != nil {
		return SendError(request, err)
	}

	return events.APIGatewayProxyResponse{
		Body:       string(marshal),
		StatusCode: status,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}
