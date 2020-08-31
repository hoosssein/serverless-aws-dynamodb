package manager

import "moghimi/myservice/src/model/device"

type DeviceManager interface {
	SaveDevice(deviceModel *device.DeviceModel) (*device.DeviceModel, error)
	GetDevice(id string) (*device.DeviceModel, error)
}
