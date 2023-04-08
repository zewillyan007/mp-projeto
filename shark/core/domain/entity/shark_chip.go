package entity

import (
	"mp-projeto/shark/core/err"
	"strings"
	"time"
)

type SharkChip struct {
	Id               int64      `json:"id"`
	IdShark          int64      `json:"id_shark"`
	IdChip           int64      `json:"id_chip"`
	ChipNumber       string     `json:"chip_number"`
	Status           string     `json:"status"`
	CreationDateTime *time.Time `json:"creation_date_time"`
}

func NewSharkChip() *SharkChip {
	return &SharkChip{
		Id:               0,
		IdShark:          0,
		IdChip:           0,
		ChipNumber:       "",
		Status:           "",
		CreationDateTime: &time.Time{},
	}
}

func (o *SharkChip) IsValid() error {
	if o.IdShark == 0 {
		return err.SharkChipErrorIdShark
	}

	if o.IdChip == 0 {
		return err.SharkChipErrorIdChip
	}

	if len(strings.TrimSpace(o.ChipNumber)) == 0 {
		return err.SharkChipErrorChipNumber
	}

	if len(strings.TrimSpace(o.Status)) == 0 {
		return err.SharkChipErrorStatus
	}

	return nil
}
