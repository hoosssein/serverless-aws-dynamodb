package device

type DeviceModel struct {
	Id          string `json:"id" validate:"startswith=/devices/"`
	DeviceModel string `json:"deviceModel" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Note        string `json:"note" validate:"required"`
	Serial      string `json:"serial" validate:"required"`
}
