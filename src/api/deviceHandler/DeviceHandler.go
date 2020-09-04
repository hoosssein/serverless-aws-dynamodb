package deviceHandler

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"serverless-aws-dynamodb/src/api"
	"serverless-aws-dynamodb/src/model/device"
	"serverless-aws-dynamodb/src/model/device/dao"
	"serverless-aws-dynamodb/src/model/device/manager"
	"serverless-aws-dynamodb/src/model/device/repository"
	"serverless-aws-dynamodb/src/utils"
	"serverless-aws-dynamodb/src/utils/config"
)

type DeviceHandler struct {
	Manager manager.DeviceManager
}

var Handler = DeviceHandler{
	Manager: manager.DefaultDeviceManager{
		Dao: dao.DynamoDeviceDao{
			Repository: repository.NewDynamoDBAPIExtended(),
		},
	},
}

func (this DeviceHandler) Post(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	deviceModel := &device.DeviceModel{}
	err := json.Unmarshal([]byte(request.Body), &deviceModel)
	if err != nil {
		return api.SendError(&request, utils.HttpError{OriginalError: err, Code: 400, Message: err.Error()})
	}

	deviceModel, err = this.Manager.SaveDevice(deviceModel)
	if err != nil {
		return api.SendError(&request, err)
	}
	return api.SendJson(&request, &deviceModel, 201)
}

func (this DeviceHandler) Get(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := request.PathParameters["id"]
	if id == "" {
		return api.SendError(&request, utils.HttpError{Code: 400, Message: "id is empty"})
	}

	id = config.IdPrefix + id

	deviceModel, err := this.Manager.GetDevice(id)
	if err != nil {
		return api.SendError(&request, err)
	}

	return api.SendJson(&request, &deviceModel, 200)
}
