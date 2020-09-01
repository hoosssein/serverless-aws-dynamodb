package manager

import (
	"fmt"
	"moghimi/myservice/src/model/device"
	"moghimi/myservice/src/model/device/dao"
	. "moghimi/myservice/src/utils"
	"moghimi/myservice/src/utils/config"
	"testing"
)

func TestDefaultDeviceManager_SaveDevice_EmptyField_returnError(t *testing.T) {
	manager := DefaultDeviceManager{}
	_, err := manager.SaveDevice(&device.DeviceModel{})
	AssertTrue(t, err != nil, " expect error ")
	httpError, ok := err.(HttpError)
	AssertTrue(t, ok, fmt.Sprint(" bad error type ", ok))
	AssertEquals(t, httpError.Code, 400, "bad status code")
}

func TestDefaultDeviceManager_SaveDevice_IdIsInvalid_returnError(t *testing.T) {
	manager := DefaultDeviceManager{}
	deviceModel := device.DeviceModel{
		Id:          "/badPrefix/id1",
		DeviceModel: "STH",
		Name:        "STH",
		Note:        "STH",
		Serial:      "STH",
	}
	_, err := manager.SaveDevice(&deviceModel)
	AssertTrue(t, err != nil, " expect error ")
	httpError, ok := err.(HttpError)
	AssertTrue(t, ok, fmt.Sprint(" bad error type ", ok))
	AssertEquals(t, httpError.Code, 400, "bad status code")
}

func TestDefaultDeviceManager_SaveDevice_AllFieldsAreFull_callDao(t *testing.T) {
	deviceModel := device.DeviceModel{
		Id:          config.IdPrefix + "id1",
		DeviceModel: "STH",
		Name:        "STH",
		Note:        "STH",
		Serial:      "STH",
	}
	manager := DefaultDeviceManager{
		Dao: dao.NewMockDeviceDao(func(device *device.DeviceModel) (*device.DeviceModel, error) {
			return device, nil
		}, nil),
	}
	savedDeviceModel, err := manager.SaveDevice(&deviceModel)
	AssertTrue(t, err == nil, " expect no error ")
	AssertEquals(t, deviceModel, *savedDeviceModel, " deviceModel not match")

}

func TestDefaultDeviceManager_GetDevice_callDao(t *testing.T) {
	id := "id1"
	deviceModel := device.DeviceModel{
		Id:          config.IdPrefix + id,
		DeviceModel: "STH",
		Name:        "STH",
		Note:        "STH",
		Serial:      "STH",
	}
	manager := DefaultDeviceManager{
		Dao: dao.NewMockDeviceDao(nil, func(id string) (*device.DeviceModel, error) {
			return &deviceModel, nil
		}),
	}
	savedDeviceModel, err := manager.GetDevice(id)
	AssertTrue(t, err == nil, " expect no error ")
	AssertEquals(t, deviceModel, *savedDeviceModel, " deviceModel not match")

}
