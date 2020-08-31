package deviceHandler

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"moghimi/myservice/src/api"
	"moghimi/myservice/src/dao"
	"moghimi/myservice/src/model"
	"moghimi/myservice/src/utils"
)

type DeviceHandler struct {
	Dao dao.DeviceDao
}

const IdPrefix = "/devices/"

func (handler DeviceHandler) Post(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	device := &model.Device{}
	err := json.Unmarshal([]byte(request.Body), &device)
	if err != nil {
		return api.SendError(&request, utils.HttpError{OriginalError: err, Code: 400, Message: err.Error()})
	}
	//todo check if id is valid

	err = device.Validate()
	if err != nil {
		return api.SendError(&request, utils.HttpError{OriginalError: err, Code: 400, Message: err.Error()})
	}

	device, err = handler.Dao.SaveDevice(device)
	if err != nil {
		return api.SendError(&request, err)
	}
	return api.SendJson(&request, &device, 201)
}

func (handler DeviceHandler) Get(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := request.PathParameters["id"]
	id = IdPrefix + id

	device, err := handler.Dao.GetDevice(id)
	if err != nil {
		return api.SendError(&request, err)
	}

	return api.SendJson(&request, &device, 200)
}
