package dao

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	. "moghimi/myservice/src/model/device"
	"moghimi/myservice/src/model/device/repository"
	"moghimi/myservice/src/utils"
	"testing"
)

type DynamodbRepositoryMock struct {
	repository.DynamoDBAPIExtendedImpl
	DynamodbMock        DynamodbMock
	SuccessFulUnmarshal bool
}

func (d DynamodbRepositoryMock) CreateSession() dynamodbiface.DynamoDBAPI {
	return d.DynamodbMock
}

const successfulUnmarshal = "valid id"

func (d DynamodbRepositoryMock) UnmarshalMap(m map[string]*dynamodb.AttributeValue, out interface{}) error {
	model, ok := out.(*DeviceModel)
	if !ok {
		return errors.New("invalid interface passed to UnmarshalMap")
	}
	if d.SuccessFulUnmarshal {
		model.Id = successfulUnmarshal
		return nil
	} else {
		return errors.New("fail to unmarshal")
	}

}

type DynamodbMock struct {
	dynamodbiface.DynamoDBAPI
	putItem func(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
	getItem func(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
}

func (d DynamodbMock) PutItem(p *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return d.putItem(p)
}
func (d DynamodbMock) GetItem(p *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	return d.getItem(p)
}

func Test_persistDevice_success(t *testing.T) {
	model := DeviceModel{
		Id:          "Id",
		DeviceModel: "DeviceModel",
		Name:        "Name",
		Note:        "Note",
		Serial:      "Serial",
	}
	dynamoDb := DynamodbMock{
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

	dao := DynamoDeviceDao{Repository: DynamodbRepositoryMock{
		DynamodbMock: dynamoDb,
	}}

	deviceModel, err := dao.PersistDevice(&model)
	utils.AssertEquals(t, nil, err, "error is not nil")
	utils.AssertEquals(t, model, *deviceModel, "model is not what expected")

}
func Test_persistDevice_dynamoReturnError(t *testing.T) {
	model := DeviceModel{
		Id:          "Id",
		DeviceModel: "DeviceModel",
		Name:        "Name",
		Note:        "Note",
		Serial:      "Serial",
	}
	dynamoDb := DynamodbMock{
		putItem: func(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
			return nil, errors.New("fail to persist")

		},
	}

	dao := DynamoDeviceDao{Repository: DynamodbRepositoryMock{
		DynamodbMock: dynamoDb,
	}}

	deviceModel, err := dao.PersistDevice(&model)
	utils.AssertEquals(t, "fail to persist", err.Error(), "error is not what expected")
	utils.AssertTrue(t, deviceModel == nil, "model is not nil")

}

func Test_getDevice_success(t *testing.T) {
	existingId := "Id"

	dynamoDb := DynamodbMock{
		getItem: func(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
			if *(input.Key["id"].S) == existingId {
				return &dynamodb.GetItemOutput{
					Item: map[string]*dynamodb.AttributeValue{
						"Id": {
							S: &existingId,
						},
					},
				}, nil
			}
			return nil, errors.New("entity not found")
		},
	}

	dao := DynamoDeviceDao{Repository: DynamodbRepositoryMock{
		DynamodbMock:        dynamoDb,
		SuccessFulUnmarshal: true,
	}}
	deviceModel, err := dao.LoadDevice(existingId)
	utils.AssertEquals(t, nil, err, "expecting no error")
	utils.AssertEquals(t, successfulUnmarshal, deviceModel.Id, "expecting successful unmarshal")

}

func Test_getDevice_fialToUnmarshal(t *testing.T) {
	existingId := "Id"

	dynamoDb := DynamodbMock{
		getItem: func(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
			if *(input.Key["id"].S) == existingId {
				return &dynamodb.GetItemOutput{
					Item: map[string]*dynamodb.AttributeValue{
						"Id": {
							S: &existingId,
						},
					},
				}, nil
			}
			return nil, errors.New("entity not found")
		},
	}

	dao := DynamoDeviceDao{Repository: DynamodbRepositoryMock{
		DynamodbMock:        dynamoDb,
		SuccessFulUnmarshal: false,
	}}
	deviceModel, err := dao.LoadDevice(existingId)
	var expectingDeviceModel *DeviceModel = nil
	utils.AssertEquals(t, "fail to unmarshal", err.Error(), "expecting error")
	utils.AssertEquals(t, expectingDeviceModel, deviceModel, "expecting successful unmarshal")

}

func Test_getDevice_idNotExist(t *testing.T) {
	existingId := "existingId"

	dynamoDb := DynamodbMock{
		getItem: func(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
			if *(input.Key["id"].S) == existingId {
				return nil, nil
			}
			return &dynamodb.GetItemOutput{}, nil
		},
	}

	dao := DynamoDeviceDao{Repository: DynamodbRepositoryMock{
		DynamodbMock:        dynamoDb,
		SuccessFulUnmarshal: true,
	}}
	deviceModel, err := dao.LoadDevice("not existing id")
	expectingError := utils.HttpError{
		OriginalError: nil,
		Code:          404,
		Message:       "NOT FOUND",
	}
	var expectingDeviceModel *DeviceModel = nil
	utils.AssertEquals(t, err, expectingError, "expecting error")
	utils.AssertEquals(t, expectingDeviceModel, deviceModel, "expecting nil as deviceModel")

}

func Test_getDevice_tooManyRequest(t *testing.T) {
	internalError := awserr.New(dynamodb.ErrCodeProvisionedThroughputExceededException, "my error", nil)
	dynamoDb := DynamodbMock{
		getItem: func(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
			return nil, internalError
		},
	}

	dao := DynamoDeviceDao{Repository: DynamodbRepositoryMock{
		DynamodbMock: dynamoDb,
	}}
	deviceModel, err := dao.LoadDevice("id")
	expectingError := utils.HttpError{
		OriginalError: internalError,
		Code:          429,
		Message:       "message limit exceed",
	}
	var expectingDeviceModel *DeviceModel = nil
	utils.AssertEquals(t, err, expectingError, "expecting error")
	utils.AssertEquals(t, expectingDeviceModel, deviceModel, "expecting nil as deviceModel")

}

func Test_getDevice_internalError(t *testing.T) {
	internalError := awserr.New("internal error", "my error", nil)
	dynamoDb := DynamodbMock{
		getItem: func(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
			return nil, internalError
		},
	}

	dao := DynamoDeviceDao{Repository: DynamodbRepositoryMock{
		DynamodbMock: dynamoDb,
	}}
	deviceModel, err := dao.LoadDevice("id")
	expectingError := internalError
	var expectingDeviceModel *DeviceModel = nil
	utils.AssertEquals(t, err, expectingError, "expecting error")
	utils.AssertEquals(t, expectingDeviceModel, deviceModel, "expecting nil as deviceModel")

}
