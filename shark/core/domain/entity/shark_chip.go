package entity

import (
	"mp-projeto/shark/core/err"
	"time"
)

type SharkChip struct {
	Id               int64      `json:"id"`
	IdShark          int64      `json:"id_shark"`
	IdChip           int64      `json:"id_chip"`
	CreationDateTime *time.Time `json:"creation_date_time"`
}

func NewSharkChip() *SharkChip {
	return &SharkChip{
		Id:               0,
		IdShark:          0,
		IdChip:           0,
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

	return nil
}
