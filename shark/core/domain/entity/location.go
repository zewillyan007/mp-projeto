package entity

import (
	"mp-projeto/shark/core/err"
	"strings"
	"time"
)

type Location struct {
	Id               int64      `json:"id"`
	Name             string     `json:"name"`
	CreationDateTime *time.Time `json:"creation_date_time"`
	ChangeDateTime   *time.Time `json:"change_date_time"`
}

func NewLocation() *Location {
	return &Location{
		Id:               0,
		Name:             "",
		CreationDateTime: &time.Time{},
		ChangeDateTime:   &time.Time{},
	}
}

func (o *Location) IsValid() error {
	if len(strings.TrimSpace(o.Name)) == 0 {
		return err.LocationErrorName
	}

	return nil
}
