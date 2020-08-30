package dao

/*
import "moghimi/myservice/src/model"

func SaveDevice(device *model.Device) (*model.Device, error) {
	return device,nil
}*/

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"moghimi/myservice/src/model"
	"moghimi/myservice/src/utils"
	"os"
)

func SaveDevice(device *model.Device) (*model.Device, error) {
	sess := CreateSession()
	svc := dynamodb.New(sess)

	dynamoAttribute, err := dynamodbattribute.MarshalMap(device)
	if err != nil {
		return nil, err
	}
	fmt.Println("dynamoAttribute: ", dynamoAttribute)
	input := &dynamodb.PutItemInput{
		Item:      dynamoAttribute,
		TableName: tableName(),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		return nil, err
	}

	return device, nil
}

func GetDevice(id string) (*model.Device, error) {
	sess := CreateSession()
	svc := dynamodb.New(sess)

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: tableName(),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})
	if err != nil {
		return handelError(err)
	}

	if result.Item == nil {
		return nil, utils.HttpError{
			OriginalError: nil,
			Code:          404,
			Message:       "NOT FOUND",
		}
	}
	ans := model.Device{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &ans)
	if err != nil {
		return nil, err
	}
	return &ans, nil

}

func handelError(err error) (*model.Device, error) {
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeProvisionedThroughputExceededException, dynamodb.ErrCodeRequestLimitExceeded:
				return nil, utils.HttpError{OriginalError: err, Code: 429, Message: "message limit exceed"}
			}
		}
	}
	return nil, err

}

func tableName() *string {
	return aws.String(os.Getenv("DEVICES_TABLE"))
}

func CreateSession() *session.Session {
	return session.Must(session.NewSession(&aws.Config{
		Region:   aws.String(os.Getenv("AWS_REGION")),
		Endpoint: aws.String(os.Getenv("AWS_ENDPOINT")),
	}))
}
