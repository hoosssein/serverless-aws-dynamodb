package deviceHandler

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"moghimi/myservice/src/api"
	"moghimi/myservice/src/model/device"
	"moghimi/myservice/src/model/device/manager"
	"moghimi/myservice/src/utils"
	"moghimi/myservice/src/utils/config"
)

type DeviceHandler struct {
	Manager manager.DeviceManager
}

func (handler DeviceHandler) Post(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	deviceModel := &device.DeviceModel{}
	err := json.Unmarshal([]byte(request.Body), &deviceModel)
	if err != nil {
		return api.SendError(&request, utils.HttpError{OriginalError: err, Code: 400, Message: err.Error()})
	}

	deviceModel, err = handler.Manager.SaveDevice(deviceModel)
	if err != nil {
		return api.SendError(&request, err)
	}
	return api.SendJson(&request, &deviceModel, 201)
}

func (handler DeviceHandler) Get(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := request.PathParameters["id"]
	if id == "" {
		return api.SendError(&request, utils.HttpError{Code: 400, Message: "id is empty"})
	}

	id = config.IdPrefix + id

	deviceModel, err := handler.Manager.GetDevice(id)
	if err != nil {
		return api.SendError(&request, err)
	}

	return api.SendJson(&request, &deviceModel, 200)
}
