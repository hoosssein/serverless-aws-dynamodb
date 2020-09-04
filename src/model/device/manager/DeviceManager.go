package manager

import "serverless-aws-dynamodb/src/model/device"

type DeviceManager interface {
	SaveDevice(deviceModel *device.DeviceModel) (*device.DeviceModel, error)
	GetDevice(id string) (*device.DeviceModel, error)
}
