package dto

type IncidenceDtoIn struct {
	Id                string
	IdShark           string
	Name              string
	IncidenceDateTime string
}

func NewIncidenceDtoIn() *IncidenceDtoIn {
	return &IncidenceDtoIn{
		Id:                "",
		IdShark:           "",
		Name:              "",
		IncidenceDateTime: "",
	}
}

type IncidenceDtoOut struct {
	Id                string
	IdShark           string
	Name              string
	IncidenceDateTime string
}

func NewIncidenceDtoOut() *IncidenceDtoOut {
	return &IncidenceDtoOut{
		Id:                "",
		IdShark:           "",
		Name:              "",
		IncidenceDateTime: "",
	}
}
