package entity

import (
	"mp-projeto/shark/core/err"
	"strings"
	"time"
)

type Chip struct {
	Id               int64      `json:"id"`
	Number           string     `json:"number"`
	Status           string     `json:"status"`
	CreationDateTime *time.Time `json:"creation_date_time"`
	ChangeDateTime   *time.Time `json:"change_date_time"`
}

func NewChip() *Chip {
	return &Chip{
		Id:               0,
		Number:           "",
		Status:           "",
		CreationDateTime: &time.Time{},
		ChangeDateTime:   &time.Time{},
	}
}

func (o *Chip) IsValid() error {
	if len(strings.TrimSpace(o.Number)) == 0 {
		return err.ChipErrorNumber
	}

	if len(strings.TrimSpace(o.Status)) == 0 {
		return err.ChipErrorStatus
	}

	return nil
}
