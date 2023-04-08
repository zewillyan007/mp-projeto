package entity

import (
	"mp-projeto/shark/core/err"
	"strings"
	"time"
)

type Shark struct {
	Id               int64      `json:"id"`
	Species          string     `json:"species"`
	Length           float64    `json:"length"`
	Weight           float64    `json:"weight"`
	Sex              string     `json:"sex"`
	CreationDateTime *time.Time `json:"creation_date_time"`
	ChangeDateTime   *time.Time `json:"change_date_time"`
}

func NewShark() *Shark {
	return &Shark{
		Id:               0,
		Species:          "",
		Length:           0,
		Weight:           0,
		Sex:              "",
		CreationDateTime: &time.Time{},
		ChangeDateTime:   &time.Time{},
	}
}

func (o *Shark) IsValid() error {
	if len(strings.TrimSpace(o.Species)) == 0 {
		return err.SharkErrorSpecies
	}

	return nil
}
