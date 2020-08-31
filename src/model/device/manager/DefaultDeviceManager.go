package manager

import (
	"moghimi/myservice/src/model/device"
	"moghimi/myservice/src/model/device/dao"
	"moghimi/myservice/src/utils"
)

type DefaultDeviceManager struct {
	Dao dao.DeviceDao
}

func (this DefaultDeviceManager) SaveDevice(deviceModel *device.DeviceModel) (*device.DeviceModel, error) {
	//todo check if id is valid

	err := deviceModel.Validate()
	if err != nil {
		return nil, utils.HttpError{OriginalError: err, Code: 400, Message: err.Error()}
	}
	return this.Dao.PersistDevice(deviceModel)
}

func (this DefaultDeviceManager) GetDevice(id string) (*device.DeviceModel, error) {

	return this.Dao.LoadDevice(id)
}
