package entity

import (
	"mp-projeto/shark/core/err"
	"strings"
	"time"
)

type Chip struct {
	Id               int64      `json:"id"`
	Number           string     `json:"number"`
	CreationDateTime *time.Time `json:"creation_date_time"`
	ChangeDateTime   *time.Time `json:"change_date_time"`
}

func NewChip() *Chip {
	return &Chip{
		Id:     0,
		Number: "",
	}
}

func (o *Chip) IsValid() error {
	if len(strings.TrimSpace(o.Number)) == 0 {
		return err.ChipErrorNumber
	}

	return nil
}
