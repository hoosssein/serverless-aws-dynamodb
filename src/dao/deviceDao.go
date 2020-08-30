package dao

/*
import "moghimi/myservice/src/model"

func SaveDevice(device *model.Device) (*model.Device, error) {
	return device,nil
}*/

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"moghimi/myservice/src/model"
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

func tableName() *string {
	return aws.String(os.Getenv("DEVICES_TABLE"))
}

func CreateSession() *session.Session {
	return session.Must(session.NewSession(&aws.Config{
		Region:   aws.String(os.Getenv("AWS_REGION")),
		Endpoint: aws.String(os.Getenv("AWS_ENDPOINT")),
	}))
}
