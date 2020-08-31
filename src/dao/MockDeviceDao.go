package dao

import "moghimi/myservice/src/model"

type MockDeviceDao struct {
	SaveDevice_ func(device *model.Device) (*model.Device, error)
	GetDevice_  func(id string) (*model.Device, error)
}

func (m MockDeviceDao) SaveDevice(device *model.Device) (*model.Device, error) {
	return m.SaveDevice_(device)
}
func (m MockDeviceDao) GetDevice(id string) (*model.Device, error) {
	return m.GetDevice_(id)
}
