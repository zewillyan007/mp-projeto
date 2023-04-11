package entity

import (
	"mp-projeto/shark/core/err"
	"strings"
	"time"
)

type Incidence struct {
	Id                          int64      `json:"id"`
	ChipNumber                  string     `json:"chip_number"`
	MicrocontrollerSerialNumber string     `json:"microcontroller_serial_number"`
	IncidenceDateTime           *time.Time `json:"incidence_date_time"`
}

func NewIncidence() *Incidence {
	return &Incidence{
		Id:                          0,
		ChipNumber:                  "",
		MicrocontrollerSerialNumber: "",
		IncidenceDateTime:           &time.Time{},
	}
}

func (o *Incidence) IsValid() error {
	if len(strings.TrimSpace(o.ChipNumber)) == 0 {
		return err.IncidenceErrorChipNumber
	}

	if len(strings.TrimSpace(o.MicrocontrollerSerialNumber)) == 0 {
		return err.IncidenceErrorMicrocontrollerSerialNumber
	}

	return nil
}
