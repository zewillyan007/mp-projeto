package entity

import (
	"mp-projeto/shark/core/err"
	"strings"
	"time"
)

type Microcontroller struct {
	Id               int64      `json:"id"`
	IdLocation       int64      `json:"id_location"`
	SerialNumber     string     `json:"serial_number"`
	Model            string     `json:"model"`
	Status           string     `json:"status"`
	CreationDateTime *time.Time `json:"creation_date_time"`
	ChangeDateTime   *time.Time `json:"change_date_time"`
}

func NewMicrocontroller() *Microcontroller {
	return &Microcontroller{
		Id:               0,
		SerialNumber:     "",
		Status:           "",
		CreationDateTime: &time.Time{},
		ChangeDateTime:   &time.Time{},
	}
}

func (o *Microcontroller) IsValid() error {
	if o.IdLocation == 0 {
		return err.MicrocontrollerErrorIdLocation
	}

	if len(strings.TrimSpace(o.SerialNumber)) == 0 {
		return err.MicrocontrollerErrorSerialNumber
	}

	if len(strings.TrimSpace(o.Model)) == 0 {
		return err.MicrocontrollerErrorModel
	}

	if len(strings.TrimSpace(o.Status)) == 0 {
		return err.MicrocontrollerErrorStatus
	}

	return nil
}
