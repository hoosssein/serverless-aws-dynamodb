package dao

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
	"serverless-aws-dynamodb/src/model/device"
	"serverless-aws-dynamodb/src/model/device/repository"
	"serverless-aws-dynamodb/src/utils"
)

type DynamoDeviceDao struct {
	Repository repository.DynamoDBAPIExtended
}

func (this DynamoDeviceDao) PersistDevice(device *device.DeviceModel) (*device.DeviceModel, error) {
	dynamoAttribute, err := this.Repository.MarshalMap(device)
	if err != nil {
		return nil, err
	}
	fmt.Println("dynamoAttribute: ", dynamoAttribute)
	input := &dynamodb.PutItemInput{
		Item:      dynamoAttribute,
		TableName: tableName(),
	}

	_, err = this.Repository.CreateSession().PutItem(input)
	if err != nil {
		return nil, err
	}

	return device, nil
}

func (this DynamoDeviceDao) LoadDevice(id string) (*device.DeviceModel, error) {
	result, err := this.Repository.CreateSession().GetItem(&dynamodb.GetItemInput{
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
	ans := device.DeviceModel{}
	err = this.Repository.UnmarshalMap(result.Item, &ans)
	if err != nil {
		return nil, err
	}
	return &ans, nil
}

func handelError(err error) (*device.DeviceModel, error) {
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
