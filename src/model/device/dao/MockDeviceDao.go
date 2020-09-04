package dao

import (
	"serverless-aws-dynamodb/src/model/device"
)

type MockDeviceDao struct {
	persistDevice func(device *device.DeviceModel) (*device.DeviceModel, error)
	loadDevice    func(id string) (*device.DeviceModel, error)
}

func NewMockDeviceDao(persistDevice func(device *device.DeviceModel) (*device.DeviceModel, error), loadDevice func(id string) (*device.DeviceModel, error)) *MockDeviceDao {
	return &MockDeviceDao{persistDevice: persistDevice, loadDevice: loadDevice}
}

func (m MockDeviceDao) PersistDevice(device *device.DeviceModel) (*device.DeviceModel, error) {
	return m.persistDevice(device)
}
func (m MockDeviceDao) LoadDevice(id string) (*device.DeviceModel, error) {
	return m.loadDevice(id)
}
