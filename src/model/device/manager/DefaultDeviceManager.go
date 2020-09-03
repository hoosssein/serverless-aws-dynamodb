package manager

import (
	"errors"
	"moghimi/myservice/src/model/device"
	"moghimi/myservice/src/model/device/dao"
	"moghimi/myservice/src/utils"
	"moghimi/myservice/src/utils/config"
	"strings"
)

type DefaultDeviceManager struct {
	Dao dao.DeviceDao
}

func (this DefaultDeviceManager) SaveDevice(deviceModel *device.DeviceModel) (*device.DeviceModel, error) {

	err := Validate(deviceModel)
	if err != nil {
		return nil, utils.HttpError{OriginalError: err, Code: 400, Message: err.Error()}
	}
	return this.Dao.PersistDevice(deviceModel)
}

func (this DefaultDeviceManager) GetDevice(id string) (*device.DeviceModel, error) {

	return this.Dao.LoadDevice(id)
}

func Validate(d *device.DeviceModel) error {
	if d.DeviceModel == "" {
		return errors.New("deviceModel is empty")
	}
	if d.Name == "" {
		return errors.New("name is empty")
	}
	if d.Note == "" {
		return errors.New("note is empty")
	}
	if d.Serial == "" {
		return errors.New("serial is empty")
	}
	if !strings.HasPrefix(d.Id, config.IdPrefix) {
		return errors.New("id should start with " + config.IdPrefix)
	}
	if len(d.Id) <= len(config.IdPrefix) {
		return errors.New("id should contains strings after " + config.IdPrefix)
	}
	return nil
}
