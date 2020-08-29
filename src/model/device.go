package model

import (
	"github.com/go-playground/validator"
)

type Device struct {
	Id          string `json:"id" validate:"required"`
	DeviceModel string `json:"deviceModel" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Note        string `json:"note" validate:"required"`
	Serial      string `json:"serial" validate:"required"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (d Device) Validate() error {
	err := validate.Struct(d)
	return err
}
