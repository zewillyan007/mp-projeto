package entity

import "time"

type Incidence struct {
	Id                int64      `json:"id"`
	IdShark           int64      `json:"id_shark"`
	Name              string     `json:"name"`
	IncidenceDateTime *time.Time `json:"incidence_date_time"`
}

func NewIncidence() *Incidence {
	return &Incidence{
		Id:                0,
		IdShark:           0,
		Name:              "",
		IncidenceDateTime: &time.Time{},
	}
}
