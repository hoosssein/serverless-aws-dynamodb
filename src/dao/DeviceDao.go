package dao

import "moghimi/myservice/src/model"

type DeviceDao interface {
	SaveDevice(device *model.Device) (*model.Device, error)
	GetDevice(id string) (*model.Device, error)
}
