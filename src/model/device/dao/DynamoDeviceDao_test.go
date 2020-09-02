package dao

import (
	"errors"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"moghimi/myservice/src/model/device"
	"moghimi/myservice/src/utils"
	"testing"
)

type DynamoMock struct {
	dynamodbiface.DynamoDBAPI
	putItem func(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
}

func (d DynamoMock) PutItem(p *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return d.putItem(p)
}

func Test_persistDevice_success(t *testing.T) {
	model := device.DeviceModel{
		Id:          "Id",
		DeviceModel: "DeviceModel",
		Name:        "Name",
		Note:        "Note",
		Serial:      "Serial",
	}
	dynamoDb := DynamoMock{
		putItem: func(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
			expected := map[string]*dynamodb.AttributeValue{
				"id":          {S: &(model.Id)},
				"deviceModel": {S: &(model.DeviceModel)},
				"name":        {S: &(model.Name)},
				"note":        {S: &(model.Note)},
				"serial":      {S: &(model.Serial)},
			}
			item := input.Item
			if *(expected["id"].S) == *(item["id"].S) && *(expected["note"].S) == *(item["note"].S) {
				return &dynamodb.PutItemOutput{}, nil
			}
			return nil, errors.New("invalid input for putItem")
		},
	}

	deviceModel, err := persistDevice(dynamoDb, &model)
	utils.AssertEquals(t, nil, err, "error is not nil")
	utils.AssertEquals(t, model, *deviceModel, "model is not what expected")

}
func Test_persistDevice_dynamoReturnError(t *testing.T) {
	model := device.DeviceModel{
		Id:          "Id",
		DeviceModel: "DeviceModel",
		Name:        "Name",
		Note:        "Note",
		Serial:      "Serial",
	}
	dynamoDb := DynamoMock{
		putItem: func(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
			return nil, errors.New("fail to persist")

		},
	}

	deviceModel, err := persistDevice(dynamoDb, &model)
	utils.AssertEquals(t, "fail to persist", err.Error(), "error is not what expected")
	utils.AssertTrue(t, deviceModel == nil, "model is not nil")

}
