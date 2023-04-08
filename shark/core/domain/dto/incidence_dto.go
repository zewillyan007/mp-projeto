package dto

type IncidenceDtoIn struct {
	Id                string
	ChipNumber        string
	IncidenceDateTime string
}

func NewIncidenceDtoIn() *IncidenceDtoIn {
	return &IncidenceDtoIn{
		Id:                "",
		ChipNumber:        "",
		IncidenceDateTime: "",
	}
}

type IncidenceDtoOut struct {
	Id                string
	ChipNumber        string
	IncidenceDateTime string
}

func NewIncidenceDtoOut() *IncidenceDtoOut {
	return &IncidenceDtoOut{
		Id:                "",
		ChipNumber:        "",
		IncidenceDateTime: "",
	}
}
