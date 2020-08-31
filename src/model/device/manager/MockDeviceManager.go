package manager

import (
	"moghimi/myservice/src/model/device"
)

type MockDeviceManager struct {
	saveDevice func(device *device.DeviceModel) (*device.DeviceModel, error)
	getDevice  func(id string) (*device.DeviceModel, error)
}

func NewMockDeviceManager(saveDevice func(device *device.DeviceModel) (*device.DeviceModel, error), getDevice func(id string) (*device.DeviceModel, error)) *MockDeviceManager {
	return &MockDeviceManager{saveDevice: saveDevice, getDevice: getDevice}
}

func (m MockDeviceManager) SaveDevice(device *device.DeviceModel) (*device.DeviceModel, error) {
	return m.saveDevice(device)
}
func (m MockDeviceManager) GetDevice(id string) (*device.DeviceModel, error) {
	return m.getDevice(id)
}
