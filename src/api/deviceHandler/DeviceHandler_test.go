package deviceHandler

import (
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"moghimi/myservice/src/model/device"
	"moghimi/myservice/src/model/device/manager"
	"moghimi/myservice/src/utils"
	"moghimi/myservice/src/utils/config"
	"reflect"
	"testing"
)

const existingId = "existingId"

var existingDevice = device.DeviceModel{
	Id:          config.IdPrefix + existingId,
	DeviceModel: "DM",
	Name:        "NAME",
	Note:        "NOTE",
	Serial:      "SERIAL",
}

func TestDeviceHandler_Get(t *testing.T) {
	type fields struct {
		Manager manager.DeviceManager
	}
	type args struct {
		request events.APIGatewayProxyRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    events.APIGatewayProxyResponse
		wantErr bool
	}{
		{
			name: "find existing id",
			fields: fields{
				Manager: mockManagerGetDevice(),
			},
			args: args{
				request: events.APIGatewayProxyRequest{
					PathParameters: map[string]string{
						"id": existingId,
					},
				},
			},
			want:    existingDeviceResponse(200),
			wantErr: false,
		},
		{
			name:   "id is empty",
			fields: fields{},
			args: args{
				request: events.APIGatewayProxyRequest{
					PathParameters: map[string]string{},
				},
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: 400,
				Body:       "id is empty",
			},
			wantErr: false,
		},
		{
			name: "id not exists",
			fields: fields{
				Manager: mockManagerGetDevice(),
			},
			args: args{
				request: events.APIGatewayProxyRequest{
					PathParameters: map[string]string{
						"id": "not existing",
					},
				},
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: 404,
				Body:       "NOT FOUND",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := DeviceHandler{
				Manager: tt.fields.Manager,
			}
			got, err := handler.Get(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func existingDeviceResponse(statusCode int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: ignoreError(json.Marshal(existingDevice)),
	}
}

func ignoreError(input []byte, err error) string {
	return string(input)
}

func mockManagerGetDevice() *manager.MockDeviceManager {
	return manager.NewMockDeviceManager(nil, func(id string) (*device.DeviceModel, error) {
		if id == config.IdPrefix+existingId {
			return &existingDevice, nil
		}
		return nil, utils.HttpError{
			OriginalError: nil,
			Code:          404,
			Message:       "NOT FOUND",
		}
	})
}
func TestDeviceHandler_Post(t *testing.T) {
	type fields struct {
		Manager manager.DeviceManager
	}
	type args struct {
		request events.APIGatewayProxyRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    events.APIGatewayProxyResponse
		wantErr bool
	}{
		{
			name: "save successfully",
			fields: fields{
				Manager: manager.NewMockDeviceManager(func(device *device.DeviceModel) (*device.DeviceModel, error) {
					return device, nil
				}, nil),
			},
			args: args{
				request: events.APIGatewayProxyRequest{
					Body: ignoreError(json.Marshal(existingDevice)),
				},
			},
			want:    existingDeviceResponse(201),
			wantErr: false,
		},
		{
			name: "manager reject device",
			fields: fields{
				Manager: manager.NewMockDeviceManager(func(device *device.DeviceModel) (*device.DeviceModel, error) {
					return device, utils.HttpError{
						OriginalError: nil,
						Code:          402,
						Message:       "SOME ERROR",
					}
				}, nil),
			},
			args: args{
				request: events.APIGatewayProxyRequest{
					Body: ignoreError(json.Marshal(existingDevice)),
				},
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: 402,
				Body:       "SOME ERROR",
			},
			wantErr: false,
		}, {
			name: "manager reject device with internal error",
			fields: fields{
				Manager: manager.NewMockDeviceManager(func(device *device.DeviceModel) (*device.DeviceModel, error) {
					return device, errors.New("bad error")
				}, nil),
			},
			args: args{
				request: events.APIGatewayProxyRequest{
					Body: ignoreError(json.Marshal(existingDevice)),
				},
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: 500,
				Body:       "INTERNAL SERVER ERROR",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := DeviceHandler{
				Manager: tt.fields.Manager,
			}
			got, err := handler.Post(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Post() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Post() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeviceHandler_Post_failToParse(t *testing.T) {
	handler := DeviceHandler{}
	ans, err := handler.Post(events.APIGatewayProxyRequest{
		Body: "error" + ignoreError(json.Marshal(existingDevice)),
	})
	if err != nil {
		t.Error("method return err")
		return
	}
	if ans.StatusCode != 400 {
		t.Error("expecting statusCode: 400, got:", ans.StatusCode)
	}

}
