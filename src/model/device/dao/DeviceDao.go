package dao

import (
	"serverless-aws-dynamodb/src/model/device"
)

type DeviceDao interface {
	PersistDevice(device *device.DeviceModel) (*device.DeviceModel, error)
	LoadDevice(id string) (*device.DeviceModel, error)
}
